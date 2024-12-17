# OTel LogPipeline set-up validation

## 1. Set-up configuration steps

### With Helm

``` bash
k apply -f telemetry-manager/config/samples/operator_v1alpha1_telemetry.yaml

kh-cls-log // Execute knowledge-hub/scripts/create_cls_log_pipeline.sh with the corresponding environment variables 

helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts

tm && helm install -n kyma-system logging open-telemetry/opentelemetry-collector -f ./docs/contributor/pocs/assets/otel-log-agent-values.yaml
```

### Manual

``` bash
k apply -f telemetry-manager/config/samples/operator_v1alpha1_telemetry.yaml

kh-cls-log // Execute knowledge-hub/scripts/create_cls_log_pipeline.sh with the corresponding environment variables 

k apply -f ./otlp-logs-validation.yaml
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

TODO: Performance test: overall throughput si comparable to what we have with FB
TODO: Check backpressure scenario: That agent is pausing if backend is slow/rejecting

``` bash
k create ns prometheus
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm upgrade --install -n "prometheus" "prometheus" prometheus-community/kube-prometheus-stack -f hack/load-tests/values.yaml --set grafana.adminPassword=myPwd

k apply -f telemetry-manager/hack/load-tests/log-agent-test-setup.yaml
```

``` sql
-- RECEIVED
round(sum(rate(otelcol_receiver_accepted_log_records{service="telemetry-log-agent-metrics"}[20m])))

-- EXPORTED
round(sum(rate(otelcol_exporter_sent_log_records{service="telemetry-log-agent-metrics"}[20m])))

-- QUEUE
avg(sum(otelcol_exporter_queue_size{service="telemetry-log-agent-metrics"}))

-- MEMORY
round(sum(avg_over_time(container_memory_working_set_bytes{namespace="kyma-system", container="collector"}[20m]) * on(namespace,pod) group_left(workload) avg_over_time(namespace_workload_pod:kube_pod_owner:relabel{namespace="kyma-system", workload="telemetry-log-agent"}[20m])) by (pod) / 1024 / 1024)

-- CPU
round(sum(avg_over_time(node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{namespace="kyma-system"}[20m]) * on(namespace,pod) group_left(workload) avg_over_time(namespace_workload_pod:kube_pod_owner:relabel{namespace="kyma-system", workload="telemetry-log-agent"}[20m])) by (pod), 0.1)
```