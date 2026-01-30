# Kubernetes Manifests

This directory contains Kubernetes manifests for deploying the Mini Relational DBMS.

## Files

- `deployment.yaml`: Deployment specification for the DBMS pods
- `service.yaml`: Service to expose the DBMS
- `configmap.yaml`: Configuration for the application

## Deployment

### Deploy to Kubernetes

```bash
# Apply all manifests
kubectl apply -f deploy/k8s/

# Or apply individually
kubectl apply -f deploy/k8s/configmap.yaml
kubectl apply -f deploy/k8s/deployment.yaml
kubectl apply -f deploy/k8s/service.yaml
```

### Verify Deployment

```bash
# Check pods
kubectl get pods -l app=mini-dbms

# Check service
kubectl get svc mini-dbms

# View logs
kubectl logs -l app=mini-dbms -f
```

### Access the Service

```bash
# Port forward to access locally
kubectl port-forward svc/mini-dbms 8080:8080

# Access health endpoint
curl http://localhost:8080/health
```

### Scale Deployment

```bash
# Scale to 3 replicas
kubectl scale deployment mini-dbms --replicas=3

# Verify scaling
kubectl get pods -l app=mini-dbms
```

### Delete Deployment

```bash
# Delete all resources
kubectl delete -f deploy/k8s/
```

## Production Considerations

For production deployments, consider:
- Using an Ingress controller for external access
- Configuring persistent volumes for data storage
- Setting up resource quotas and limits
- Implementing horizontal pod autoscaling
- Adding monitoring and alerting
- Using Helm charts for easier management (see `deploy/helm/`)
