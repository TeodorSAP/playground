# OTel LogPipeline set-up validation

## 1. Set-up configuration steps

```
k create ns otel-logging

k apply -f /Users/I565663/Projects/telemetry-manager/config/samples/operator_v1alpha1_telemetry.yaml

kh-cls-log // Execute knowledge-hub/scripts/create_cls_log_pipeline.sh with the corresponding environment variables 

helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts

tm && helm install -n otel-logging logging open-telemetry/opentelemetry-collector -f docs/contributor/pocs/assets/otel-logs-values.yaml

tm && helm install -n otel-logging logging open-telemetry/opentelemetry-collector -f ./docs/contributor/pocs/assets/otel-log-agent-values.yaml
```

## 2. Resulting Resources

### Agent ConfigMap

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
      readers:
      - pull:
          exporter:
            prometheus:
              host: ${MY_POD_IP}
              port: 8888

```

#### Things to take into consideration (at implementation)
- Dynamically inclusion/exclusion of namespaces, based on LogPipeline spec attributes
- Exclude FluentBit container in OTel configuration and OTel container in FluentBit configuration
- `receivers/filelog/operators`: The copy body to `attributes.original` must be avoided if `dropLogRawBody` flag is enabled


### Agent Daemonset

```
apiVersion: apps/v1
kind: DaemonSet
metadata:
  annotations:
    deprecated.daemonset.template.generation: "1"
    meta.helm.sh/release-name: logging
    meta.helm.sh/release-namespace: otel-logging
  labels:
    app.kubernetes.io/component: agent-collector
    app.kubernetes.io/instance: logging
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: opentelemetry-collector
    app.kubernetes.io/version: 0.114.0
    helm.sh/chart: opentelemetry-collector-0.110.3
  name: logging-opentelemetry-collector-agent
  namespace: otel-logging
spec:
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app.kubernetes.io/instance: logging
      app.kubernetes.io/name: opentelemetry-collector
      component: agent-collector
  template:
    metadata:
      annotations:
        checksum/config: bae3842cd432b970200d5a39fd42e45bd45459a3f89cda4a87939d99d5481fba
      creationTimestamp: null
      labels:
        app.kubernetes.io/instance: logging
        app.kubernetes.io/name: opentelemetry-collector
        component: agent-collector
    spec:
      containers:
      - args:
        - --config=/conf/relay.yaml
        env:
        - name: MY_POD_IP
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: status.podIP
        image: otel/opentelemetry-collector-k8s:0.114.0
        imagePullPolicy: Always
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /
            port: 13133
            scheme: HTTP
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
        name: opentelemetry-collector
        ports:
        - containerPort: 6831
          hostPort: 6831
          name: jaeger-compact
          protocol: UDP
        - containerPort: 14250
          hostPort: 14250
          name: jaeger-grpc
          protocol: TCP
        - containerPort: 14268
          hostPort: 14268
          name: jaeger-thrift
          protocol: TCP
        - containerPort: 8888
          name: metrics
          protocol: TCP
        - containerPort: 4317
          hostPort: 4317
          name: otlp
          protocol: TCP
        - containerPort: 4318
          hostPort: 4318
          name: otlp-http
          protocol: TCP
        - containerPort: 9411
          hostPort: 9411
          name: zipkin
          protocol: TCP
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /
            port: 13133
            scheme: HTTP
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1
        resources: {}
        securityContext:
          runAsGroup: 0
          runAsUser: 0
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /conf
          name: opentelemetry-collector-configmap
        - mountPath: /var/log/pods
          name: varlogpods
          readOnly: true
        - mountPath: /var/lib/docker/containers
          name: varlibdockercontainers
          readOnly: true
        - mountPath: /var/lib/otelcol
          name: varlibotelcol
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      serviceAccount: logging-opentelemetry-collector
      serviceAccountName: logging-opentelemetry-collector
      terminationGracePeriodSeconds: 30
      volumes:
      - configMap:
          defaultMode: 420
          items:
          - key: relay
            path: relay.yaml
          name: logging-opentelemetry-collector-agent
        name: opentelemetry-collector-configmap
      - hostPath:
          path: /var/log/pods
          type: ""
        name: varlogpods
      - hostPath:
          path: /var/lib/otelcol
          type: DirectoryOrCreate
        name: varlibotelcol
      - hostPath:
          path: /var/lib/docker/containers
          type: ""
        name: varlibdockercontainers
  updateStrategy:
    rollingUpdate:
      maxSurge: 0
      maxUnavailable: 1
    type: RollingUpdate
```

### How does checkpointing work

- By enabling the storeCheckpoint preset (Helm) the `file_storage` extension is activated in the receiver
- The `file_storage` has the path `/var/lib/otelcol`
- This is later mounted as a `hostPath` volume in the DaemonSet spec
- Also set in the `storage` property of the filelog receiver

> `storage` = The ID of a storage extension to be used to store file offsets. File offsets allow the receiver to pick up where it left off in the case of a collector restart. If no storage extension is used, the receiver will manage offsets in memory only.

## 3. Benchmarking and Performance Tests

TODO