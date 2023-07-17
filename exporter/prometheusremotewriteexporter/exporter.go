// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/multierr"

	prometheustranslator "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"
)

const maxBatchByteSize = 3000000

// prwExporter converts OTLP metrics to Prometheus remote write TimeSeries and sends them to a remote endpoint.
type prwExporter struct {
	endpointURL     *url.URL
	client          *http.Client
	wg              *sync.WaitGroup
	closeChan       chan struct{}
	concurrency     int
	userAgentHeader string
	clientSettings  *confighttp.HTTPClientSettings
	settings        component.TelemetrySettings

	wal              *prweWAL
	exporterSettings prometheusremotewrite.Settings
}

// newPRWExporter initializes a new prwExporter instance and sets fields accordingly.
func newPRWExporter(cfg *Config, set exporter.CreateSettings) (*prwExporter, error) {
	sanitizedLabels, err := validateAndSanitizeExternalLabels(cfg)
	if err != nil {
		return nil, err
	}

	endpointURL, err := url.ParseRequestURI(cfg.HTTPClientSettings.Endpoint)
	if err != nil {
		return nil, errors.New("invalid endpoint")
	}

	userAgentHeader := fmt.Sprintf("%s/%s", strings.ReplaceAll(strings.ToLower(set.BuildInfo.Description), " ", "-"), set.BuildInfo.Version)

	prwe := &prwExporter{
		endpointURL:     endpointURL,
		wg:              new(sync.WaitGroup),
		closeChan:       make(chan struct{}),
		userAgentHeader: userAgentHeader,
		concurrency:     cfg.RemoteWriteQueue.NumConsumers,
		clientSettings:  &cfg.HTTPClientSettings,
		settings:        set.TelemetrySettings,
		exporterSettings: prometheusremotewrite.Settings{
			Namespace:           cfg.Namespace,
			ExternalLabels:      sanitizedLabels,
			DisableTargetInfo:   !cfg.TargetInfo.Enabled,
			ExportCreatedMetric: cfg.CreatedMetric.Enabled,
			AddMetricSuffixes:   cfg.AddMetricSuffixes,
		},
	}

	if cfg.WAL == nil {
		return prwe, nil
	}

	prwe.wal, err = newWAL(cfg.WAL, prwe.export)
	if err != nil {
		return nil, err
	}

	return prwe, nil
}

// Start creates the prometheus client
func (prwe *prwExporter) Start(ctx context.Context, host component.Host) (err error) {
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient, err = prwe.clientSettings.ToClient(host, prwe.settings)
	if err != nil {
		return err
	}

	// Configure retry settings
	// Default settings reference - https://github.com/hashicorp/go-retryablehttp/blob/571a88bc9c3b7c64575f0e9b0f646af1510f2c76/client.go#L51-L53
	retryClient.RetryWaitMin = 100 * time.Millisecond
	retryClient.RetryWaitMax = 1 * time.Second
	retryClient.RetryMax = 3

	prwe.client = retryClient.HTTPClient
	return prwe.turnOnWALIfEnabled(contextWithLogger(ctx, prwe.settings.Logger.Named("prw.wal")))
}

func (prwe *prwExporter) shutdownWALIfEnabled() error {
	if !prwe.walEnabled() {
		return nil
	}
	return prwe.wal.stop()
}

// Shutdown stops the exporter from accepting incoming calls(and return error), and wait for current export operations
// to finish before returning
func (prwe *prwExporter) Shutdown(context.Context) error {
	select {
	case <-prwe.closeChan:
	default:
		close(prwe.closeChan)
	}
	err := prwe.shutdownWALIfEnabled()
	prwe.wg.Wait()
	return err
}

// PushMetrics converts metrics to Prometheus remote write TimeSeries and send to remote endpoint. It maintain a map of
// TimeSeries, validates and handles each individual metric, adding the converted TimeSeries to the map, and finally
// exports the map.
func (prwe *prwExporter) PushMetrics(ctx context.Context, md pmetric.Metrics) error {
	prwe.wg.Add(1)
	defer prwe.wg.Done()

	select {
	case <-prwe.closeChan:
		return errors.New("shutdown has been called")
	default:
		tsMap, err := prometheusremotewrite.FromMetrics(md, prwe.exporterSettings)
		if err != nil {
			err = consumererror.NewPermanent(err)
		}
		// Call export even if a conversion error, since there may be points that were successfully converted.
		return multierr.Combine(err, prwe.handleExport(ctx, tsMap))
	}
}

func validateAndSanitizeExternalLabels(cfg *Config) (map[string]string, error) {
	sanitizedLabels := make(map[string]string)
	for key, value := range cfg.ExternalLabels {
		if key == "" || value == "" {
			return nil, fmt.Errorf("prometheus remote write: external labels configuration contains an empty key or value")
		}
		sanitizedLabels[prometheustranslator.NormalizeLabel(key)] = value
	}

	return sanitizedLabels, nil
}

/*//Signature for the time series
func tSSignature(labels *[]prompb.Label) string {
	var labelPairs []string
	for _, label := range *labels {
		labelPairs = append(labelPairs, fmt.Sprintf("%s=%s", label.Name, label.Value ))
	}

	// Sort the labels.
	sort.Strings(labelPairs)

	sign := strings.Join(labelPairs, ",")
	hashDigest := sha256.Sum256([]byte(sign))

	return fmt.Sprintf("%x", hashDigest)
}*/

//This function takes a map of time series as input and returns a slice of arrays, each array contains the time series with the same hash value.
func partitionTimeSeries(tsMap map[string]*prompb.TimeSeries) [][]*prompb.TimeSeries {
	partitionedTS := make([][]*prompb.TimeSeries, len(tsMap))
	index := 0
	for _, ts := range tsMap {
		partitionedTS[index] = []*prompb.TimeSeries{ts}
		index++
	}
	return partitionedTS
}

// handleExport handles the export of metrics by creating worker goroutines and channels.
// It waits for all workers to finish and returns an error if any worker encountered one.
func (prwe *prwExporter) handleExport(ctx context.Context, tsMap map[string]*prompb.TimeSeries) error {
	// There are no metrics to export, so return.
	if len(tsMap) == 0 {
		return nil
	}

	partitionedTS := partitionTimeSeries(tsMap)

	// Create channels and error channel for each partitioned array.
	channels := make([]chan []*prompb.TimeSeries, len(partitionedTS))
	errCh := make(chan error, len(partitionedTS))

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(len(partitionedTS))

	// Create channels and launch goroutines to process each partition
	for i := range partitionedTS {
		channels[i] = make(chan []*prompb.TimeSeries)
		go prwe.exportWorker(ctx, channels[i], errCh, &wg)
		channels[i] <- partitionedTS[i] // Send the partitioned array to the corresponding channel
		close(channels[i])              // Close the channel after sending the data
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Collect errors from the error channel
	var errs []error
	for err := range errCh {
		if err != nil {
			errs = append(errs, err)
		}
	}

	return multierr.Combine(errs...)
}

// exportWorker is the worker function that consumes a channel of partitioned time series and performs the export.
// It sends any encountered error to the error channel.
func (prwe *prwExporter) exportWorker(ctx context.Context, tsArrayChan <-chan []*prompb.TimeSeries, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	for tsArray := range tsArrayChan {
		// Iterate over the partitioned time series sequentially
		for _, ts := range tsArray {
			// Calls the helper function to convert and batch the time series to the desired format
			requests, err := batchTimeSeries(map[string]*prompb.TimeSeries{ts.String(): ts}, maxBatchByteSize)
			if err != nil {
				errCh <- err
				return
			}

			if !prwe.walEnabled() {
				// Perform a direct export otherwise.
				if err := prwe.export(ctx, requests); err != nil {
					errCh <- err
				}
			} else {
				// Otherwise, the WAL is enabled, and just persist the requests to the WAL
				// and they'll be exported in another goroutine to the RemoteWrite endpoint.
				if err := prwe.wal.persistToWAL(requests); err != nil {
					errCh <- err
				}
			}
		}
	}

	errCh <- nil
}

// export sends a Snappy-compressed WriteRequest containing TimeSeries to a remote write endpoint in order
func (prwe *prwExporter) export(ctx context.Context, requests []*prompb.WriteRequest) error {
	input := make(chan *prompb.WriteRequest, len(requests))
	for _, request := range requests {
		input <- request
	}
	close(input)

	var wg sync.WaitGroup

	concurrencyLimit := int(math.Min(float64(prwe.concurrency), float64(len(requests))))
	wg.Add(concurrencyLimit) // used to wait for workers to be finished

	var mu sync.Mutex
	var errs error
	// Run concurrencyLimit of workers until there
	// is no more requests to execute in the input channel.
	for i := 0; i < concurrencyLimit; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done(): // Check firstly to ensure that the context wasn't cancelled.
					return

				case request, ok := <-input:
					if !ok {
						return
					}
					if errExecute := prwe.execute(ctx, request); errExecute != nil {
						mu.Lock()
						errs = multierr.Append(errs, consumererror.NewPermanent(errExecute))
						mu.Unlock()
					}
				}
			}
		}()
	}
	wg.Wait()

	return errs
}

func (prwe *prwExporter) execute(ctx context.Context, writeReq *prompb.WriteRequest) error {
	// Uses proto.Marshal to convert the WriteRequest into bytes array
	data, err := proto.Marshal(writeReq)
	if err != nil {
		return consumererror.NewPermanent(err)
	}
	buf := make([]byte, len(data), cap(data))
	compressedData := snappy.Encode(buf, data)

	// Create the HTTP POST request to send to the endpoint
	req, err := http.NewRequestWithContext(ctx, "POST", prwe.endpointURL.String(), bytes.NewReader(compressedData))
	if err != nil {
		return consumererror.NewPermanent(err)
	}

	// Add necessary headers specified by:
	// https://cortexmetrics.io/docs/apis/#remote-api
	req.Header.Add("Content-Encoding", "snappy")
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("X-Prometheus-Remote-Write-Version", "0.1.0")
	req.Header.Set("User-Agent", prwe.userAgentHeader)

	resp, err := prwe.client.Do(req)
	if err != nil {
		return consumererror.NewPermanent(err)
	}
	defer resp.Body.Close()

	// 2xx status code is considered a success
	// 5xx errors are recoverable and the exporter should retry
	// Reference for different behavior according to status code:
	// https://github.com/prometheus/prometheus/pull/2552/files#diff-ae8db9d16d8057358e49d694522e7186
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 256))
	rerr := fmt.Errorf("remote write returned HTTP status %v; err = %w: %s", resp.Status, err, body)
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		// Return a non-permanent error for 5xx status codes to indicate that it's recoverable
		return rerr
	}
	return consumererror.NewPermanent(rerr)
}

func (prwe *prwExporter) walEnabled() bool { return prwe.wal != nil }

func (prwe *prwExporter) turnOnWALIfEnabled(ctx context.Context) error {
	if !prwe.walEnabled() {
		return nil
	}
	cancelCtx, cancel := context.WithCancel(ctx)
	go func() {
		<-prwe.closeChan
		cancel()
	}()
	return prwe.wal.run(cancelCtx)
}
