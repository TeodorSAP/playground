### Base Setup
1. `istio` module deployed
2. Change `kyma-system/istio-controller-manager` Deployment image to: `europe-central2-docker.pkg.dev/sap-se-cx-kyma-goat/istio/istio-manager:PR-1374`
3. Create load-test namespace: `kubectl create ns load-test`

### OTEL Demo App Test
https://github.com/kyma-project/telemetry-manager/tree/main/docs/user/integration/opentelemetry-demo

Deploy:
```shell
export K8S_NAMESPACE="otel"
kubectl create namespace $K8S_NAMESPACE
kubectl label namespace $K8S_NAMESPACE istio-injection=enabled

export HELM_OTEL_RELEASE="otel-demo"
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update

helm upgrade --install --create-namespace -n $K8S_NAMESPACE $HELM_OTEL_RELEASE open-telemetry/opentelemetry-demo -f https://raw.githubusercontent.com/kyma-project/telemetry-manager/main/docs/user/integration/opentelemetry-demo/values.yaml
```

Test:
```shell
kubectl -n $K8S_NAMESPACE port-forward svc/frontend-proxy 8080
```

Clean Up:
```shell
helm delete -n $K8S_NAMESPACE $HELM_OTEL_RELEASE
```

### Prometheus + Grafana Setup
```shell
k create ns prometheus
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm upgrade --install -n "prometheus" "prometheus" prometheus-community/kube-prometheus-stack -f ./assets/prom-values.yaml --set grafana.adminPassword=myPwd
```

### Load Test Setup
```shell
kubectl label namespace load-test istio-injection=enabled

# Test envoy provider
kubectl apply -f ./assets/istio-telemetry-envoy.yaml
kubectl apply -f ./assets/load-test.yaml

# Test kyma-logs provider
kubectl apply -f ./assets/backend.yaml
kubectl apply -f ./assets/istio-telemetry-kyma-logs.yaml
kubectl apply -f ./assets/load-test.yaml

# Additionally, test edge-cases for kyma-logs provider (did not work-out)
# kubectl apply -f ./assets/fault-injection-config.yaml # configure fault injection parameters accordingly (see file)
# kubectl apply -f ./assets/istio-telemetry-kyma-logs.yaml
# kubectl apply -f ./assets/load-test.yaml
```

1. Test "backend unreachable" scenario by scaling the backend deployment to 0 replicas
2. Test "backend backpressure" scenario by modifying the ConfigMap of the backend, i.e. adding a new memory_limiter processor:
```yaml
    processors:
      memory_limiter:
        limit_mib: 40
        spike_limit_mib: 20
        check_interval: 1s
    # ...
    service:
      pipelines:
        logs:
          receivers:
            - otlp
          processors:
            - memory_limiter
          exporters:
            - debug
```

![Load Test Architecture](./load-test-architecture.drawio.svg)