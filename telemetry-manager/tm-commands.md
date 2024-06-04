# Telemetry Manager - My Commands

### Fully delete k3d cluster

```bash
k3d cluster delete kyma && k3d registry delete k3d-kyma-registry
```

### Deploy image using Docker Hub

```bash
# Updating Image
export IMG=teodorsap/telemetry-manager:current
make docker-build
make docker-push
# Provisioning
make provision-k3d
make install
make deploy
```

- Set `imagePullPolicy` from `IfNew` to `Always`
