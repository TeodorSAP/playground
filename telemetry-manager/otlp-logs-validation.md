# OTel LogPipeline set-up validation

## Set-up configuration steps

```
k create ns otel-logging

k apply -f /Users/I565663/Projects/telemetry-manager/config/samples/operator_v1alpha1_telemetry.yaml

kh-cls-log // Execute knowledge-hub/scripts/create_cls_log_pipeline.sh with the corresponding environment variables 

helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts

tm && helm install -n otel-logging logging open-telemetry/opentelemetry-collector -f docs/contributor/pocs/assets/otel-logs-values.yaml

tm && helm install -n otel-logging logging open-telemetry/opentelemetry-collector -f ./docs/contributor/pocs/assets/otel-log-agent-values.yaml
```

## Resulting Resources

### LogPipeline ConfigMap

```
exporters:
  otlp:
    endpoint: telemetry-otlp-logs.kyma-system:4317
    tls:
      insecure: true
    retry_on_failure:
        enabled: true
        initial_interval: 5s
        max_interval: 30s
        max_elapsed_time: 300s

extensions:
  file_storage:
    directory: /var/lib/otelcol
  health_check:
    endpoint: ${env:MY_POD_IP}:13133
  pprof:
    endpoint: 127.0.0.1:1777

processors:
  memory_limiter:
    check_interval: 5s
    limit_percentage: 80
    spike_limit_percentage: 25
  transform/set-instrumentation-scope-runtime:
    error_mode: ignore
    metric_statements:
        - context: scope
          statements:
            - set(version, "main")
            - set(name, "io.kyma-project.telemetry/runtime")
            - set(version, "main")
            - set(name, "io.kyma-project.telemetry/runtime")

receivers:
  filelog:
    exclude:
    - /var/log/pods/kyma-system_telemetry-log-agent_*/*/*.log # exclude self
    - /var/log/pods/kyma-system_telemetry-fluent-bit_*/*/*.log # exclude FluentBit
    include:
    - /var/log/pods/*/*/*.log
    include_file_name: false
    include_file_path: true
    operators:
    - id: parser-crio
      regex: ^(?P<time>[^ Z]+) (?P<stream>stdout|stderr) (?P<logtag>[^ ]*) ?(?P<log>.*)$
      timestamp:
        layout: 2006-01-02T15:04:05.999999999Z07:00
        layout_type: gotime
        parse_from: attributes.time
      type: regex_parser
    - combine_field: attributes.log
      combine_with: ""
      id: crio-recombine
      is_last_entry: attributes.logtag == 'F'
      output: extract_metadata_from_filepath
      source_identifier: attributes["log.file.path"]
      type: recombine
    - id: extract_metadata_from_filepath
      parse_from: attributes["log.file.path"]
      regex: ^.*\/(?P<namespace>[^_]+)_(?P<pod_name>[^_]+)_(?P<uid>[a-f0-9\-]+)\/(?P<container_name>[^\._]+)\/(?P<restart_count>\d+)\.log$
      type: regex_parser
    - from: attributes.stream
      to: attributes["log.iostream"]
      type: move
    - from: attributes.container_name
      to: resource["k8s.container.name"]
      type: move
    - from: attributes.namespace
      to: resource["k8s.namespace.name"]
      type: move
    - from: attributes.pod_name
      to: resource["k8s.pod.name"]
      type: move
    - from: attributes.restart_count
      to: resource["k8s.container.restart_count"]
      type: move
    - from: attributes.uid
      to: resource["k8s.pod.uid"]
      type: move
    - from: attributes.log
      to: body
      type: move
    - if: body matches "^{.*}$"
      parse_from: body
      parse_to: attributes
      type: json_parser
    - from: body
      to: attributes.original
      type: copy
    - from: attributes.message
      if: attributes.message != nil
      to: body
      type: move
    - from: attributes.msg
      if: attributes.msg != nil
      to: body
      type: move
    - if: attributes.level != nil
      parse_from: attributes.level
      type: severity_parser
    retry_on_failure:
      enabled: true
    start_at: beginning
    storage: file_storage

service:
  extensions:
  - health_check
  - pprof
  - file_storage
  pipelines:
    logs:
      exporters:
      - otlp
      processors:
      - memory_limiter
      - transform/set-instrumentation-scope-runtime
      receivers:
      - filelog
  telemetry:
    metrics:
      address: ${MY_POD_IP}:8888

```

#### Things to take into consideration (at implementation)
- Dynamically inclusion/exclusion of namespaces, based on LogPipeline spec attributes
- Exclude FluentBit container in OTel configuration and OTel container in FluentBit configuration
- `receivers/filelog/operators`: The copy body to `attributes.original` must be avoided if `dropLogRawBody` flag is enabled


### HostPath VolumeMount Daemonset spec

TODO

## Benchmarking and Performance Tests

TODO