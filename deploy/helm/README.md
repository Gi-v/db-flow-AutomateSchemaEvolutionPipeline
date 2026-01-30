# Mini DBMS Helm Chart

This Helm chart deploys the Mini Relational DBMS to Kubernetes.

## Installation

### Install from local chart

```bash
# Install with default values
helm install mini-dbms ./deploy/helm

# Install with custom values
helm install mini-dbms ./deploy/helm --values my-values.yaml
```

### Upgrade

```bash
helm upgrade mini-dbms ./deploy/helm
```

### Uninstall

```bash
helm uninstall mini-dbms
```

## Configuration

The following table lists the configurable parameters:

| Parameter | Description | Default |
|-----------|-------------|---------|
| `replicaCount` | Number of replicas | `1` |
| `image.repository` | Image repository | `mini-dbms` |
| `image.tag` | Image tag | `latest` |
| `service.type` | Service type | `ClusterIP` |
| `service.port` | Service port | `8080` |
| `resources.limits.cpu` | CPU limit | `500m` |
| `resources.limits.memory` | Memory limit | `512Mi` |
| `config.logLevel` | Log level | `info` |

## Examples

### Production deployment with 3 replicas

```bash
helm install mini-dbms ./deploy/helm \
  --set replicaCount=3 \
  --set config.logLevel=warn
```

### Enable autoscaling

```bash
helm install mini-dbms ./deploy/helm \
  --set autoscaling.enabled=true \
  --set autoscaling.minReplicas=2 \
  --set autoscaling.maxReplicas=10
```
