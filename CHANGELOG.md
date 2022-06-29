# Changelog

## Unreleased

## 🛑 Breaking changes 🛑

- `k8sclusterreceiver`: The `receiver.k8sclusterreceiver.reportCpuMetricsAsDouble` feature gate has been removed (#10838)
  - If users were disabling this feature gate, they may have to update
    monitoring for a few Kubernetes cpu metrics. For more details see [feature-gate-configurations](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.54.0/receiver/k8sclusterreceiver#feature-gate-configurations).
- `prometheusexporter`: Automatically rename metrics with units to follow Prometheus naming convention (#8950)

### 🚩 Deprecations 🚩

- `logzioexporter`: Announcing `custom_endpoint`, `drain_interval`, `queue_capacity`, `queue_max_length` configuration options will be deprecated in upcoming releases (#10821)

### 🚀 New components 🚀

### 💡 Enhancements 💡

- `logzioexporter`: Add support for logs pipeline and support for exporterhelper (#10821)
- `flinkmetricsreceiver`: add attribute values to metadata #11520
- `prometheusreceiver`: Add `target_info` labels to resource attributes. (#11034)
- `saphanareceiver`: Fix component memory query, add better error handling (#11507)
- `sqlqueryreceiver`: Add core functionality to SQL query receiver (#10867)
- `sapmexporter`: Add config option to log responses from Splunk APM. (#11425)
- `splunkhecexporter`: Update limits for max_content_length settings (#11550)
- `filterprocessor`: Add ability to filter `Spans` (#6341)
- `prometheusexporter` : Added a feature to prometheusexporter to export exemplars along with histogram metrics (#9945)
- `tracegen`: support add additional resource attributes. (#11145)
- `transformprocessor`: Add IsMatch factory function.  This function allows regex matching in conditions (#10903)
- `transformprocessor`: `replace_pattern` and `replace_all_patterns` use regex for pattern matching and replacing text in attributes/metrics. (#11125)
- `coralogixexporter`: Add support for metrics (#11065)

### 🧰 Bug fixes 🧰

- `datadogexporter`: The `traces.span_name_remappings` setting now correctly refers to the OpenTelemetry key to be renamed without any sort of normalization. (#9693)
- `datadogexporter`: Unify traces exporter behavior with Datadog Agent OTLP traces ingest. (#9693)
- `aerospikereceiver`: Fix issue where namespaces would not be collected (#11465)
- `filestorageextension`: Copy values returned by Get (#11776)
- `redactionprocessor`: respect allow_all_keys configuration (#11542)
- `sapmreceiver`: Fix issue where component instance use in multiple pipelines leads to start failures (#11518)
- `signalfxreceiver`: Fix issue where component instance use in multiple pipelines leads to start failures (#11513)
- `splunkhecreceiver`: Fix issue where component instance use in multiple pipelines leads to start failures (#11517)

## v0.54.0

## 🛑 Breaking changes 🛑

- `transformprocessor`: `metric.is_monotonic` is now accessed via a bool literal instead of a string. (#10473)
- `vcenterreceiver`: Changed the attribute `effective` on `vcenter.cluster.host.count` as it will now be reported as a bool rather than a string (#10914)

### 🚩 Deprecations 🚩

- `datadogexporter`: Deprecate `instrumentation_library_metadata_as_tags` (#11135)
- `datadogexporter`: Deprecate `Sanitize` method of `Config` struct (#8829)
- `observiqexporter`: Deprecate the observiq exporter (#10977)
- `honeycombexporter`: Deprecate honeycomb exporter (#10318)

### 🚀 New components 🚀

- `expvarreceiver`: Include `expvarreceiver` in components (#10847)
- `googlemanagedprometheusexporter` Add the Google Managed Service for Prometheus exporter. (#10840)
- `googlemanagedprometheusexporter` The Google Managed Service for Prometheus exporter is alpha. (#10925)

### 💡 Enhancements 💡

- `tailsamplingprocessor`: Add trace_state policy (#10852)
- `mongodbatlasreceiver` Add support for receiving alerts (#10854)
- `cmd/mdatagen`: Allow attribute values of any types (#9245)
- `metricstransformprocessor`: Migrate the processor from OC to pdata (#10817)
  - This behavior can be reverted by disabling the `processor.metricstransformprocessor.UseOTLPDataModel` feature gate.
- `transformprocessor`: Add byte slice literal to the grammar.  Add new SpanID and TraceID functions that take a byte slice and return a Span/Trace ID. (#10487)
- `transformprocessor`: Add Summary transform functions. (#11041)
- `transformprocessor`: Add nil literal to the grammar. (#11150)
- `elasticsearchreceiver`: Add integration test for elasticsearch receiver (#10165)
- `tailsamplingprocessor`: New sampler added that allows to sample based on minimum number of spans
- `datadogexporter`: Some config validation and unmarshaling steps are now done on `Validate` and `Unmarshal` instead of `Sanitize` (#8829)
- `datadogexporter`: Add `exporter.datadog.hostname.preview` feature flag and related warnings (#10926)
- `datadogexporter`: Add `instrumentation_scope_metadata_as_tags` instead of `instrumentation_library_metadata_as_tags` in favor of https://github.com/open-telemetry/opentelemetry-proto/releases/tag/v0.15.0 (#11135)
- `examples`: Add an example for scraping Couchbase metrics (#10894)
- `filestorageextension`: Add background compaction capability (#9327)
- `googlecloudpubsubreceiver`: Added new `Endpoint` and `Insecure` connection configuration options. (#10845)
- `dynatraceexporter`: Provide better estimated summaries for partial histograms. (#11044)
- `mongodbreceiver`: Add integration test for mongodb receiver (#10864)
- `mezmoexporter`: add logging for HTTP errors (#10875)
- `googlecloudexporter`: Support writing to multiple GCP projects by setting the `gcp.project.id` resource attribute, and support service account impersonation (#11051)
- `k8sattributeprocessor`: Add debug logs to help identify missing attributes (#11060)
- `jmxreceiver`: Add latest releases of jmx metrics gatherer & wildfly jar to supported jars hash list (#11134)
- `rabbitmqreceiver`: Add integration test for rabbitmq receiver (#10865)
- `transformprocessor`: Allow using trace_state with key-value struct (#11029)
- `awscontainerinsightsreciever`: Pod detection Logic to support k8's on containerd runtime (#11666)


### 🧰 Bug fixes 🧰

- `kubletetstatsreceiver`: Bring back `k8s.container.name` attribute (#10848)
- `transformprocessor`: Fix issue where some metric fields were not working correctly in conditions. (#10473)
- `transformprocessor`: Fix issue where some trace fields were not working correctly in conditions. (#10471)
- `transformprocessor`: Fix issue where some log fields were not working correctly in conditions. (#10903)
- `pkg/stanza`: Skip building fingerprint in case of configuration change (#10485)
- `windowseventlogreceiver`: Fixed example config in readme (#10971)
- `pkg/stanza`: Fix access to atomic variable without using atomic package (#11023)
- `exporter/awsemfexporter:`: Fix dead links in README.md. (#11027)
- `googlecloudexporter`: Fix (self-obs) point_count metric calculation, concurrent map write panic, and dropped log attributes (#11051)
- `signalfxexporter`: Event Type is a required field, if not set, set it to `unknown` to prevent signalfx ingest from dropping it (#11121)
- `prometheusreceiver`: validate that combined metric points (e.g. histograms) have the same timestamp (#9385)
- `splunkhecexporter`: Fix flaky test when exporting traces (#11418)
- `mongodbatlasexporter`: Fix mongodbatlas.system.memory.usage.max not being reported (#11126)
- `receiver/awsxrayreceiver`: Fix null span exception fields causing null pointer exception (#11431)
- `pkg/stanza`: use ObservedTimestamp to decide if flush log for recombine operator (#11433)

## v0.53.0

### 🛑 Breaking changes 🛑

- `jmxreceiver`: Remove properties & groovyscript parameters from JMX Receiver. Add ResourceAttributes & LogLevel parameter to supply some of the removed functionality with reduced attack surface (#9685)
- `resourcedetectionprocessor`: 'gke' and 'gce' resource detectors are replaced with a single 'gcp' detector (#10347)
- `pkg/stanza`: Removed reference to deprecated `ClusterName` (#10426)
- `couchbasereceiver`: Fully removed unimplemented Couchbase receiver (#10482)
- `hostmetricsreciever`: Fix Load Scraper to normalize 1m, 5m, and 15m averages independently (#8267)

### 🚀 New components 🚀

- `flinkmetricsreceiver`: Add implementation of Flink Metric Receiver (#10121)
- `windowseventlogreceiver` Added implementation of Windows Event Log Receiver (#9228)
- `vcenterreceiver`: Add metrics receiver for new vcenterreceiver component (#9224)
- `googlecloudpubsubreceiver` Activate the Google Cloud Pubsub receiver. (#10580)
- `googlecloudpubsubexporter` Activate the Google Cloud Pubsub exporter. (#10580)
- `aerospikereceiver`: Add implementation of Aerospike Metric Receiver. (#9961)

### 💡 Enhancements 💡

- `awsemfexporter`: Add min and max support for histograms (#10577)
- `tailsamplingprocessor`: Add support for string invert matching to `and` policy (#9553)
- `mezemoexporter`: Add user agent string to outgoing HTTP requests (#10470)
- `prometheusreceiver`: Improve performance of metrics builder (#10546)
- `transformprocessor`: Add functions for conversion of scalar metric types (`gauge_to_sum` and `sum_to_gauge`) (#10255)
- `dynatraceexporter`: Use min and max when provided in a data point for histograms (#10815)
- `dynatraceexporter`: Truncate unmarshalable responses to avoid long log lines (#10568)
- `scrapertest`: Add `IgnoreResourceAttributeValue` option to metric comparison (#10828)

### 🧰 Bug fixes 🧰

- `transformprocessor`: Fix issue where incorrect error was returned if a bad path was passed to a function (#10141)
- `tanzuobservabilityexporter`: Improve how negative values in exponential histograms are handled. (#10135)
- `dynatraceexporter`: Ensure min is always less than or equal to mean and max is always greater or equal to mean for histogram estimation. (#10257)
- `resourcedetectionprocessor`: GCP resource detector now properly detects zone/region on GKE (#10347)
- `resourcedetectionprocessor`: GCP resource detector no longer fails to detect resource when using workload identity (#10486)
- `tailsamplingprocessor`: Fix composite sampler with inverse policy
- `awsprometheusremotewriteexporter`: Fix signing of empty request bodies. (#10578)
- `sigv4authextension`: Fix signing of empty request bodies. (#10578)
- `prometheusexporter`: Converting monotonic Delta to Cumulative sums (#9919)
- `statsdreceiver`: Update the lastIntervalTime for Counter metrics (#9919)
- `resourcedetectionprocessor`: GCP resource detector now correctly detects region on Google App Engine standard (#10814)
- `apachereceiver`: Update units to follow semconv (#10587)

## v0.52.0

### 🛑 Breaking changes 🛑

- `jmxreceiver`: Hash the jars provided to JMX Receiver and only allow if they match an approved list (#9687)
- `jmxreceiver`: Remove properties & groovyscript parameters from JMX Receiver. Add ResourceAttributes & LogLevel parameter to supply some of the removed functionality with reduced attack surface (#9685)

### 🚀 New components 🚀

- `aerospikereceiver`: Add implementation of Aerospike Metrics Receiver (#9961)
- `bigipreceiver`: Add implementation of F5 Big-IP Metric Receiver (#9680)
- `expvarreceiver`: Initial work for a receiver designed to scrape `memstats` from Golang applications. (#9747)
- `mezmoexporter`: Add implementation of Mezmo Log exporter (#9743)
- `nsxtreceiver`: Added implementation of NSX-T Metric Receiver (#9568)
- `expvarreceiver`: Add implementation of new receiver. (#10183)
- `telemetrygen`: Started implementing an upgraded version of `tracegen` generating traces and metrics (#9597)

### 💡 Enhancements 💡

- `transformprocessor`: Add transformation of metrics (#10100)
- `transformprocessor`: Include transform processor in components (#10134)
- `kubeletstatsreceiver`: Update receiver to use new Metrics Builder. All emitted metrics remain the same. (#9744)
- `transformprocessor`: Add new `replace_match` and `replace_all_matches` functions (#10132)
- `resourcedetectionprocessor`: Add "cname" and "lookup" hostname sources
- `jmxreceiver`: Communicate with JMX metrics gatherer subprocess via properties file (#9685)
- `pkg/stanza`: make multiline tests more like integration tests #10353 

### 🧰 Bug fixes 🧰

- `datadogexporter`: add error checks for datadog exporter (#9964)
- `datadogexporter`: Fix host aliases not being properly sent to the Datadog backend (#9748)
- `groupbyattrsprocessor`: copied aggregationtemporality when grouping metrics. (#9088)
- `jaeger`: Update OTLP-Jaeger translation of span events according to the OTel Spec: use `event` log field instead
  of `message` to represent OTel Span Event Name (#10273)
- `mongodbreceiver`: Fix issue where receiver startup could hang (#10111)
- `transformprocessor`: Fix issue where metric.aggregation_temporality and metric.is_monotic were not actually gettable or settable (#10197)
- `signalfxexporter`: Emit prometheus compatible histogram/summary to signalfx #10299
  - This behavior can be reverted using the `exporter.signalfxexporter.PrometheusCompatible` featuregate.
- `podmanreceiver`: Container Stats Error structure (#9397)
- `pkg/stanza`: pipeline.Operators() will return a consistently ordered list of operators whenever possible (#9761)
- `tanzuobservabilityexporter`: add  error checks for tanzuobservability exporter (#10188)

## v0.51.0

### 🛑 Breaking changes 🛑

- `datadogexporter`: Replace HistogramMode defined as string with enum. (#9589)
- `pkg/translator/signalfx`: Change signalfx translator to expose To/From translator structs. (#9740)
- `transformprocessor`: Add parameter validation to `truncate_all` and `limit` functions.  The `limit` parameter can no longer be negative. (#9783)
- `newrelicexporter` deleted. Use New Relic [native OTLP ingest](https://docs.newrelic.com/docs/more-integrations/open-source-telemetry-integrations/opentelemetry/opentelemetry-setup/) instead. (#9894)
- `k8sclusterreceiver`: Removing `ClusterName` as per https://github.com/kubernetes/apimachinery/commit/430b920312ca0fa10eca95967764ff08f34083a3. (#9885)

### 🚩 Deprecations 🚩

- `exporter/azuremonitor`: Deprecate use of LogRecord.Name as the log envelope category name. There is no replacement. (#9258)
- `processor/k8sattributes`: Deprecate use of k8s.cluster.name metadata parameter (obsolete) (#9968)

### 🚀 New components 🚀

- `schemaprocessor`: Starting the initial work to allow from translating from semantic convention to another (#8371)
- `saphanareceiver`: Added implementation of SAP HANA Metric Receiver (#8827)
- `logstransformprocessor`: Add implementation of Logs Transform Processor (#9335)

### 💡 Enhancements 💡

- `cmd/mdatagen`: Replace enum attributes values with typed constants (#9683)
- `elasticsearchreceiver`: Update metrics scope name from `otelcol/elasticsearch`
  to `otelcol/elasticsearchreceiver` (#9757)
- `k8sclusterreceiver`: Validate that k8s API supports a resource before setting up a watcher for it (#9523)
- `internal/stanza`: Add support for `remove` operator (#9524)
- `k8sattributesprocessor`: Support regex capture groups in tag_name (#9525)
- `mongoreceiver`: Update metrics scope name from `otelcol/mongodb` to `otelcol/mongodbreceiver` (#9759)
- `transformprocessor`: Add new `truncation` function to allow truncating string values in maps such as `attributes` or `resource.attributes` (#9546)
- `datadogexporter`: Add `api.fail_on_invalid_key` to fail fast if api key is invalid (#9426)
- `transformprocessor`: Add support for functions to validate parameters (#9563)
- `googlecloudexporter`: Add GCP cloud logging exporter (#9679)
- `transformprocessor`: Add new `limit` function to allow limiting the number of items in a map, such as the number of attributes in `attributes` or `resource.attributes` (#9552)
- `processor/attributes`: Support attributes set by server authenticator (#9420)
- `datadogexporter`: Experimental support for Exponential Histograms with delta aggregation temporality (#8350)
- `prometheusreceiver`: Support OpenMetrics Info and Stateset metrics (#9378)

### 🧰 Bug fixes 🧰

- `k8sclusterreceiver`: Fix the receiver to work with 1.19 and 1.20 k8s API versions (#9523)
- `azuremonitorexporter`: Fix log exporter bug related to incorrectly mapping SpanId (#9579)
- `mysqlreceiver`: Fix attribute values mismatch with its definition (#9688)
- `opencensusreceiver`: Do not report fatal error if err is server closed (#9559).
- `sqlserverreceiver`: Fix the receiver to have integer types on metrics where applicable (#9601)
- `prometheusreceiver`: Fix the memory issue introduced in the 0.49.0 release (#9718)
- `couchdbreceiver`: Fix issue where the receiver would not respect custom metric settings (#9598)
- `nginxreceiver`: Include nginxreceiver in components (#9572)
- `pkg/translator/prometheusremotewrite`: Fix data race when used with other exporters (#9736)
- `examples/demo`: fix baggage not work in trace demo app. (#9418)
- `prometheusreceiver`: Handle the condition where `up` metric value is NaN (#9253)
- `tanzuobservabilityexporter`: Make metrics stanza in config be optional (#9098)
- `filelogreceiver`: Update Kubernetes examples to fix native OTel logs collection issue where 0 length logs cause errors (#9754)
- `logstransformprocessor`: Resolve node ordering to fix intermittent failures (#9761)
- `awsinsightreceiver`: Migrate from `ConfigMapsResourceLock` to `ConfigMapsLeasesResourceLock` as per https://github.com/kubernetes/client-go/commit/276ea3ed979947d7cdd4b3d708862245ddcd8883 (#9885)
- `filelog`, `journald`, `syslog`, `tcplog`, `udplog`: Add support for []string type for converting log record entries (#9887)

## v0.50.0

### 🛑 Breaking changes 🛑

- `stackdriverexporter`: Remove the stackdriver exporter in favor of the identical googlecloud exporter (#9274)
- `filelog`, `journald`, `syslog`, `tcplog`, `udplog`: Remove `preserve_to` field from sub-parsers (#9331)
- `kafkametricsreceiver`: instrumentation name updated from `otelcol/kafkametrics` to `otelcol/kafkametricsreceiver` (#9406)
- `kubeletstatsreceiver`: instrumentation name updated from `kubeletstats` to `otelcol/kubeletstatsreceiver` (#9400)
- `datadogexporter`: Remove `GetHostTags` method from `TagsConfig` struct (#9423)
- `googlecloudexporter`: Graduate the `exporter.googlecloud.OTLPDirect` feature-gate to Beta.  This includes changes to the configuration structure, and many changes to default behavior. (#9471)

### 🚩 Deprecations 🚩

- `cumulativetodeltaprocessor`: Deprecated `metrics` configuration option in favor of `include` and `exclude` (#8952)
- `datadogexporter`: Deprecate `metrics::report_quantiles` in favor of `metrics::summaries::mode` (#8846)
- `datadogexporter`: Deprecate `traces.sample_rate` setting. It was never used anywhere. (#9771)

### 🚀 New components 🚀

- `iisreceiver`: Add implementation of IIS Metric Receiver (#8832)
- `sqlserverreceiver`: Add implementation of SQL Server Metric Receiver (#8398)
- `activedirectorydsreceiver`: Add implementation of Active Directory Domain Services metric receiver (#9359)
- `sqlqueryreceiver`: Add readme, factory, and config to initial implementation of SQL receiver (#9408)

### 💡 Enhancements 💡

- `pkg/translator/prometheusremotewrite`: Allow to disable sanitize metric labels (#8270)
- `basicauthextension`: Implement `configauth.ClientAuthenticator` so that the extension can also be used as HTTP client basic authenticator.(#8847)
- `azuremonitorexporter`, `lokiexporter`, `observiqexporter`: Update timestamp processing logic (#9130)
- `cumulativetodeltaprocessor`: add new include/exclude configuration options with regex support (#8952)
- `datadogexporter`: Update deprecation messages to reflect new deprecation plan (#9422)
- `cmd/mdatagen`: Update generated functions to have simple parse function to handle string parsing consistently and limit code duplication across receivers (#7574)
- `attributesprocessor`: Support filter by severity (#9132)
- `transformprocessor`: Add transformation of logs (#9368)
- `datadogexporter`: Add `metrics::summaries::mode` to specify export mode for summaries (#8846)
- `prometheusreceiver`: Add resource attributes for kubernetes resource discovery labels (#9416)

### 🧰 Bug fixes 🧰

- `fluentforwardreceiver`: Release port on shutdown (#9111)
- `prometheusexporter`: Prometheus fails to generate logs when prometheus exporter produced a check exception occurs. (#8949)
- `resourcedetectionprocessor`: Wire docker detector (#9372)
- `kafkametricsreceiver`: The kafkametricsreceiver was changed to connect to kafka during scrape, rather than startup. If kafka is unavailable the receiver will attempt to connect during subsequent scrapes until succcessful (#8817).
- `datadogexporter`: Update Kubernetes example manifest to new executable name. (#9425).
- `riakreceiver`: Fix issue where user configured metric settings were ignored. (#9561)
- `sqlserverreceiver`: Update `sqlserver.transaction_log.growth.count` and `sqlserver.transaction_log.shrink.count` to be monotonic sums. (#9522)

## v0.49.0

### ⚠️ Warning  ⚠️

This release contains an issue in
[Prometheus receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver)
causing 30% memory consumption increase when there is a lot of target churn. The issue is currently being
investigated and will be fixed in one of the new releases. More details:
https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/9278.

### 🛑 Breaking changes 🛑

- `filelogreceiver`, `journaldreceiver`, `syslogreceiver`, `tcplogreceiver`, `udplogreceiver`:
  - Updated data model to align with stable logs data model, which includes various breaking changes. (#9139, #8835)
    - A detailed [Upgrade Guide](https://github.com/open-telemetry/opentelemetry-log-collection/releases/tag/v0.28.0) is available in the log-collection v0.29.0 release notes.
- `datadogexporter`: Remove `OnlyMetadata` method from `Config` struct (#8980)
- `datadogexporter`: Remove `GetCensoredKey` method from `APIConfig` struct (#8980)
- `mongodbatlasreceiver`: Updated to uses newer metric builder which changed some metric and resource attributes (#9093)
- `dynatraceexporter`: Make `serialization` package `/internal` (#9097)
- `attributesprocessor`: Remove log names from filters (#9131)
- `k8sclusterreceiver`: The `receiver.k8sclusterreceiver.reportCpuMetricsAsDouble` feature gate is now enabled by default (#9367)
  - Users may have to update monitoring for a few Kubernetes cpu metrics, for
    more details see [feature-gate-configurations](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.54.0/receiver/k8sclusterreceiver#feature-gate-configurations).

### 🚩 Deprecations 🚩

- `datadogexporter`: Deprecate `service` setting in favor of `service.name` semantic convention (#8784)
- `datadogexporter`: Deprecate `version` setting in favor of `service.version` semantic convention (#8784)
- `datadogexporter`: Deprecate `env` setting in favor of `deployment.environment` semantic convention (#9017)
- `datadogexporter`: Deprecate `GetHostTags` method from `TagsConfig` struct (#8975)
- `datadogexporter`: Deprecate `tags` setting in favor of `host_metadata::tags` (#9100)
- `datadogexporter`: Deprecate `send_metadata` setting in favor of `host_metadata::enabled` (#9100)
- `datadogexporter`: Deprecate `use_resource_metadata` setting in favor of `host_metadata::hostname_source` (#9100)
- `prometheusexecreceiver`: Deprecate prom_exec receiver (#9058)
- `fluentbitextension`: Deprecate Fluentbit extension (#9062)

### 🚀 New components 🚀

- `riakreceiver`: Riak Metric Receiver (#8548)

### 💡 Enhancements 💡
- `splunkhecexporter`: Add support for batching traces (#8995)
- `hostmetricsreceiver`: Migrate Processes scraper to the Metrics builder (#8855)
- `tanzuobservabilityexporter`: Use resourcetotelemetry helper (#8338)
- Add `make crosslink` target to ensure replace statements are included in `go.mod` for all transitive dependencies within repository (#8822)
- `filestorageextension`: Change bbolt DB settings for better performance (#9004)
- `jaegerremotesamplingextension`: Add local and remote sampling stores (#8818)
- `attributesprocessor`: Add support to filter on log body (#8996)
- `prometheusremotewriteexporter`: Translate resource attributes to the target info metric (#8493)
- `prometheusexporter`: Add `job` and `instance` labels to metrics so they can be scraped with `honor_labels: true` (#9115)
- `podmanreceiver`: Add API timeout configuration option (#9014)
- `cmd/mdatagen`: Add `sem_conv_version` field to metadata.yaml that is used to set metrics SchemaURL (#9010)
- `splunkheceporter`: Add an option to disable log or profiling data (#9065)
- `windowsperfcountersreceiver`: Move code into separate package for use in other windowsperfcounter receivers (#9108)
- `datadogexporter`: Add `host_metadata` configuration section to configure host metadata export (#9100)
- `cmd/mdatagen`: Update documentation generated for attributes to list enumerated values and show the "value" that will be visible on metrics when it is different from the attribute key in metadata.yaml (#8983)
- `routingprocessor`: add option to drop resource attribute used for routing (#8990)

### 🧰 Bug fixes 🧰

- `filestorageextension`: use correct bbolt options for compaction (#9134)
- `hostmetricsreceiver`: Use cpu times for time delta in cpu.utilization calculation (#8857)
- `dynatraceexporter`: Remove overly verbose stacktrace from certain logs (#8989)
- `googlecloudexporter`: fix the `exporter.googlecloud.OTLPDirect` fature-gate, which was not applied when the flag was provided (#9116)
- `signalfxexporter`: Fix bug to enable timeouts for correlating traces and metrics (#9101)
- `windowsperfcountersreceiver`: fix exported values being integers instead of doubles (#9138)
- `prometheusreceiver`: Fix issues with relabelling the `job` and `instance` labels. (#8780)
- `dynatraceexporter`: Continue processing data points after a serialization error. (#9330)

## v0.48.0

### 💡 Enhancements 💡

- `k8seventsreceiver`: Add Api_version and resource_version (#8539)
- `datadogexporter`: Add `metrics::sums::cumulative_monotonic_mode` to specify export mode for cumulative monotonic sums (#8490)
- `dynatraceexporter`: add multi-instance deployment note to README.md (#8848)
- `resourcedetectionprocessor`: Add attribute allowlist (#8547)
- `datadogexporter`:  Metrics payload data and Sketches payload data will be logged if collector is started in debug mode (#8929)
- `cmd/mdatagen`: Add resource attributes definition to metadata.yaml and move `pdata.Metrics` creation to the
  generated code (#8555)

### 🛑 Breaking changes 🛑

- `windowsperfcountersreceiver`: Added metrics configuration (#8376)
- `lokiexporter`: Remove deprecated LogRecord.name field (#8951)
- `splunkhecexporter`: Remove deprecated LogRecord.name field (#8951)

### 🚩 Deprecations 🚩

- `datadogexporter`: Deprecate `OnlyMetadata` method from `Config` struct (#8359)
- `datadogexporter`: Deprecate `GetCensoredKey` method from `APIConfig` struct (#8830)
- `datadogexporter`: Deprecate `metrics::send_monotonic_counter` in favor of `metrics::sums::cumulative_monotonic_mode` (#8490)

### 🚀 New components 🚀

- `sigv4authextension`: Enable component (#8518)

## v0.47.0

### 💡 Enhancements 💡

- `googlecloudexporter`: Add Validate method in config (#8559)
- `attributesprocessor`: Add convert action (#7930)
- `attributesprocessor`: Add metric support (#8111)
- `prometheusremotewriteexporter`: Write-Ahead Log support enabled (#7304)
- `hostreceiver/filesystemscraper`: Add filesystem utilization (#8027)
- `hostreceiver/pagingscraper`: Add paging.utilization (#6221)
- `googlecloudexporter`: [Alpha] Translate metrics directly from OTLP to gcm using the `exporter.googlecloud.OTLPDirect` feature-gate (#7177)
- `simpleprometheusreceiver`: Add support for static labels (#7908)
- `spanmetricsprocessor`: Dropping the condition to replace _ with key_ as __ label is reserved and _ is not (#8057)
- `podmanreceiver`: Add container.runtime attribute to container metrics (#8262)
- `dockerstatsreceiver`: Add container.runtime attribute to container metrics (#8261)
- `tanzuobservabilityexporter`: instrumentation Library and Dropped Counts to Span Tags (#8120)
- `clickhouseexporter`: Implement consume log logic. (#9705)
- `influxdbexporter`: Add support for cumulative, non-monotonic metrics. (#8348)
- `oauth2clientauthextension`: Add support for EndpointParams (#7307)
- Add `NewMetricData` function to `MetricsBuilder` to consistently set instrumentation library name (#8255)
- `googlecloudpubsubreceiver` Added implementation of Google Cloud Pubsub receiver. (#8391)
- `googlecloudpubsubexporter` Added implementation of Google Cloud Pubsub exporter. (#8391)
- `coralogixexporter` Allow exporter timeout to be configured (#7957)
- `prometheusremotewriteexporter` support adding trace id and span id attached to exemplars (#8380)
- `influxdbexporter`: accept histogram metric missing infinity bucket. (#8462)
- `skywalkingreceiver`: Added implementation of Skywalking receiver. (#8549)
- `prometheusreceiver`: Fix staleness bug for histograms and summaries (#8561)

### 🛑 Breaking changes 🛑

- `mongodbatlasreceiver`: rename mislabeled attribute `memory_state` to correct `disk_status` on partition disk metrics (#7747)
- `mongodbatlasreceiver`: Correctly set initial lookback for querying mongodb atlas api (#8246)
- `nginxreceiver`: instrumentation name updated from `otelcol/nginx` to `otelcol/nginxreceiver` (#8255)
- `postgresqlreceiver`: instrumentation name updated from `otelcol/postgresql` to `otelcol/postgresqlreceiver` (#8255)
- `redisreceiver`: instrumentation name updated from `otelcol/redis` to `otelcol/redisreceiver` (#8255)
- `apachereceiver`: instrumentation name updated from `otelcol/apache` to `otelcol/apachereceiver` ()
- `couchdbreceiver`: instrumentation name updated from `otelcol/couchdb` to `otelcol/couchdbreceiver` (#8366)
- `prometheusreceiver` Change resource attributes on metrics: `instance` -> `service.instance.id`, `host.name` -> `net.host.name`,  `port` -> `net.host.port`, `scheme` -> `http.scheme`, `job` removed (#8266)
- `prometheusremotewriteexporter` Use `service.*` resource attributes instead of `job` and `instance` resource attributes when adding job and instance labels to metrics (#8266)
- `mysqlreceiver`: instrumentation name updated from `otel/mysql` to `otelcol/mysqlreceiver` (#8387)
- `zookeeperreceiver`: instrumentation name updated from `otelcol/zookeeper` to `otelcol/zookeeperreceiver` (#8389)
- `coralogixexporter`: Create dynamic subsystem name (#7957)
  - Deprecate configuration changed. Dynamic subsystem name from traces service name property.
- `rabbitmqreceiver`: instrumentation name updated from `otelcol/rabbitmq` to `otelcol/rabbitmqreceiver` (#8400)

### 🧰 Bug fixes 🧰

- `zipkinexporter`: Set "error" tag value when status is set to error (#8187)
- `prometheusremotewriteexporter`: Correctly handle metric labels which collide after sanitization (#8378)
- `prometheusremotewriteexporter`: Drop labels when exemplar attributes exceed the max number of characters (#8379)
- `k8sclusterreceiver`: Add support to enable k8s node and container cpu metrics to be reported as double values (#8245)
  - Use "--feature-gates=receiver.k8sclusterreceiver.reportCpuMetricsAsDouble" to enable reporting node and container
    cpu metrics as a double values.
- `tanzuobservabilityexporter`: Fix a typo in Instrumentation Library name and version tags (#8384)
- `logreceivers`: Fix an issue where receiver would sometimes fail to build using Go 1.18 (#8521)
- `awsxrayreceiver`: Add defaults for optional stack frame parameters (#8790)

### 🚩 Deprecations 🚩

- `datadogexporter`: Deprecate automatic environment variable detection (#8397)

### 🚀 New components 🚀
- `sigv4authextension`: New Component: Sigv4 Authenticator Extension (#8263)

## v0.46.0

### 💡 Enhancements 💡

- `internal/stanza`: Export metrics from Stanza receivers (#8025)
- `hostreceiver/pagingscraper`: Migrate the scraper to the mdatagen metrics builder (#7139)
- Do not drop zero trace/span id spans in the jaeger conversion (#7946)
- Upgrade to use semantic conventions 1.6.1 (#7926)
- `dynatraceexporter`: Validate QueueSettings and perform config validation in Validate() instead (#8020)
- `sapmexporter`: Add validation for `sending_queue` setting (#8023)
- `signalfxexporter`: Add validation for `sending_queue` setting (#8026)
- `internal/stanza`: Add support for arbitrary attribute types (#8081)
- `resourcedetectionprocessor`: Add confighttp.HTTPClientSettings To Resource Detection Config Fixes (#7397)
- `hostmetricsreceiver`: Add cpu.utilization metrics to cpu scrapper (#7130)
- `honeycombexporter`: Add validation for `sending_queue` setting (#8113)
- `routingprocessor`: Expand error handling on failure to build exporters (#8125)
- `skywalkingreceiver`: Add new skywalking receiver component folder and structure (#8107)
- `groupbyattrsprocesor`: Allow empty keys, which allows to use the processor for compaction (#7793)
- `datadogexporter`: Add rbac to example k8s manifest file (#8186)
- `splunkhecexporter`: Add validation for `sending_queue` setting (#8256)

### 🛑 Breaking changes 🛑

- Remove deprecated functions from jaeger translator (#8032)
- `internal/stanza`: Remove `write_to` setting from input operators (#8081)
- `mongodbatlasreceiver`: rename `mongodb.atlas.*` attributes to `mongodb_atlas.*` adhering to naming guidelines. Adding 3 new attributes (#7960)

### 🧰 Bug fixes 🧰

- `prometheusreceiver`: Fix segfault that can occur after receiving stale metrics (#8056)
- `filelogreceiver`: Fix issue where logs could occasionally be duplicated (#8123)
- `prometheusremotewriteexporter`: Fix empty non-string resource attributes (#8116)

### 🚀 New components 🚀

## v0.45.1

### 💡 Enhancements 💡

- `sumologicexporter`: Move validation to Config (#7936)
- `elasticsearchexporter`: Fix crash with batch processor (#7953).
- `splunkhecexporter`: Batch metrics payloads (#7760)
- `tanzuobservabilityexporter`: Add internal SDK metric tag (#7826)
- `hostreceiver/processscraper`: Migrate the scraper to the mdatagen metrics builder (#7287)

### 🧰 Bug fixes 🧰

- `awsprometheusremotewriteexporter`: fix dependencies issue (#7963)

### 🚀 New components 🚀

- `awsfirehose` receiver: Add AWS Kinesis Data Firehose Receiver (#7918)

## v0.45.0

### 💡 Enhancements 💡

- `hostreceiver/filesystemscraper`: Migrate the scraper to the mdatagen metrics builder (#7772)
- `hostreceiver/memoryscraper`: Migrate the scraper to the mdatagen metrics builder (#7312)
- `lokiexporter`: Use record attributes as log labels (#7569)
- `routingprocessor`: Do not err on failure to build exporters (#7423)
- `apachereceiver`: Update to mdatagen v2 (#7573)
- `datadogexporter`: Don't send host metadata if hostname is empty (#7426)
- `datadogexporter`: Add insecure_skip_verify flag to configuration (#7422)
- `coralogixexporter`: Update readme (#7785)
- `awscloudwatchlogsexporter`: Remove name from aws cloudwatch logs exporter (#7554)
- `tanzuobservabilityexporter`: Update OTel Collector's Exporter to match WF Proxy Handling of source (#7929)
- `hostreceiver/memoryscraper`: Add memory.utilization (#6221)
- `awskinesisexporter`: Add Queue Config Validation AWS Kinesis Exporter (#7835)
- `elasticsearchexporter`: Remove usage of deprecated LogRecord.Name field (#7829).
- `loadbalancingexporter`: Allow non-exist hostname on startup (#7935)
- `datadogexporter`: Use exact sum, count and average on Datadog distributions (#7830)
- `storage/filestorage`: add optional compaction to filestorage (#7768)
- `tanzuobservabilityexporter`: Add attributes from the Resource to the resulting WF metric tags & set `source` value in WF metric (#8101)

### 🛑 Breaking changes 🛑

- Use go mod compat, drops support for reproducibility with go 1.16 (#7915)
- `apachereceiver`: Update instrumentation library name from `otel/apache` to `otelcol/apache` (#7754)
- `pkg/translator/prometheusremotewrite`: Cleanup prw translator public functions (#7776)
- `prometheusreceiver`: The OpenCensus-based metric conversion pipeline has
  been removed.
  - The `receiver.prometheus.OTLPDirect` feature gate has been removed as
    the direct pipeline is the only remaining pipeline.
- `translator/jaeger`: Cleanup jaeger translator function names (#7775)
  - Deprecate old funcs with Internal word.
- `mysqlreceiver`: Update data model and names for several metrics (#7924)
  - Change all metrics to Int values
  - Remove `mysql.buffer_pool_pages`. Replace with:
    - `mysql.buffer_pool.pages`
    - `mysql.buffer_pool.data_pages`
    - `mysql.buffer_pool.page_flushes`
  - Remove `mysql.buffer_pool_size`. Replace with:
    - `mysql.buffer_pool.limit`
    - `mysql.buffer_pool.usage`
  - Rename `mysql.buffer_pool_operations` to `mysql.buffer_pool.operations`

### 🚩 Deprecations 🚩

- Deprecated log_names setting from filter processor. (#7552)

### 🧰 Bug fixes 🧰

 - `tailsamplingprocessor`: "And" policy only works as a sub policy under a composite policy (#7590)
 - `prometheusreceiver`: Correctly map description and units when converting
  Prometheus metadata directly to pdata. (#7748)
 - `sumologicexporter`: fix exporter panics on malformed histogram (#7548)
- `awsecscontainermetrics`: CPU Reserved is now 1024/vCPU for ECS Container Insights (#6734)

### 🚀 New components 🚀

- `clickhouse` exporter: Add ClickHouse Exporter (#6907)
- `pkg/translator/signalfx`: Extract signalfx to metrics conversion in a separate package (#7778)
  - Extract FromMetrics to SignalFx translator package (#7823)

## v0.44.0

### 💡 Enhancements 💡

- `kafkaexporter`: Add compression and flush max messages options.
- `dynatraceexporter`: Write error logs using plugin logger (#7360)
- `dynatraceexporter`: Fix docs for TLS settings (#7568)
- `tanzuobservabilityexporter`: Turn on metrics exporter (#7281)
- `attributesprocessor` `resourceprocessor`: Add `from_context` value source
- `resourcedetectionprocessor`: check cluster config to verify resource is on aws for eks resources (#7186)
- `awscloudwatchlogsexporter`: enable awscloudwatchlogsexporter which accepts and exports log data (#7297)
- `translator/prometheusremotewrite`: add a new module to help translate data from OTLP to Prometheus Remote Write (#7240)
- `azuremonitorexporter`: In addition to traces, export logs to Azure Application Insights (#7403)
- `jmxreceiver`: Added `additional_jars` configuration option to launch JMX Metric Gatherer JAR with extended `CLASSPATH` (#7378)
- `awscontainerinsightreceiver`: add full pod name when configured to AWS Container Insights Receiver (#7415)
- `hostreceiver/loadscraper`: Migrate the scraper to the mdatagen metrics builder (#7288)
- `awsecscontainermetricsreceiver`: Rename attributes to follow semantic conventions (#7425)
- `datadogexporter`: Always map conventional attributes to tags (#7185)
- `mysqlreceiver`: Add golden files for integration test (#7303)
- `nginxreceiver`: Standardize integration test (#7515)
- `mysqlreceiver`: Update to use mdatagen v2 (#7507)
- `postgresqlreceiver`: Add integration tests (#7501)
- `apachereceiver`: Add integration test (#7517)
- `mysqlreceiver`: Use scrapererror to report errors (#7513)
- `postgresreceiver`: Update to mdatagen v2 (#7503)
- `nginxreceiver`: Update to mdatagen v2 (#7549)
- `datadogexporter`: Fix traces exporter's initialization log (#7564)
- `tailsamplingprocessor`: Add And sampling policy (#6910)
- `coralogixexporter`: Add Coralogix Exporter (#7383)
- `prometheusexecreceiver`: Add default value for `scrape_timeout` option (#7587)

### 🛑 Breaking changes 🛑

- `resourcedetectionprocessor`: Update `os.type` attribute values according to semantic conventions (#7544)
- `awsprometheusremotewriteexporter`: Deprecation notice; may be removed after v0.49.0
  - Switch to using the `prometheusremotewriteexporter` + `sigv4authextension` instead

### 🧰 Bug fixes 🧰

- `resourcedetectionprocessor`: fix `meta` allow list excluding keys with nil values (#7424)
- `postgresqlreceiver`: Fix issue where empty metrics could be returned after failed connection (#7502)
- `resourcetotelemetry`: Ensure resource attributes are added to summary
  and exponential histogram data points. (#7523)

### 🚩 Deprecations 🚩

- Deprecated otel_to_hec_fields.name setting from splunkhec exporter. (#7560)

## v0.43.0

### 💡 Enhancements 💡

- `coralogixexporter`: First implementation of Coralogix Exporter (#6816)
- `cloudfoundryreceiver`: Enable Cloud Foundry client (#7060)
- `elasticsearchexporter`: add elasticsearchexporter to the components exporter list (#6002)
- `elasticsearchreceiver`: Add metric metadata (#6892)
- `elasticsearchreceiver`: Use same metrics as JMX receiver for JVM metrics (#7160)
- `elasticsearchreceiver`: Implement scraping logic (#7174)
- `datadogexporter`: Add http.status_code tag to trace stats (#6889)
- `datadogexporter`: Add configuration option to use OTel span name into the Datatog resource name (#6611)
- `mongodbreceiver`: Add initial client code to the component (#7125)
- `tanzuobservabilityexporter`: Support delta histograms (#6897)
- `awscloudwatchlogsexporter`: Use cwlogs package to export logs (#7152)
- `mysqlreceiver`: Add the receiver to available components (#7078)
- `tanzuobservabilityexporter`: Documentation for the memory_limiter configuration (#7164)
- `dynatraceexporter`: Do not shut down exporter when metrics ingest module is temporarily unavailable (#7161)
- `mongodbreceiver`: Add metric metadata (#7163)
- `mongodbreceiver`: Add metric scraping (#7175)
- `postgresqlreceiver`: add the receiver to available components (#7079)
- `rabbitmqreceiver`: Add scraper logic (#7299)
- `tanzuobservability exporter`: Support summary metrics (#7121)
- `mongodbatlasreceiver`: Add retry and backoff to HTTP client (#6943)
- Use Jaeger gRPC instead of Thrift in the docker-compose example (#7243)
- `tanzuobservabilityexporter`: Support exponential histograms (#7127)
- `receiver_creator`: Log added and removed endpoint env structs (#7248)
- `prometheusreceiver`: Use the OTLP data conversion path by default. (#7282)
  - Use `--feature-gates=-receiver.prometheus.OTLPDirect` to re-enable the
    OpenCensus conversion path.
- `extension/observers`: Correctly set image and tag on container endpoints (#7279)
- `tanzuobservabilityexporter`: Document how to enable memory_limiter (#7286)
- `hostreceiver/networkscraper`: Migrate the scraper to the mdatagen metrics builder (#7048)
- `hostmetricsreceiver`: Add MuteProcessNameError config flag to mute specific error reading process executable (#7176)
- `scrapertest`: Improve comparison logic (#7305)
- `hostmetricsreceiver`: add `cpu_average` option for load scraper to report the average cpu load (#6999)
- `scrapertest`: Add comparison option to ignore specific attributes (#6519)
- `tracegen`: Add option to pass in custom headers to export calls via command line (#7308)
- `tracegen`: Provide official container images (#7179)
- `scrapertest`: Add comparison function for pdata.Metrics (#7400)
- `prometheusremotewriteexporter` : Dropping the condition to replace _ with key_ as __ label is reserved and _ is not (#7112)

### 🛑 Breaking changes 🛑

- `tanzuobservabilityexporter`: Remove status.code
- `tanzuobservabilityexporter`: Use semantic conventions for status.message (#7126)
- `k8sattributesprocessor`: Move `kube` and `observability` packages to `internal` folder (#7159)
- `k8sattributesprocessor`: Unexport processor `Option`s (#7311)
- `zookeeperreceiver`: Refactored metrics to have correct units, types, and combined some metrics via attributes. (#7280)
- `prometheusremotewriteexporter`: `PRWExporter` struct and `NewPRWExporter()`
  function are now unexported. (#TBD)
- `newrelicexporter` marked as deprecated (#7284)

### 🚀 New components 🚀

- `rabbitmqreceiver`: Establish codebase for RabbitMQ metrics receiver (#7239)
- Add `basicauth` extension (#7167)
- `k8seventsreceiver`: Implement core logic (#6885)

### 🧰 Bug fixes 🧰

- `k8sattributeprocessor`: Parse IP out of net.Addr to correctly tag k8s.pod.ip (#7077)
- `k8sattributeprocessor`: Process IP correctly for net.Addr instances that are not typed (#7133)
- `mdatagen`: Fix validation of `enabled` field in metadata.yaml (#7166)
- `elasticsearch`: Fix timestamp for each metric being startup time (#7255)
- `prometheusremotewriteexporter`: Fix index out of range panic caused by expiring metrics (#7149)
- `resourcedetection`: Log the error when checking for ec2metadata availability (#7296)

## v0.42.0

### 💡 Enhancements 💡

- `couchbasereceiver`: Add couchbase client (#7122)
- `couchdbreceiver`: Add couchdb scraper (#7131)
- `couchdbreceiver`: Add couchdb client (#6880)
- `elasticsearchreceiver`: Implement scraper client (#7019)
- `couchdbreceiver`: Add metadata metrics (#6878)
- `prometheusremotewriteexporter`: Handling Staleness flag from OTLP (#6679)
- `prometheusexporter`: Handling Staleness flag from OTLP (#6805)
- `prometheusreceiver`: Set OTLP no-data-present flag for stale scraped metrics. (#7043)
- `mysqlreceiver`: Add Integration test (#6916)
- `datadogexporter`: Add compatibility with ECS Fargate semantic conventions (#6670)
- `k8s_observer`: discover k8s.node endpoints (#6820)
- `redisreceiver`: Add missing description fields to keyspace metrics (#6940)
- `redisreceiver`: Set start timestamp uniformly for gauge and sum metrics (#6941)
- `kafkaexporter`: Allow controlling Kafka acknowledgment behaviour  (#6301)
- `lokiexporter`: Log the first part of the http body on failed pushes to loki (#6946)
- `resourcedetectionprocessor`: add the [consul](https://www.consul.io/) detector (#6382)
- `awsemfexporter`: refactor cw_client logic into separate `cwlogs` package (#7072)
- `prometheusexporter`: Dropping the condition to replace _ with key_ as __ label is reserved and _ is not (#7506)

### 🛑 Breaking changes 🛑

- `memcachedreceiver`: Update metric names (#6594)
- `memcachedreceiver`: Fix some metric units and value types (#6895)
- `sapm` receiver: Use Jaeger status values instead of OpenCensus (#6682)
- `jaeger` receiver/exporter: Parse/set Jaeger status with OTel spec values (#6682)
- `awsecscontainermetricsreceiver`: remove tag from `container.image.name` (#6436)
- `k8sclusterreceiver`: remove tag from `container.image.name` (#6436)

### 🚀 New components 🚀

- `ecs_task_observer`: Discover running containers in AWS ECS tasks (#6894)
- `mongodbreceiver`: Establish codebase for MongoDB metrics receiver (#6972)
- `couchbasereceiver`: Establish codebase for Couchbase metrics receiver (#7046)
- `dbstorage`: New experimental dbstorage extension (#7061)
- `redactionprocessor`: Remove sensitive data from traces (#6495)

### 🧰 Bug fixes 🧰

- `ecstaskobserver`: Fix "Incorrect conversion between integer types" security issue (#6939)
- Fix typo in "direction" metrics attribute description (#6949)
- `zookeeperreceiver`: Fix issue where receiver could panic during shutdown (#7020)
- `prometheusreceiver`: Fix metadata fetching when metrics differ by trimmable suffixes (#6932)
- Sanitize URLs being logged (#7021)
- `prometheusreceiver`: Fix start time tracking for long scrape intervals (#7053)
- `signalfxexporter`: Don't use syscall to avoid compilation errors on some platforms (#7062)
- `tailsamplingprocessor`: Add support for new policies as composite sub-policies (#6975)

### 💡 Enhancements 💡

- `lokiexporter`: add complete log record to body (#6619)
- `k8sclusterreceiver` add `container.image.tag` attribute (#6436)
- `spanmetricproccessor`: use an LRU cache for the cached Dimensions key-value pairs (#2179)
- `skywalkingexporter`: add skywalking metrics exporter (#6528)
- `deltatorateprocessor`: add int counter support (#6982)
- `filestorageextension`: document default values (#7022)
- `redisreceiver`: Migrate the scraper to the mdatagen metrics builder (#6938)

## v0.41.0

### 🛑 Breaking changes 🛑

- None

### 🚀 New components 🚀

- `asapauthextension` (#6627)
- `mongodbatlasreceiver` (#6367)

### 🧰 Bug fixes 🧰

- `filestorageextension`: fix panic when configured directory cannot be accessed (#6103)
- `hostmetricsreceiver`: fix set of attributes for system.cpu.time metric (#6422)
- `k8sobserver`: only record pod endpoints for running pods (#5878)
- `mongodbatlasreceiver`: fix attributes fields in metadata.yaml (#6440)
- `prometheusexecreceiver`: command line processing on Windows (#6145)
- `spanmetricsprocessor`: fix exemplars support (#6140)
-  Remap arm64 to aarch64 on rpm/deb packages (#6635)

### 💡 Enhancements 💡

- `datadogexporter`: do not use attribute localhost-like hostnames (#6477)
- `datadogexporter`: retry per network call (#6412)
- `datadogexporter`: take hostname into account for cache (#6223)
- `exporter/lokiexporter`: adding a feature for loki exporter to encode JSON for log entry (#5846)
- `googlecloudspannerreceiver`: added fallback to ADC for database connections. (#6629)
- `googlecloudspannerreceiver`: added parsing only distinct items for sample lock request label. (#6514)
- `googlecloudspannerreceiver`: added request tag label to metadata config for top query stats. (#6475)
- `googlecloudspannerreceiver`: added sample lock requests label to the top lock stats metrics. (#6466)
- `googlecloudspannerreceiver`: added transaction tag label to metadata config for top transaction stats. (#6433)
- `groupbyattrsprocessor`: added support for metrics signal (#6248)
- `hostmetricsreceiver`: ensure SchemaURL is set (#6482)
- `kubeletstatsreceiver`: add support for read-only kubelet endpoint (#6488)
- `mysqlreceiver`: enable native authentication (#6628)
- `mysqlreceiver`: remove requirement for password on MySQL (#6479)
- `receiver/prometheusreceiver`: do not add host.name to metrics from localhost/unspecified targets (#6476)
- `spanmetricsprocessor`: add setStatus operation (#5886)
- `splunkhecexporter`: remove duplication of host.name attribute (#6527)
- `tanzuobservabilityexporter`: add consumer for sum metrics. (#6385)
- Update log-collection library to v0.23.0 (#6593)

## v0.40.0

### 🛑 Breaking changes 🛑

- `tencentcloudlogserviceexporter`: change `Endpoint` to `Region` to simplify configuration (#6135)

### 🚀 New components 🚀

- Add `memcached` receiver (#5839)

### 🧰 Bug fixes 🧰

- Fix token passthrough for HEC (#5435)
- `datadogexporter`: Fix missing resource attributes default mapping when resource_attributes_as_tags: false (#6359)
- `tanzuobservabilityexporter`: Log and report missing metric values. (#5835)
- `mongodbatlasreceiver`: Fix metrics metadata (#6395)

### 💡 Enhancements 💡

- `awsprometheusremotewrite` exporter: Improve error message when failing to sign request
- `mongodbatlas`: add metrics (#5921)
- `healthcheckextension`: Add path option (#6111)
- Set unprivileged user to container image (#6380)
- `k8sclusterreceiver`: Add allocatable type of metrics (#6113)
- `observiqexporter`: Allow Dialer timeout to be configured (#5906)
- `routingprocessor`: remove broken debug log fields (#6373)
- `prometheusremotewriteexporter`: Add exemplars support (#5578)
- `fluentforwardreceiver`: Convert attributes with nil value to AttributeValueTypeEmpty (#6630)

## v0.39.0

### 🛑 Breaking changes 🛑

- `httpdreceiver` renamed to `apachereceiver` to match industry standards (#6207)
- `tencentcloudlogserviceexporter` change `Endpoint` to `Region` to simplify configuration (#6135)

### 🚀 New components 🚀

- Add `postgresqlreceiver` config and factory (#6153)
- Add TencentCloud LogService exporter `tencentcloudlogserviceexporter` (#5722)
- Restore `jaegerthrifthttpexporter` (#5666)
- Add `skywalkingexporter` (#5690, #6114)

### 🧰 Bug fixes 🧰

- `datadogexporter`: Improve cumulative metrics reset detection using `StartTimestamp` (#6120)
- `mysqlreceiver`: Address issues in shutdown function (#6239)
- `tailsamplingprocessor`: End go routines during shutdown (#5693)
- `googlecloudexporter`: Update google cloud exporter to correctly close the metric exporter (#5990)
- `statsdreceiver`: Fix the summary point calculation (#6155)
- `datadogexporter` Correct default value for `send_count_sum_metrics` (#6130)

### 💡 Enhancements 💡

- `datadogexporter`: Increase default timeout to 15 seconds (#6131)
- `googlecloudspannerreceiver`: Added metrics cardinality handling for Google Cloud Spanner receiver (#5981, #6148, #6229)
- `mysqlreceiver`: Mysql add support for different protocols (#6138)
- `bearertokenauthextension`: Added support of Bearer Auth for HTTP Exporters (#5962)
- `awsxrayexporter`: Fallback to rpc.method for segment operation when aws.operation missing (#6231)
- `healthcheckextension`: Add new health check feature for collector pipeline (#5643)
- `datadogexporter`: Always add current hostname (#5967)
- `k8sattributesprocessor`: Add code to fetch all annotations and labels by specifying key regex (#5780)
- `datadogexporter`: Do not rely on collector to resolve envvar when possible to resolve them (#6122)
- `datadogexporter`: Add container tags to attributes package (#6086)
- `datadogexporter`: Preserve original TraceID (#6158)
- `prometheusreceiver`: Enhance prometheus receiver logger to determine errors, test real e2e usage (#5870)
- `awsxrayexporter`: Added support for AWS AppRunner origin (#6141)

## v0.38.0

### 🛑 Breaking changes 🛑

- `datadogexporter` Make distributions the default histogram export option. (#5885)
- `redisreceiver` Update Redis receiver's metric names. (#5837)
- Remove `scraperhelper` from contrib, use the core version. (#5826)

### 🚀 New components 🚀

- `googlecloudspannerreceiver` Added implementation of Google Cloud Spanner receiver. (#5727)
- `awsxrayproxy` Wire up awsxrayproxy extension. (#5747)
- `awscontainerinsightreceiver` Enable AWS Container Insight receiver. (#5960)

### 🧰 Bug fixes 🧰

- `statsdreceiver`: fix start timestamp / temporality for counters. (#5714)
- Fix security issue related to github.com/tidwall/gjson. (#5936)
- `datadogexporter` Fix cumulative histogram handling in distributions mode (#5867)
- `datadogexporter` Skip nil sketches (#5925)

### 💡 Enhancements 💡

- Extend `kafkareceiver` configuration capabilities. (#5677)
- Convert `mongodbatlas` receiver to use scraperhelper. (#5827)
- Convert `dockerstats` receiver to use scraperhelper. (#5825)
- Convert `podman` receiver to use scraperhelper. (#5822)
- Convert `redisreceiver` to use scraperhelper. (#5796)
- Convert `kubeletstats` receiver to use scraperhelper. (#5821)
- `googlecloudspannerreceiver` Migrated Google Cloud Spanner receiver to scraper approach. (#5868)
- `datadogexporter` Use a `Consumer` interface for decoupling from zorkian's package. (#5315)
- `mdatagen` - Add support for extended metric descriptions (#5688)
- `signalfxexporter` Log datapoints option. (#5689)
- `cumulativetodeltaprocessor`: Update cumulative to delta. (#5772)
- Update configuration default values in log receivers docs. (#5840)
- `fluentforwardreceiver`: support more complex fluent-bit objects. (#5676)
- `datadogexporter` Remove spammy logging. (#5856)
- `datadogexporter` Remove obsolete report_buckets config. (#5858)
- Improve performance of metric expression matcher. (#5864)
- `tanzuobservabilityexporter` Introduce metricsConsumer and gaugeMetricConsumer. (#5426)
- `awsxrayexporter` rpc.system has priority to determine aws namespace. (#5833)
- `tailsamplingprocessor` Add support for composite sampling policy to the tailsampler. (#4958)
- `kafkaexporter` Add support for AWS_MSK_IAM SASL Auth (#5763)
- Refactor the client Authenticators  for the new "ClientAuthenticator" interfaces (#5905)
- `mongodbatlasreceiver` Add client wrapper for MongoDB Atlas support (#5386)
- `redisreceiver` Update Redis config options (#5861)
- `routingprocessor`: allow routing for all signals (#5869)
- `extension/observer/docker` add ListAndWatch to observer (#5851)

## v0.37.1

### 🧰 Bug fixes 🧰

- Fixes a problem with v0.37.0 which contained dependencies on v0.36.0 components. They should have been updated to v0.37.0.

## v0.37.0

### 🚀 New components 🚀

- [`journald` receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/journaldreceiver) to parse Journald events from systemd journal using the [opentelemetry-log-collection](https://github.com/open-telemetry/opentelemetry-log-collection) library

### 🛑 Breaking changes 🛑

- Remove squash on configtls.TLSClientSetting for splunkhecexporter (#5541)
- Remove squash on configtls.TLSClientSetting for elastic components (#5539)
- Remove squash on configtls.TLSClientSetting for observiqexporter (#5540)
- Remove squash on configtls.TLSClientSetting for AWS components (#5454)
- Move `k8sprocessor` to `k8sattributesprocessor`.
- Rename `k8s_tagger` configuration `k8sattributes`.
- filelog receiver: use empty value for `SeverityText` field instead of `"Undefined"` (#5423)
- Rename `configparser.ConfigMap` to `config.Map`
- Rename `pdata.AggregationTemporality*` to `pdata.MetricAggregationTemporality*`
- Remove deprecated `batchpertrace` package/module (#5380)

### 💡 Enhancements 💡

- `k8sattributes` processor: add container metadata enrichment (#5467, #5572)
- `resourcedetection` processor: Add an option to force using hostname instead of FQDN (#5064)
- `dockerstats` receiver: Move docker client into new shared `internal/docker` (#4702)
- `spanmetrics` processor:
  - Add exemplars to metrics (#5263)
  - Support resource attributes in metrics dimensions (#4624)
- `filter` processor:
  - Add log filtering by `regexp` type filters (#5237)
  - Add record level log filtering (#5418)
- `dynatrace` exporter: Handle non-gauge data types (#5056)
- `datadog` exporter:
  - Add support for exporting histograms as sketches (#5082)
  - Scrub sensitive information from errors (#5575)
  - Add option to send instrumentation library metadata tags with metrics (#5431)
- `podman` receiver: Add `api_version`, `ssh_key`, and `ssh_passphrase` config options (#5430)
- `signalfx` exporter:
  - Add `max_connections` config option (#5432)
  - Add dimension name to log when value > 256 chars (#5258)
  - Discourage setting of endpoint path (#4851)
- `kubeletstats` receiver: Convert to pdata instead of using OpenCensus (#5458)
- `tailsampling` processor: Add `invert_match` config option to `string_attribute` policy (#4393)
- `awsemf` exporter: Add a feature flag in UserAgent for AWS backend to monitor the adoptions (#5178)
- `splunkhec` exporter: Handle explicitly NaN and Inf values (#5581)
- `hostmetrics` receiver:
  - Collect more process states in processes scraper (#4856)
  - Add device label to paging scraper (#4854)
- `awskinesis` exporter: Extend to allow for dynamic export types (#5440)

### 🧰 Bug fixes 🧰

- `datadog` exporter:
  - Fix tags on summary and bucket metrics (#5416)
  - Fix cache key generation for cumulative metrics (#5417)
- `resourcedetection` processor: Fix failure to start collector if at least one detector returns an error (#5242)
- `prometheus` exporter: Do not record obsreport calls (#5438)
- `prometheus` receiver: Metric type fixes to match Prometheus functionality (#4865)
- `sentry` exporter: Fix sentry tracing (#4320)
- `statsd` receiver: Set quantiles for metrics (#5647)

## v0.36.0

### 🛑 Breaking changes 🛑

- `filter` processor: The configs for `logs` filter processor have been changed to be consistent with the `metrics` filter processor. (#4895)
- `splunk_hec` receiver:
  - `source_key`, `sourcetype_key`, `host_key` and `index_key` have now moved under `hec_metadata_to_otel_attrs` (#4726)
  - `path` field on splunkhecreceiver configuration is removed: We removed the `path` attribute as any request going to the Splunk HEC receiver port should be accepted, and added the `raw_path` field to explicitly map the path accepting raw HEC data. (#4951)
- feat(dynatrace): tags is deprecated in favor of default_dimensions (#5055)

### 💡 Enhancements 💡

- `filter` processor: Add ability to `include` logs based on resource attributes in addition to excluding logs based on resource attributes for strict matching. (#4895)
- `kubelet` API: Add ability to create an empty CertPool when the system run environment is windows
- `JMX` receiver: Allow JMX receiver logging level to be configured (#4898)
- `datadog` exporter: Export histograms as in OpenMetrics Datadog check (#5065)
- `dockerstats` receiver: Set Schema URL (#5239)
- Rename memorylimiter -> memorylimiterprocessor (#5262)
- `awskinesis` exporter: Refactor AWS kinesis exporter to be synchronous  (#5248)

## v0.35.0

### 🛑 Breaking changes 🛑

- Rename configparser.Parser to configparser.ConfigMap (#5070)
- Rename TelemetryCreateSettings -> TelemetrySettings (#5169)

### 💡 Enhancements 💡

- chore: update influxdb exporter and receiver (#5058)
- chore(dynatrace): use payload limit from api constants (#5077)
- Add documentation for filelog's new force_flush_period parameter (#5066)
- Reuse the gzip reader with a sync.Pool (#5145)
- Add a trace observer when splunkhecreceiver is used for logs (#5063)
- Remove usage of deprecated pdata.AttributeValueMapToMap (#5174)
- Podman Stats Receiver: Receiver and Metrics implementation (#4577)

### 🧰 Bug fixes 🧰

- Use staleness markers generated by prometheus, rather than making our own (#5062)
- `datadogexporter` exporter: skip NaN and infinite values (#5053)

## v0.34.0

### 🚀 New components 🚀

- [`cumulativetodelta` processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/cumulativetodeltaprocessor) to convert cumulative sum metrics to cumulative delta

- [`file` exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/fileexporter) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`jaeger` exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/jaegerexporter) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`kafka` exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/kafkaexporter) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`opencensus` exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/opencensusexporter) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`prometheus` exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/prometheusexporter) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`prometheusremotewrite` exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/prometheusremotewriteexporter) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`zipkin` exporter](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/zipkinexporter) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`attribute` processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/attributesprocessor) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`filter` processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/filterprocessor) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`probabilisticsampler` processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/probabilisticsamplerprocessor) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`resource` processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/resourceprocessor) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`span` processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/spanprocessor) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`hostmetrics` receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/hostmetricsreceiver) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`jaeger` receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/jaegerreceiver) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`kafka` receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/kafkareceiver) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`opencensus` receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/opencensusreceiver) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`prometheus` receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`zipkin` receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/zipkinreceiver) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`bearertokenauth` extension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/bearertokenauthextension) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`healthcheck` extension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/healthcheckextension) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`oidcauth` extension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/oidcauthextension) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`pprof` extension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/pprofextension) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))
- [`testbed`](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/testbed) from core repository ([#3474](https://github.com/open-telemetry/opentelemetry-collector/issues/3474))

### 💡 Enhancements 💡

- `tailsampling` processor: Add new policy `probabilistic` (#3876)

## v0.33.0

# 🎉 OpenTelemetry Collector Contrib v0.33.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.32.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- [`cumulativetodelta` processor](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/processor/cumulativetodeltaprocessor) to convert cumulative sum metrics to cumulative delta

### 💡 Enhancements 💡

- Collector contrib has now full support for metrics proto v0.9.0.

## v0.32.0

# 🎉 OpenTelemetry Collector Contrib v0.32.0 (Beta) 🎉

This release is marked as "bad" since the metrics pipelines will produce bad data.

- See https://github.com/open-telemetry/opentelemetry-collector/issues/3824

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.32.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🛑 Breaking changes 🛑

- `splunk_hec` receiver/exporter: `com.splunk.source` field is mapped to `source` field in Splunk instead of `service.name` (#4596)
- `redis` receiver: Move interval runner package to `internal/interval` (#4600)
- `datadog` exporter: Export summary count and sum as monotonic counts (#4605)

### 💡 Enhancements 💡

- `logzio` exporter:
  - New implementation of an in-memory queue to store traces, data compression with gzip, and queue configuration options (#4395)
  - Make `Hclog2ZapLogger` struct and methods private for public go api review (#4431)
- `newrelic` exporter (#4392):
  - Marked unsupported metric as permanent error
  - Force the interval to be valid even if 0
- `awsxray` exporter: Add PHP stacktrace parsing support (#4454)
- `file_storage` extension: Implementation of batch storage API (#4145)
- `datadog` exporter:
  - Skip sum metrics with no aggregation temporality (#4597)
  - Export delta sums as counts (#4609)
- `elasticsearch` exporter: Add dedot support (#4579)
- `signalfx` exporter: Add process metric to translation rules (#4598)
- `splunk_hec` exporter: Add profiling logs support (#4464)
- `awsemf` exporter: Replace logGroup and logStream pattern with metric labels (#4466)

### 🧰 Bug fixes 🧰

- `awsxray` exporter: Fix the origin on ECS/EKS/EB on EC2 cases (#4391)
- `splunk_hec` exporter: Prevent re-sending logs that were successfully sent (#4467)
- `signalfx` exporter: Prefix temporary metric translations (#4394)

## v0.31.0

# 🎉 OpenTelemetry Collector Contrib v0.31.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.31.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🛑 Breaking changes 🛑

- `influxdb` receiver: Removed `metrics_schema` config option (#4277)

### 💡 Enhancements 💡

- Update to OTLP 0.8.0:
  - Remove use of `IntHistogram` (#4276)
  - Update exporters/receivers for `NumberDataPoint`
- Remove use of deprecated `pdata` slice `Resize()` (#4203, #4208, #4209)
- `awsemf` exporter: Added the option to have a user who is sending metrics from EKS Fargate Container Insights to reformat them to look the same as insights from ECS so that they can be ingested by CloudWatch (#4130)
- `k8scluster` receiver: Support OpenShift cluster quota metrics (#4342)
- `newrelic` exporter (#4278):
  - Requests are now retry-able via configuration option (defaults to retries enabled). Permanent errors are not retried.
  - The exporter monitoring metrics now include an untagged summary metric for ease of use.
  - Improved error logging to include URLs that fail to post messages to New Relic.
- `datadog` exporter: Upscale trace stats when global sampling rate is set (#4213)

### 🧰 Bug fixes 🧰

- `statsd` receiver: Add option to set Counter to be monotonic (#4154)
- Fix `internal/stanza` severity mappings (#4315)
- `awsxray` exporter: Fix the wrong AWS env resource setting (#4384)
- `newrelic` exporter (#4278):
  - Configuration unmarshalling did not allow timeout value to be set to 0 in the endpoint specific section.
  - Request cancellation was not propagated via context into the http request.
  - The queued retry logger is set to a zap.Nop logger as intended.

## v0.30.0

# 🎉 OpenTelemetry Collector Contrib v0.30.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.30.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀
- `oauth2clientauth` extension: ported from core (#3848)
- `metrics-generation` processor: is now enabled and available (#4047)

### 🛑 Breaking changes 🛑

- Removed `jaegerthrifthttp` exporter (#4089)

### 💡 Enhancements 💡

- `tailsampling` processor:
  - Add new policy `status_code` (#3754)
  - Add new tail sampling processor policy: status_code (#3754)
- `awscontainerinsights` receiver:
  - Integrate components and fix bugs for EKS Container Insights (#3846)
  - Add Cgroup to collect ECS instance metrics for container insights receiver #3875
- `spanmetrics` processor: Support sub-millisecond latency buckets (#4091)
- `sentry` exporter: Add exception event capture in sentry (#3854)

## v0.29.0

# 🎉 OpenTelemetry Collector Contrib v0.29.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.29.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🛑 Breaking changes 🛑

- `redis` receiver (#3808)
  - removed configuration `service_name`. Use resource processor or `resource_attributes` setting if using `receivercreator`
  - removed `type` label and set instrumentation library name to `otelcol/redis` as other receivers do

### 💡 Enhancements 💡

- `tailsampling` processor:
  - Add new policy `latency` (#3750)
  - Add new policy `status_code` (#3754)
- `splunkhec` exporter: Include `trace_id` and `span_id` if set (#3850)
- `newrelic` exporter: Update instrumentation naming in accordance with otel spec (#3733)
- `sentry` exporter: Added support for insecure connection with Sentry (#3446)
- `k8s` processor:
  - Add namespace k8s tagger (#3384)
  - Add ignored pod names as config parameter (#3520)
- `awsemf` exporter: Add support for `TaskDefinitionFamily` placeholder on log stream name (#3755)
- `loki` exporter: Add resource attributes as Loki label (#3418)

### 🧰 Bug fixes 🧰

- `datadog` exporter:
  - Ensure top level spans are computed (#3786)
  - Update `env` clobbering behavior (#3851)
- `awsxray` exporter: Fixed filtered attribute translation (#3757)
- `splunkhec` exporter: Include trace and span id if set in log record (#3850)

## v0.28.0

# 🎉 OpenTelemetry Collector Contrib v0.28.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.28.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `humio` exporter to export data to Humio using JSON over the HTTP [Ingest API](https://docs.humio.com/reference/api/ingest/)
- `udplog` receiver to receives logs from udp using the [opentelemetry-log-collection](https://github.com/open-telemetry/opentelemetry-log-collection) library
- `tanzuobservability` exporter to send traces to [Tanzu Observability](https://tanzu.vmware.com/observability)

### 🛑 Breaking changes 🛑

- `f5cloud` exporter (#3509):
  - Renamed the config 'auth' field to 'f5cloud_auth'. This will prevent a config field name collision when [Support for Custom Exporter Authenticators as Extensions](https://github.com/open-telemetry/opentelemetry-collector/pull/3128) is ready to be integrated.

### 💡 Enhancements 💡

- Enabled Dependabot for Github Actions (#3543)
- Change obsreport helpers for receivers to use the new pattern created in Collector (#3439,#3443,#3449,#3504,#3521,#3548)
- `datadog` exporter:
  - Add logging for unknown or unsupported metric types (#3421)
  - Add collector version tag to internal health metrics (#3394)
  - Remove sublayer stats calc and mutex (#3531)
  - Deduplicate hosts for which we send running metrics (#3539)
  - Add support for summary datatype (#3660)
  - Add datadog span operation name remapping config option (#3444)
  - Update error formatting for error spans that are not exceptions (#3701)
- `nginx` receiver: Update the nginx metrics to more closely align with the conventions (#3420)
- `elasticsearch` exporter: Init JSON encoding support (#3101)
- `jmx` receiver:
  - Allow setting system properties (#3450)
  - Update tested JMX Metric Gatherer release (#3695)
- Refactor components for the Client Authentication Extensions (#3507)
- Remove redundant conversion calls (#3688)
- `storage` extension: Add a `Close` method to Client interface (#3506)
- `splunkhec` exporter: Add `metric_type` as key which maps to the type of the metric (#3696)
- `k8s` processor: Add semantic conventions to k8s-tagger for pod metadata (#3544)
- `kubeletstats` receiver: Refactor kubelet client to internal folder (#3698)
- `newrelic` exporter (#3690):
  - Updates the log level from error to debug when New Relic rate limiting occurs
  - Updates the sanitized api key that is reported via metrics
- `filestorage` extension: Add ability to specify name (#3703)
- `awsemf` exporter: Store the initial value for cumulative metrics (#3425)
- `awskinesis` exporter: Refactor to allow for extended types of encoding (#3655)
- `ecsobserver` extension:
  - Add task definition, ec2, and service fetcher (#3503)
  - Add exporter to convert task to target (#3333)

### 🧰 Bug fixes 🧰

- `awsemf` exporter: Remove delta adjustment from summaries by default (#3408)
- `alibabacloudlogservice` exporter: Sanitize labels for metrics (#3454)
- `statsd` receiver: Fix StatsD drop metrics tags when using summary as observer_type for timer/histogram (#3440)
- `awsxray` exporter: Restore setting of Throttle for HTTP throttle response (#3685)
- `awsxray` receiver: Fix quick start bug (#3653)
- `metricstransform` processor: Check all data points for matching metric label values (#3435)

## v0.27.0

# 🎉 OpenTelemetry Collector Contrib v0.27.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.27.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `tcplog` receiver to receive logs from tcp using the [opentelemetry-log-collection](https://github.com/open-telemetry/opentelemetry-log-collection) library
- `influxdb` receiver to accept metrics data as [InfluxDB Line Protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/)

### 💡 Enhancements 💡

- `splunkhec` exporter:
  - Include the response in returned 400 errors (#3338)
  - Map summary metrics to Splunk HEC metrics (#3344)
  - Add HEC telemetry (#3260)
- `newrelic` exporter: Include dropped attributes and events counts (#3187)
- `datadog` exporter:
  - Add Fargate task ARN to container tags (#3326)
  - Improve mappings for span kind dd span type (#3368)
- `signalfx` exporter: Add info log for host metadata properties update (#3343)
- `awsprometheusremotewrite` exporter: Add SDK and system information to User-Agent header (#3317)
- `metricstransform` processor: Add filtering capabilities matching metric label values for applying changes (#3201)
- `groupbytrace` processor: Added workers for queue processing (#2902)
- `resourcedetection` processor: Add docker detector (#2775)
- `tailsampling` processor: Support regex on span attribute filtering (#3335)

### 🧰 Bug fixes 🧰

- `datadog` exporter:
  - Update Datadog attributes to tags mapping (#3292)
  - Consistent `hostname` and default metrics behavior (#3286)
- `signalfx` exporter: Handle character limits on metric names and dimensions (#3328)
- `newrelic` exporter: Fix timestamp value for cumulative metrics (#3406)

## v0.26.0

# 🎉 OpenTelemetry Collector Contrib v0.26.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.26.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `influxdb` exporter to support sending tracing, metrics, and logging data to [InfluxDB](https://www.influxdata.com/products/)

### 🛑 Breaking changes 🛑

- `signalfx` exporter (#3207):
  - Additional metrics excluded by default by signalfx exporter
    - system.disk.io_time
    - system.disk.operation_time
    - system.disk.weighted_io_time
    - system.network.connections
    - system.processes.count
    - system.processes.created

### 💡 Enhancements 💡

- Add default config and systemd environment file support for DEB/RPM packages (#3123)
- Log errors on receiver start/stop failures (#3208)
- `newrelic` exporter: Update API key detection logic (#3212)
- `splunkhec` exporter:
  - Mark permanent errors to avoid futile retries (#3253)
  - Add TLS certs verification (#3204)
- `datadog` exporter:
  - Add env and tag name normalization to trace payloads (#3200)
  - add `ignore_resource`s configuration option (#3245)
- `jmx` receiver: Update for latest snapshot and header support (#3283)
- `awsxray` exporter: Added support for stack trace translation for .NET language (#3280)
- `statsd` receiver: Add timing/histogram for statsD receiver as OTLP summary (#3261)

### 🧰 Bug fixes 🧰

- `awsprometheusremotewrite` exporter:
  - Remove `sending_queue` (#3186)
  - Use the correct default for aws_auth.service (#3161)
  - Identify the Amazon Prometheus region from the endpoint (#3210)
  - Don't panic in case session can't be constructed (#3221)
- `datadog` exporter: Add max tag length (#3185)
- `sapm` exporter: Fix crash when passing the signalfx access token (#3294)
- `newrelic` exporter: Update error conditions (#3322)

## v0.25.0

# 🎉 OpenTelemetry Collector Contrib v0.25.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.25.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `kafkametricsreceiver` new receiver component for collecting metrics about a kafka cluster - primarily lag and offset. [configuration instructions](receiver/kafkametricsreceiver/README.md)
- `file_storage` extension to read and write data to the local file system (#3087)

### 🛑 Breaking changes 🛑

- `newrelic` exporter (#3091):
  - Removal of common attributes (use opentelemetry collector resource processor to add attributes)
  - Drop support for cumulative metrics being sent to New Relic via a collector

### 💡 Enhancements 💡

- Update `opentelemetry-log-collection` to v0.17.0 for log receivers (#3017)
- `datadog` exporter:
  - Add `peer.service` priority instead of `service.name` (#2817)
  - Improve support of semantic conventions for K8s, Azure and ECS (#2623)
- Improve and batch logs translation for stanza (#2892)
- `statsd` receiver: Add timing/histogram as OTLP gauge (#2973)
- `honeycomb` exporter: Add Retry and Queue settings (#2714)
- `resourcedetection` processor:
  - Add AKS resource detector (#3035)
  - Use conventions package constants for ECS detector (#3171)
- `sumologic` exporter: Add graphite format (#2695)
- Add trace attributes to the log entry for stanza (#3018)
- `splunk_hec` exporter: Send log record name as part of the HEC log event (#3119)
- `newrelic` exporter (#3091):
  - Add support for logs
  - Performance improvements
  - Optimizations to the New Relic payload to reduce payload size
  - Metrics generated for monitoring the exporter
  - Insert Key vs License keys are auto-detected in some cases
  - Collector version information is properly extracted via the application start info parameters

### 🧰 Bug fixes 🧰

- `splunk_hec` exporter: Fix sending log payload with missing the GZIP footer (#3032)
- `awsxray` exporter: Remove propagation of error on shutdown (#2999)
- `resourcedetection` processor:
  - Correctly report DRAGONFLYBSD value (#3100)
  - Fallback to `os.Hostname` when FQDN is not available (#3099)
- `httpforwarder` extension: Do not report ErrServerClosed when shutting down the service (#3173)
- `collectd` receiver: Do not report ErrServerClosed when shutting down the service (#3178)

## v0.24.0

# 🎉 OpenTelemetry Collector Contrib v0.24.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.24.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `fluentbit` extension and `fluentforward` receiver moved from opentelemetry-collector

### 💡 Enhancements 💡

- Check `NO_WINDOWS_SERVICE` environment variable to force interactive mode on Windows (#2819)
- `resourcedetection `processor:
  - Add task revision to ECS resource detector (#2814)
  - Add GKE detector (#2821)
  - Add Amazon EKS detector (#2820)
  - Add `VMScaleSetName` field to Azure detector (#2890)
- `awsemf` exporter:
  - Add `parse_json_encoded_attr_values` config option to decode json-encoded strings in attribute values (#2827)
  - Add `output_destination` config option to support AWS Lambda (#2720)
- `googlecloud` exporter: Handle `cloud.availability_zone` semantic convention (#2893)
- `newrelic` exporter: Add `instrumentation.provider` to default attributes (#2900)
- Set unprivileged user to container image (#2925)
- `splunkhec` exporter: Add `max_content_length_logs` config option to send log data in payloads less than max content length (#2524)
- `k8scluster` and `kubeletstats` receiver: Replace package constants in favor of constants from conventions in core (#2996)

### 🧰 Bug fixes 🧰

- `spanmetrics` processor:
  - Rename `calls` metric to `calls_total` and set `IsMonotonic` to true (#2837)
  - Validate duplicate dimensions at start (#2844)
- `awsemf` exporter: Calculate delta instead of rate for cumulative metrics (#2512)
- `signalfx` exporter:
  - Remove more unnecessary translation rules (#2889)
  - Implement summary type (#2998)
- `awsxray` exporter: Remove translation to HTTP status from OC status (#2978)
- `awsprometheusremotewrite` exporter: Close HTTP body after RoundTrip (#2955)
- `splunkhec` exporter: Add ResourceAttributes to Splunk Event (#2843)

## v0.23.0

# 🎉 OpenTelemetry Collector Contrib v0.23.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.23.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `groupbyattrs` processor to group the records by provided attributes
- `dotnetdiagnostics` receiver to read metrics from .NET processes

### 🛑 Breaking changes 🛑

- `stackdriver` exporter marked as deprecated and renamed to `googlecloud`
- Change the rule expression in receiver creator for matching endpoints types from `type.port`, `type.hostport` and `type.pod` to `type == "port"`, `type == "hostport"` and `type == "pod"` (#2661)

### 💡 Enhancements 💡

- `loadbalancing` exporter: Add support for logs (#2470)
- `sumologic` exporter: Add carbon formatter (#2562)
- `awsecscontainermetrics` receiver: Add new metric for stopped container (#2383)
- `awsemf` exporter:
  - Send EMF logs in batches (#2572)
  - Add prometheus type field for CloudWatch compatibility (#2689)
- `signalfx` exporter:
  - Add resource attributes to events (#2631)
  - Add translation rule to drop dimensions (#2660)
  - Remove temporary host translation workaround (#2652)
  - Remove unnecessary default translation rules (#2672)
  - Update `exclude_metrics` option so that the default exclude rules can be overridden by setting the option to `[]` (#2737)
- `awsprometheusremotewrite` exporter: Add support for given IAM roles (#2675)
- `statsd` receiver: Change to use OpenTelemetry type instead of OpenCensus type (#2733)
- `resourcedetection` processor: Add missing entries for `cloud.infrastructure_service` (#2777)

### 🧰 Bug fixes 🧰

- `dynatrace` exporter: Serialize each datapoint into separate line (#2618)
- `splunkhec` exporter: Retain all otel attributes (#2712)
- `newrelic` exporter: Fix default metric URL (#2739)
- `googlecloud` exporter: Add host.name label if hostname is present in node (#2711)

## v0.22.0

# 🎉 OpenTelemetry Collector Contrib v0.22.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.22.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `filelog` receiver to tail and parse logs from files using the [opentelemetry-log-collection](https://github.com/open-telemetry/opentelemetry-log-collection) library

### 💡 Enhancements 💡

- `dynatrace` exporter: Send metrics to Dynatrace in chunks of 1000 (#2468)
- `k8s` processor: Add ability to associate metadata tags using pod UID rather than just IP (#2199)
- `signalfx` exporter:
  - Add statusCode to logging field on dimension client (#2459)
  - Add translation rules for `cpu.utilization_per_core` (#2540)
  - Updates to metadata handling (#2531)
  - Calculate extra network I/O metrics (#2553)
  - Calculate extra disk I/O metrics (#2557)
- `statsd` receiver: Add metric type label and `enable_metric_type` option (#2466)
- `sumologic` exporter: Add support for carbon2 format (#2562)
- `resourcedetection` processor: Add Azure detector (#2372)
- `k8scluster` receiver: Use OTel conventions for metadata (#2530)
- `newrelic` exporter: Multi-tenant support for sending trace data and performance enhancements (#2481)
- `stackdriver` exporter: Enable `retry_on_failure` and `sending_queue` options (#2613)
- Use standard way to convert from time.Time to proto Timestamp (#2548)

### 🧰 Bug fixes 🧰

- `signalfx` exporter:
  - Fix calculation of `network.total` metric (#2551)
  - Correctly convert dimensions on metadata updates (#2552)
- `awsxray` exporter and receiver: Fix the type of content_length (#2539)
- `resourcedetection` processor: Use values in accordance to semantic conventions for AWS (#2556)
- `awsemf` exporter: Fix concurrency issue (#2571)

## v0.21.0

# 🎉 OpenTelemetry Collector Contrib v0.21.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.21.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `loki` exporter to export data via HTTP to Loki

### 🛑 Breaking changes 🛑

- `signalfx` exporter: Allow periods to be sent in dimension keys (#2456). Existing users who do not want to change this functionality can set `nonalphanumeric_dimension_chars` to `_-`

### 💡 Enhancements 💡

- `awsemf` exporter:
  - Support unit customization before sending logs to AWS CloudWatch (#2318)
  - Group exported metrics by labels (#2317)
- `datadog` exporter: Add basic span events support (#2338)
- `alibabacloudlogservice` exporter: Support new metrics interface (#2280)
- `sumologic` exporter:
  - Enable metrics pipeline (#2117)
  - Add support for all types of log body (#2380)
- `signalfx` exporter: Add `nonalphanumeric_dimension_chars` config option (#2442)

### 🧰 Bug fixes 🧰

- `resourcedetection` processor: Fix resource attribute environment variable (#2378)
- `k8scluster` receiver: Fix nil pointer bug (#2450)

## v0.20.0

# 🎉 OpenTelemetry Collector Contrib v0.20.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.20.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `spanmetrics` processor to aggregate Request, Error and Duration (R.E.D) metrics from span data
- `awsxray` receiver to accept spans in the X-Ray Segment format
- `groupbyattrs` processor to group the records by provided attributes

### 🛑 Breaking changes 🛑

- Rename `kinesis` exporter to `awskinesis` (#2234)
- `signalfx` exporter: Remove `send_compatible_metrics` option, use `translation_rules` instead (#2267)
- `datadog` exporter: Remove default prefix from user metrics (#2308)

### 💡 Enhancements 💡

- `signalfx` exporter: Add k8s metrics to default excludes (#2167)
- `stackdriver` exporter: Reduce QPS (#2191)
- `datadog` exporter:
  - Translate otel exceptions to DataDog errors (#2195)
  - Use resource attributes for metadata and generated metrics (#2023)
- `sapm` exporter: Enable queuing by default (#1224)
- `dynatrace` exporter: Allow underscores anywhere in metric or dimension names (#2219)
- `awsecscontainermetrics` receiver: Handle stopped container's metadata (#2229)
- `awsemf` exporter: Enhance metrics batching in AWS EMF logs (#2271)
- `f5cloud` exporter: Add User-Agent header with version to requests (#2292)

### 🧰 Bug fixes 🧰

- `signalfx` exporter: Reinstate network/filesystem translation rules (#2171)

## v0.19.0

# 🎉 OpenTelemetry Collector Contrib v0.19.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.19.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `f5cloud` exporter to export metric, trace, and log data to F5 Cloud
- `jmx` receiver to report metrics from a target MBean server in conjunction with the [JMX Metric Gatherer](https://github.com/open-telemetry/opentelemetry-java-contrib/blob/v1.0.0-alpha/contrib/jmx-metrics/README.md)

### 🛑 Breaking changes 🛑

- `signalfx` exporter: The `exclude_metrics` option now takes slice of metric filters instead of just metric names (slice of strings) (#1951)

### 💡 Enhancements 💡

- `datadog` exporter: Sanitize datadog service names (#1982)
- `awsecscontainermetrics` receiver: Add more metadata (#2011)
- `azuremonitor` exporter: Favor RPC over HTTP spans (#2006)
- `awsemf` exporter: Always use float64 as calculated rate (#2019)
- `splunkhec` receiver: Make the HEC receiver path configurable, and use `/*` by default (#2137)
- `signalfx` exporter:
  - Drop non-default metrics and add `include_metrics` option to override (#2145, #2146, #2162)
  - Rename `system.network.dropped_packets` metric to `system.network.dropped` (#2160)
  - Do not filter cloud attributes from dimensions (#2020)
- `redis` receiver: Migrate to pdata metrics #1889

### 🧰 Bug fixes 🧰

- `datadog` exporter: Ensure that version tag is added to trace stats (#2010)
- `loadbalancing` exporter: Rolling update of collector can stop the periodical check of DNS updates (#1798)
- `awsecscontainermetrics` receiver: Change the type of `exit_code` from string to int and deal with the situation when there is no data (#2147)
- `groupbytrace` processor: Make onTraceReleased asynchronous to fix processor overload (#1808)
- Handle cases where the time field of Splunk HEC events is encoded as a String (#2159)

## v0.18.0

# 🎉 OpenTelemetry Collector Contrib v0.18.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.18.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `sumologic` exporter to send logs and metrics data to Sumo Logic
- `dynatrace` exporter to send metrics to Dynatrace

### 💡 Enhancements 💡

- `datadog` exporter:
  - Add resource attributes to tags conversion feature (#1782)
  - Add Kubernetes conventions for hostnames (#1919)
  - Add container tags to datadog export for container infra metrics in service view (#1895)
  - Update resource naming and span naming (#1861)
  - Add environment variables support for config options (#1897)
- `awsxray` exporter: Add parsing of JavaScript stack traces (#1888)
- `elastic` exporter: Translate exception span events (#1858)
- `signalfx` exporter: Add translation rules to aggregate per core CPU metrics in default translations (#1841)
- `resourcedetection` processor: Gather tags associated with the EC2 instance and add them as resource attributes (#1899)
- `simpleprometheus` receiver: Add support for passing params to the prometheus scrape config (#1949)
- `azuremonitor` exporter: Implement Span status code specification changes - gRPC (#1960)
- `metricstransform` processor: Add grouping option ($1887)
- `alibabacloudlogservice` exporter: Use producer to send data to improve performance (#1981)

### 🧰 Bug fixes 🧰

- `datadog` exporter: Handle monotonic metrics client-side (#1805)
- `awsxray` exporter: Log error when translating span (#1809)

## v0.17.0

# 🎉 OpenTelemetry Collector Contrib v0.17.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.17.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 💡 Enhancements 💡

- `awsemf` exporter: Add collector version to EMF exporter user agent (#1778)
- `signalfx` exporter: Add configuration for trace correlation (#1795)
- `statsd` receiver: Add support for metric aggregation (#1670)
- `datadog` exporter: Improve logging of hostname detection (#1796)

### 🧰 Bug fixes 🧰

- `resourcedetection` processor: Fix ecs detector to not use the default golang logger (#1745)
- `signalfx` receiver: Return 200 when receiver succeed (#1785)
- `datadog` exporter: Use a singleton for sublayer calculation (#1759)
- `awsxray` and `awsemf` exporters: Change the User-Agent content order (#1791)

## v0.16.0

# 🎉 OpenTelemetry Collector Contrib v0.16.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.16.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🛑 Breaking changes 🛑

- `honeycomb` exporter: Update to use internal data format (#1689)

### 💡 Enhancements 💡

- `newrelic` exporter: Add support for span events (#1643)
- `awsemf` exporter:
  - Add placeholder support in `log_group_name` and `log_stream_name` config (#1623, #1661)
  - Add label matching filtering rule (#1619)
- `resourcedetection` processor: Add new resource detector for AWS Elastic Beanstalk environments (#1585)
- `loadbalancing` exporter:
  - Add sort of endpoints in static resolver (#1692)
  - Allow specifying port when using DNS resolver (#1650)
- Add `batchperresourceattr` helper library that splits an incoming data based on an attribute in the resource (#1694)
- `alibabacloudlogservice` exporter:
  - Add logs exporter (#1609)
  - Change trace type from opencensus to opentelemetry (#1713)
- `datadog` exporter:
  - Improve trace exporter performance (#1706, #1707)
  - Add option to only send metadata (#1723)
- `awsxray` exporter:
  - Add parsing of Python stack traces (#1676)
  - Add collector version to user agent (#1730)

### 🧰 Bug fixes 🧰

- `loadbalancing` exporter:
  - Fix retry queue for exporters (#1687)
  - Fix `periodicallyResolve` for DNS resolver checks (#1678)
- `datadog` exporter: Fix status code handling (#1691)
- `awsxray` exporter:
  - Fix empty traces in X-Ray console (#1709)
  - Stricter requirements for adding http request url (#1729)
  - Fix status code handling for errors/faults (#1740)
- `signalfx` exporter:
  - Split incoming data requests by access token before enqueuing (#1727)
  - Disable retry on 400 and 401, retry with backoff on 429 and 503 (#1672)
- `awsecscontainermetrics` receiver: Improve error handling to fix seg fault (#1738)

## v0.15.0

# 🎉 OpenTelemetry Collector Contrib v0.15.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.15.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `zookeeper` receiver: Collects metrics from a Zookeeper instance using the `mntr` command
- `loadbalacing` exporter: Consistently exports spans belonging to the same trace to the same backend
- `windowsperfcounters` receiver: Captures the configured system, application, or custom performance counter data from the Windows registry using the PDH interface
- `awsprometheusremotewrite` exporter:  Sends metrics data in Prometheus TimeSeries format to a Prometheus Remote Write Backend and signs each outgoing HTTP request following the AWS Signature Version 4 signing process

### 💡 Enhancements 💡

- `awsemf` exporter:
  - Add `metric_declarations` config option for metric filtering and dimensions (#1503)
  - Add SummaryDataType and remove Min/Max from Histogram (#1584)
- `signalfxcorrelation` exporter: Add ability to translate host dimension (#1561)
- `newrelic` exporter: Use pdata instead of the OpenCensus for traces (#1587)
- `metricstransform` processor:
  - Add `combine` action for matched metrics (#1506)
  - Add `submatch_case` config option to specify case of matched label values (#1640)
- `awsecscontainermetrics` receiver: Extract cluster name from ARN (#1626)
- `elastic` exporter: Improve handling of span status if the status code is unset (#1591)

### 🧰 Bug fixes 🧰

- `awsemf` exporter: Add check for unhandled metric data types (#1493)
- `groupbytrace` processor: Make buffered channel to avoid goroutines leak (#1505)
- `stackdriver` exporter: Set `options.UserAgent` so that the OpenCensus exporter does not override the UA ($1620)

## v0.14.0

# 🎉 OpenTelemetry Collector Contrib v0.14.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.14.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `datadog` exporter to send metric and trace data to Datadog (#1352)
- `tailsampling` processor moved from core to contrib (#1383)

### 🛑 Breaking changes 🛑

- `jmxmetricsextension` migrated to `jmxreceiver` (#1182, #1357)
- Move signalfx correlation code out of `sapm` to `signalfxcorrelation` exporter (#1376)
- Move Splunk specific utils outside of common (#1306)
- `stackdriver` exporter:
    - Config options `metric_prefix` & `skip_create_metric_descriptor` are now nested under `metric`, see [README](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/v0.14.0/exporter/stackdriverexporter/README.md).
    - Trace status codes no longer reflect gRPC codes as per spec changes: open-telemetry/opentelemetry-specification#1067
- `datadog` exporter: Remove option to change the namespace prefix (#1483)

### 💡 Enhancements 💡

- `splunkhec` receiver: Add ability to ingest metrics (#1276)
- `signalfx` receiver: Improve pipeline error handling (#1329)
- `datadog` exporter:
  - Improve hostname resolution (#1285)
  - Add flushing/export of traces and trace-related statistics (#1266)
  - Enable traces on Windows (#1340)
  - Send otel.exporter running metric (#1354)
  - Add tag normalization util method (#1373)
  - Send host metadata (#1351)
  - Support resource conventions for hostnames (#1434)
  - Add version tag extract (#1449)
- Add `batchpertrace` library to split the incoming batch into several batches, one per trace (#1257)
- `statsd` receiver:
  - Add timer support (#1335)
  - Add sample rate support for counter, transfer gauge to double and transfer counter to int only (#1361)
- `awsemf` exporter: Restructure metric translator logic (#1353)
- `resourcedetection` processor:
  - Add EC2 hostname attribute (#1324)
  - Add ECS Resource detector (#1360)
- `sapm` exporter: Add queue settings (#1390)
- `metrictransform` processor: Add metric filter option (#1447)
- `awsxray` exporter: Improve ECS attribute and origin translation (#1428)
- `resourcedetection` processor: Initial system detector (#1405)

### 🧰 Bug fixes 🧰

- Remove duplicate definition of cloud providers with core conventions (#1288)
- `kubeletstats` receiver: Handle nil references from the kubelet API (#1326)
- `awsxray` receiver:
  - Add kind type to root span to fix the empty parentID problem (#1338)
  - Fix the race condition issue (#1490)
- `awsxray` exporter:
  - Setting the tlsconfig InsecureSkipVerify using NoVerifySSL (#1350)
  - Drop invalid xray trace id (#1366)
- `elastic` exporter: Ensure span name is limited (#1371)
- `splunkhec` exporter: Don't send 'zero' timestamps to Splunk HEC (#1157)
- `stackdriver` exporter: Skip processing empty metrics slice (#1494)

## v0.13.0

# 🎉 OpenTelemetry Collector Contrib v0.13.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.13.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 💡 Enhancements 💡

- `sapm` exporter:
  - Enable queuing by default (#1224)
  - Add SignalFx APM correlation (#1205)
  - Make span source attribute and destination dimension names configurable (#1286)
- `signalfx` exporter:
  - Pass context to the http client requests (#1225)
  - Update `disk.summary_utilization` translation rule to accommodate new labels (#1258)
- `newrelic` exporter: Add `span.kind` attribute (#1263)
- `datadog` exporter:
  - Add Datadog trace translation helpers (#1208)
  - Add API key validation (#1216)
- `splunkhec` receiver: Add the ability to ingest logs (#1268)
- `awscontainermetrics` receiver: Report `CpuUtilized` metric in percentage (#1283)
- `awsemf` exporter: Only calculate metric rate for cumulative counter and avoid SingleDimensionRollup for metrics with only one dimension (#1280)

### 🧰 Bug fixes 🧰

- Make `signalfx` exporter a metadata exporter (#1252)
- `awsecscontainermetrics` receiver: Check for empty network rate stats and set zero (#1260)
- `awsemf` exporter: Remove InstrumentationLibrary dimension in CloudWatch EMF Logs if it is undefined (#1256)
- `awsxray` receiver: Fix trace/span id transfer (#1264)
- `datadog` exporter: Remove trace support for Windows for now (#1274)
- `sapm` exporter: Correlation enabled check inversed (#1278)

## v0.12.0

# 🎉 OpenTelemetry Collector Contrib v0.12.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.12.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- `awsemf` exporter to support exporting metrics to AWS CloudWatch (#498, #1169)
- `http_forwarder` extension that forwards HTTP requests to a specified target (#979, #1014, #1150)
- `datadog` exporter that sends metric and trace data to Datadog (#1142, #1178, #1181, #1212)
- `awsecscontainermetrics` receiver to collect metrics from Amazon ECS Task Metadata Endpoint (#1089, #1148, #1160)

### 💡 Enhancements 💡

- `signalfx` exporter:
  - Add host metadata synchronization (#1039, #1118)
  - Add `copy_dimensions` translator option (#1126)
  - Update `k8s_cluster` metric translations (#1121)
  - Add option to exclude metrics (#1156)
  - Add `avg` aggregation method (#1151)
  - Fallback to host if cloud resource id not found (#1170)
  - Add backwards compatible translation rules for the `dockerstatsreceiver` (#1201)
  - Enable queuing and retries (#1223)
- `splunkhec` exporter:
  - Add log support (#875)
  - Enable queuing and retries (#1222)
- `k8scluster` receiver: Standardize metric names (#1119)
- `awsxray` exporter:
  - Support AWS EKS attributes (#1090)
  - Store resource attributes in X-Ray segments (#1174)
- `honeycomb` exporter:
  - Add span kind to the event sent to Honeycomb (#474)
  - Add option to adjust the sample rate using an attribute on the span (#1162)
- `jmxmetrics` extension: Add subprocess manager to manage child java processes (#1028)
- `elastic` exporter: Initial metrics support (#1173)
- `k8s` processor: Rename default attr names for label/annotation extraction (#1214)
- Add common SignalFx host id extraction (#1100)
- Allow MSI upgrades (#1165)

### 🧰 Bug fixes 🧰

- `awsxray` exporter: Don't set origin to EC2 when not on AWS (#1115)

## v0.11.0

# 🎉 OpenTelemetry Collector Contrib v0.11.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.11.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀
- add `dockerstats` receiver as top level component (#1081)
- add `tracegen` utility (#956)

### 💡 Enhancements 💡
- `stackdriver` exporter: Allow overriding client options via config (#1010)
- `k8scluster` receiver: Ensure informer caches are synced before initial data sync (#842)
- `elastic` exporter: Translate `deployment.environment` resource attribute to Elastic APM's semantically equivalent `service.environment` (#1022)
- `k8s` processor: Add logs support (#1051)
- `awsxray` exporter: Log response error with zap (#1050)
- `signalfx` exporter
  - Add dimensions to renamed metrics (#1041)
  - Add translation rules for `disk_ops.total` and `disk_ops.pending` metrics (#1082)
  - Add event support (#1036)
- `kubeletstats` receiver: Cache detailed PVC labels to reduce API calls (#1052)
- `signalfx` receiver: Add event support (#1035)

## v0.10.0

# 🎉 OpenTelemetry Collector Contrib v0.10.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.10.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀
- add initial docker stats receiver, without sourcing in top level components (#495)
- add initial jmx metrics extension structure, without sourcing in top level components (#740)
- `routing` processor for routing spans based on HTTP headers (#907)
- `splunkhec` receiver to receive Splunk HEC metrics, traces and logs (#840)
- Add skeleton for `http_forwarder` extension that forwards HTTP requests to a specified target (#979)

### 💡 Enhancements 💡
- `stackdriver` exporter
  - Add timeout parameter (#835)
  - Add option to configurably set UserAgent string (#758)
- `signalfx` exporter
  - Reduce memory allocations for big batches processing (#871)
  - Add AWSUniqueId and gcp_id generation (#829)
  - Calculate cpu.utilization compatibility metric (#839, #974, #954)
- `metricstransform` processor: Replace `{{version}}` in label values (#876)
- `resourcedetection` processor: Logs Support (#970)
- `statsd` receiver: Add parsing for labels and gauges (#903)

### 🧰 Bug fixes 🧰
- `k8s` processor
  - Wrap metrics before sending further down the pipeline (#837)
  - Fix setting attributes on metrics passed from agent (#836)
- `awsxray` exporter: Fix "pointer to empty string" is not omitted bug (#830)
- `azuremonitor` exporter: Treat UNSPECIFIED span kind as INTERNAL (#844)
- `signalfx` exporter: Remove misleading warnings (#869)
- `newrelic` exporter: Fix panic if service name is empty (#969)
- `honeycomb` exporter: Don't emit default proc id + starttime (#972)

## v0.9.0

# 🎉 OpenTelemetry Collector Contrib v0.9.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.9.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🛑 Breaking changes 🛑
- Remove deprecated `lightstep` exporter (#828)

### 🚀 New components 🚀
- `statsd` receiver for ingesting StatsD messages (#566)

### 💡 Enhancements 💡
- `signalfx` exporter
   - Add disk usage translations (#760)
   - Add disk utilization translations (#782)
   - Add translation rule to drop redundant metrics (#809)
- `kubeletstats` receiver
  - Sync available volume metadata from /pods endpoint (#690)
  - Add ability to collect detailed data from PVC (#743)
- `awsxray` exporter: Translate SDK name/version into xray model (#755)
- `elastic` exporter: Translate semantic conventions to Elastic destination fields (#671)
- `stackdriver` exporter: Add point count metric (#757)
- `awsxray` receiver
  - Ported the TCP proxy from the X-Ray daemon (#774)
  - Convert to OTEL trace format (#691)

### 🧰 Bug fixes 🧰
- `kubeletstats` receiver: Do not break down metrics batch (#754)
- `host` observer: Fix issue on darwin where ports listening on all interfaces are not correctly accounted for (#582)
- `newrelic` exporter: Fix panic on missing span status (#775)

## v0.8.0

# 🎉 OpenTelemetry Collector Contrib v0.8.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.8.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- Receivers
  - `prometheusexec` subprocess manager (##499)

### 💡 Enhancements 💡

- `signalfx` exporter
  - Add/Update metric translations (#579, #584, #639, #640, #652, #662)
  - Add support for calculate new metric translator (#644)
  - Add renaming rules for load metrics (#664)
  - Update `container.name` to `k8s.container.name` in default translation rule (#683)
  - Rename working-set and page-fault metrics (#679)
- `awsxray` exporter
  - Translate exception event into xray exception (#577)
  - Add ingestion of X-Ray segments via UDP (#502)
  - Parse Java stacktrace and populate in xray cause (#687)
- `kubeletstats` receiver
  - Add metric_groups option (#648)
  - Set datapoint timestamp in receiver (#661)
  - Change `container.name` label to `k8s.container.name` (#680)
  - Add working-set and page-fault metrics (#666)
  - Add basic support for volume metrics (#667)
- `stackdriver` trace exporter: Move to new interface and pdata (#486)
- `metricstranform` processor: Keep timeseries and points in order after aggregation (#663)
- `k8scluster` receiver: Change `container.spec.name` label to `k8s.container.name` (#681)
- Migrate receiver creator to internal data model (#701)
- Add ec2 support to `resourcedetection` processor (#587)
- Enable timeout, sending queue and retry for SAPM exporter (#707)

### 🧰 Bug fixes 🧰

- `azuremonitor` exporter: Correct HTTP status code success mapping (#588)
- `k8scluster` receiver: Fix owner reference in metadata updates (#649)
- `awsxray` exporter: Fix handling of db system (#697)

### 🚀 New components 🚀

- Skeleton for AWS ECS container metrics receiver (#463)
- `prometheus_exec` receiver (#655)

## v0.7.0

# 🎉 OpenTelemetry Collector Contrib v0.7.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.7.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🛑 Breaking changes 🛑

- `awsxray` receiver updated to support udp: `tcp_endpoint` config option renamed to `endpoint` (#497)
- TLS config changed for `sapmreceiver` (#488) and `signalfxreceiver` receivers (#488)

### 🚀 New components 🚀

- Exporters
  - `sentry` adds tracing exporter for [Sentry](https://sentry.io/) (#565)
- Extensions
  - `endpoints` observer: adds generic endpoint watcher (#427)
  - `host` observer: looks for listening network endpoints on host (#432)

### 💡 Enhancements 💡

- Update `honeycomb` exporter for v0.8.0 compatibility
- Extend `metricstransform` processor to be able to add a label to an existing metric (#441)
- Update `kubeletstats` metrics according to semantic conventions (#475)
- Updated `awsxray` receiver config to use udp (#497)
- Add `/pods` endpoint support in `kubeletstats` receiver to add extra labels (#569)
- Add metric translation options to `signalfx` exporter (#477, #501, #571, #573)

### 🧰 Bug fixes 🧰

- `azuremonitor` exporter: Mark spanToEnvelope errors as permanent (#500)

## v0.6.0

# 🎉 OpenTelemetry Collector Contrib v0.6.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.6.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🛑 Breaking changes 🛑

- Removed `jaegarlegacy` (#397) and `zipkinscribe` receivers (#410)
- `kubeletstats` receiver: Renamed `k8s.pod.namespace` pod label to `k8s.namespace.name` and `k8s.container.name` container label to `container.name`

### 🚀 New components 🚀

- Processors
  - `metricstransform` renames/aggregates within individual metrics (#376) and allow changing the data type between int and float (#402)

### 💡 Enhancements 💡

- `awsxray` exporter: Use `peer.service` as segment name when set. (#385)
- `splunk` exporter: Add trace exports support (#359, #399)
- Build and publish Windows MSI (#408) and DEB/RPM Linux packages (#405)

### 🧰 Bug fixes 🧰

- `kubeletstats` receiver:
  - Fixed NPE for newly created pods (#404)
  - Updated to latest change in the ReceiverFactoryOld interface (#401)
  - Fixed logging and self reported metrics (#357)
- `awsxray` exporter: Only convert SQL information for SQL databases. (#379)
- `resourcedetection` processor: Correctly obtain machine-type info from gce metadata (#395)
- `k8scluster` receiver: Fix container resource metrics (#416)

## v0.5.0

Released 01-07-2020

# 🎉 OpenTelemetry Collector Contrib v0.5.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.5.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🚀 New components 🚀

- Processors
  - `resourcedetection` to automatically detect the resource based on the configured set of detectors (#309)

### 💡 Enhancements 💡

- `kubeletstats` receiver: Support for ServiceAccount authentication (#324)
- `signalfx` exporter and receiver
  - Add SignalFx metric token passthrough and config option (#325)
  - Set default endpoint of `signalfx` receiver to `:9943` (#351)
- `awsxray` exporter: Support aws plugins EC2/ECS/Beanstalk (#343)
- `sapm` exporter and receiver: Add SAPM access token passthrough and config option (#349)
- `k8s` processor: Add metrics support (#358)
- `k8s` observer: Separate annotations from labels in discovered pods (#363)

### 🧰 Bug fixes 🧰

- `honeycomb` exporter: Remove shared use of libhoney from goroutines (#305)

## v0.4.0

Released 17-06-2020

# 🎉 OpenTelemetry Collector Contrib v0.4.0 (Beta) 🎉

The OpenTelemetry Collector Contrib contains everything in the [opentelemetry-collector release](https://github.com/open-telemetry/opentelemetry-collector/releases/tag/v0.4.0) (be sure to check the release notes here as well!). Check out the [Getting Started Guide](https://opentelemetry.io/docs/collector/getting-started/) for deployment and configuration information.

### 🛑 Breaking changes 🛑

  - `signalfx` exporter `url` parameter changed to `ingest_url` (no impact if only using `realm` setting)

### 🚀 New components 🚀

- Receivers
  - `receiver_creator` to create receivers at runtime (#145), add observer support to receiver_creator (#173), add rules support (#207), add dynamic configuration values (#235)
  - `kubeletstats` receiver (#237)
  - `prometheus_simple` receiver (#184)
  - `kubernetes-cluster` receiver (#175)
  - `redis` receiver (#138)
- Exporters
  - `alibabacloudlogservice` exporter (#259)
  - `SplunkHEC` metrics exporter (#246)
  - `elastic` APM exporter (#240)
  - `newrelic` exporter (#229)
- Extensions
  - `k8s` observer (#185)

### 💡 Enhancements 💡

- `awsxray` exporter
  - Use X-Ray convention of segment name == service name (#282)
  - Tweak xray export to improve rendering of traces and improve parity (#241)
  - Add handling for spans received with nil attributes (#212)
- `honeycomb` exporter
  - Use SendPresampled (#291)
  - Add span attributes as honeycomb event fields (#271)
  - Support resource labels in Honeycomb exporter (#20)
- `k8s` processor
  - Add support of Pod UID extraction to k8sprocessor (#219)
  - Use `k8s.pod.ip` to record resource IP instead of just `ip` (#183)
  - Support same authentication mechanism as other kubernetes components do (#307)
- `sapm` exporter: Add TLS for SAPM and SignalFx receiver (#215)
- `signalfx` exporter
  - Add metric metadata syncer to SignalFx exporter (#231)
  - Add TLS for SAPM and SignalFx receiver (#215)
- `stackdriver` exporter: Add support for resource mapping in config (#163)

### 🧰 Bug fixes 🧰

- `awsxray` exporter: Wrap bad request errors for proper handling by retry queue (#205)
- `lightstep` exporter: Ensure Lightstep exporter doesnt crash on nil node (#250)
- `sapm` exporter: Do not break Jaeger traces before sending downstream (#193)
- `k8s` processor: Ensure Jaeger spans work in passthrough mode (262)

## 🧩 Components 🧩

### Receivers

|       Traces        |      Metrics      |
| :-----------------: | :---------------: |
|    Jaeger Legacy    |      Carbon       |
| SAPM (SignalFx APM) |     Collectd      |
|    Zipkin Scribe    |    K8s Cluster    |
|                     |       Redis       |
|                     |     SignalFx      |
|                     | Simple Prometheus |
|                     |     Wavefront     |

### Processors

- K8s

### Exporters

|        Commercial         |   Community   |
| :-----------------------: | :-----------: |
| Alibaba Cloud Log Service |    Carbon     |
|         AWS X-ray         |    Elastic    |
|       Azure Monitor       | Jaeger Thrift |
|         Honeycomb         |    Kinesis    |
|         Lightstep         |
|         New Relic         |
|    SAPM (SignalFx APM)    |
|    SignalFx (Metrics)     |
|        Splunk HEC         |
|   Stackdriver (Google)    |

### Extensions

- Observer
  - K8s

## v0.3.0 Beta

Released 2020-03-30

### Breaking changes

-  Make prometheus receiver config loading strict. #697
Prometheus receiver will now fail fast if the config contains unused keys in it.

### Changes and fixes

- Enable best effort serve by default of Prometheus Exporter (https://github.com/orijtech/prometheus-go-metrics-exporter/pull/6)
- Fix null pointer exception in the logging exporter #743
- Remove unnecessary condition to have at least one processor #744
- Updated Honeycomb exported to `honeycombio/opentelemetry-exporter-go v0.3.1`

### Features

Receivers / Exporters:

* AWS X-Ray
* Carbon
* CollectD
* Honeycomb
* Jaeger
* Kinesis
* LightStep
* OpenCensus
* OpenTelemetry
* SAPM
* SignalFx
* Stackdriver
* Wavefront
* Zipkin
* Zipkin Scribe


Processors:

* Attributes
* Batch
* Memory Limiter
* Queued Retry
* Resource
* Sampling
* Span
* Kubernetes

Extensions:

* Health Check
* Performance Profiler
* zPages


## v0.2.8

Released 2020-03-25

Alpha v0.2.8 of OpenTelemetry Collector Contrib.

- Implemented OTLP receiver and exporter.
- Added ability to pass config to the service programmatically (useful for custom builds).
- Improved own metrics / observability.


## v0.2.7

Released 2020-03-17

### Self-Observability
- New command-line switch to control legacy and new metrics. Users are encouraged
to experiment and migrate to the new metrics.
- Improved error handling on shutdown.


### Processors
- Fixed passthrough mode k8sprocessor.
- Added `HASH` action to attribute processor.

### Receivers and Exporters
- Added Honeycomb exporter.
- Added LightStep exporter.
- Added regular expression for Carbon receiver, allowing the metric name to be broken into proper label keys and values.
- Updated Stackdriver exporter to use a new batch API.


## v0.2.6 Alpha

Released 2020-02-18

### Self-Observability
- Updated metrics prefix to `otelcol` and expose command line argument to modify the prefix value.
- Batch dropped span now emits zero when no spans are dropped.

### Processors
- Extended Span processor to have include/exclude span logic.
- Ability to choose strict or regexp matching for include/exclude filters.

### Receivers and Exporters
- Added Carbon receiver and exporter.
- Added Wavefront receiver.


## v0.0.5 Alpha

Released 2020-01-30

- Regexp-based filtering of span names.
- Ability to extract attributes from span names and rename span.
- File exporter for debugging.
- Span processor is now enabled by default.

## v0.0.1 Alpha

Released 2020-01-11

First release of OpenTelemetry Collector Contrib.


[v0.3.0]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.2.8...v0.3.0
[v0.2.8]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.2.7...v0.2.8
[v0.2.7]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.2.6...v0.2.7
[v0.2.6]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.0.5...v0.2.6
[v0.0.5]: https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.0.1...v0.0.5
[v0.0.1]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.0.1
