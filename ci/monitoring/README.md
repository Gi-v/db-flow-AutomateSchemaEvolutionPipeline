# Monitoring Setup

This directory contains configuration for monitoring the Mini Relational DBMS.

## Prometheus

Prometheus is used to collect metrics from the DBMS.

### Setup

```bash
# Run Prometheus with docker-compose (add to docker-compose.yml)
# Or deploy to Kubernetes
kubectl apply -f prometheus-deployment.yaml
```

### Metrics to Monitor

- **Query Performance**
  - `dbms_query_duration_seconds`: Query execution time
  - `dbms_query_total`: Total queries executed
  
- **Storage**
  - `dbms_pages_allocated`: Number of pages allocated
  - `dbms_buffer_pool_hit_ratio`: Buffer pool efficiency
  
- **Transactions**
  - `dbms_transactions_total`: Total transactions
  - `dbms_transactions_committed`: Committed transactions
  - `dbms_transactions_aborted`: Aborted transactions
  
- **System**
  - `go_goroutines`: Number of goroutines
  - `go_memstats_alloc_bytes`: Memory allocated

## Grafana

Grafana provides visualization for Prometheus metrics.

### Dashboards

1. **Overview Dashboard**: System health and performance
2. **Query Performance**: Query latency and throughput
3. **Storage Metrics**: Page usage and buffer pool stats
4. **Transaction Metrics**: Transaction rates and states

### Access

```bash
# Port forward Grafana
kubectl port-forward svc/grafana 3000:3000

# Default credentials: admin/admin
```

## Future Enhancements

- Alert rules for critical metrics
- Custom Grafana dashboards
- Log aggregation with Loki
- Distributed tracing with Jaeger
