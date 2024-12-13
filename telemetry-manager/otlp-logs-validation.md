# OTel LogPipeline set-up validation

## 1. Set-up configuration steps

### With Helm

``` bash
k apply -f /Users/I565663/Projects/telemetry-manager/config/samples/operator_v1alpha1_telemetry.yaml

kh-cls-log // Execute knowledge-hub/scripts/create_cls_log_pipeline.sh with the corresponding environment variables 

helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts

tm && helm install -n kyma-system logging open-telemetry/opentelemetry-collector -f ./docs/contributor/pocs/assets/otel-log-agent-values.yaml
```

### Manual

```
```

## 2. Resulting Resources

### Agent ConfigMap (OTel Config)

See [OTLP Logs Validation YAML](./otlp-logs-validation.yaml)

#### Things to take into consideration (at implementation)
- Dynamically inclusion/exclusion of namespaces, based on LogPipeline spec attributes
- Exclude FluentBit container in OTel configuration and OTel container in FluentBit configuration
- `receivers/filelog/operators`: The copy body to `attributes.original` must be avoided if `dropLogRawBody` flag is enabled


### Agent DaemonSet

See [OTLP Logs Validation YAML](./otlp-logs-validation.yaml)

### How does checkpointing work

- By enabling the storeCheckpoint preset (Helm) the `file_storage` extension is activated in the receiver
- The `file_storage` has the path `/var/lib/otelcol`
- This is later mounted as a `hostPath` volume in the DaemonSet spec
- Also set in the `storage` property of the filelog receiver

> `storage` = The ID of a storage extension to be used to store file offsets. File offsets allow the receiver to pick up where it left off in the case of a collector restart. If no storage extension is used, the receiver will manage offsets in memory only.

## 3. Benchmarking and Performance Tests

TODO