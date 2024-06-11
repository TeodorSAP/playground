# KCP - Useful Commands

### Get all clusters
```bash
kcp rt
```

### Run tasks on clusters
```bash
kcp taskrun --target <TARGET> -p <PARALLEL_PROCESSES> --gardener-kubeconfig <GARDENER_ENVIRONMENT_KUBECONFIG> <SCRIPT>
```

Example:
```bash
kcp taskrun --target all -p 8 --gardener-kubeconfig "$HOME/.kcp/gardener-kubeconfig/kubeconfig-garden-kyma-dev.yaml" /Users/I565663/Projects/playground/telemetry-manager/scripts/check-ca-otlp.sh
```