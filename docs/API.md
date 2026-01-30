# API Documentation

## Base URL

```
http://localhost:8080
```

## Endpoints

### 1. Health Check

Check if the service is healthy and running.

**Endpoint**: `GET /health`

**Response**:
```json
{
  "status": "healthy",
  "service": "minidb"
}
```

**Example**:
```bash
curl http://localhost:8080/health
```

---

### 2. API Information

Get information about the API and supported features.

**Endpoint**: `GET /api/info`

**Response**:
```json
{
  "name": "Mini Relational DBMS",
  "version": "0.1.0",
  "features": [
    "SQL: SELECT, INSERT, UPDATE, DELETE",
    "Storage: Disk-backed pages with buffer pool",
    "Indexing: B+ tree",
    "Transactions: WAL and MVCC",
    "Concurrency: Lock manager"
  ]
}
```

**Example**:
```bash
curl http://localhost:8080/api/info
```

---

### 3. Execute SQL Query

Execute a SQL query and get results.

**Endpoint**: `POST /query`

**Request Headers**:
```
Content-Type: application/json
```

**Request Body**:
```json
{
  "sql": "SELECT * FROM users WHERE id = 1"
}
```

**Response (Success)**:
```json
{
  "success": true,
  "result": {
    "rows": [
      {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com"
      }
    ],
    "rows_affected": 0,
    "message": "Selected 1 rows from users"
  }
}
```

**Response (Error)**:
```json
{
  "error": "Parse error: unexpected token"
}
```

**HTTP Status Codes**:
- `200 OK`: Query executed successfully
- `400 Bad Request`: Invalid SQL syntax or request format
- `500 Internal Server Error`: Execution error

---

### 4. Prometheus Metrics

Get Prometheus metrics for monitoring.

**Endpoint**: `GET /metrics`

**Response**:
```
# HELP minidb_requests_total Total number of requests
# TYPE minidb_requests_total counter
minidb_requests_total{endpoint="/query",method="POST"} 42

# HELP minidb_query_duration_seconds Query execution duration in seconds
# TYPE minidb_query_duration_seconds histogram
minidb_query_duration_seconds_bucket{query_type="SELECT",le="0.005"} 10
minidb_query_duration_seconds_bucket{query_type="SELECT",le="0.01"} 25
...
```

**Example**:
```bash
curl http://localhost:8080/metrics
```

---

## SQL Query Examples

### CREATE TABLE

Create a new table with a schema.

```json
{
  "sql": "CREATE TABLE users (id int, name string)"
}
```

**Response**:
```json
{
  "success": true,
  "result": {
    "message": "Created table users"
  }
}
```

---

### INSERT

Insert data into a table.

```json
{
  "sql": "INSERT INTO users VALUES (1, 'John Doe')"
}
```

**Response**:
```json
{
  "success": true,
  "result": {
    "rows_affected": 1,
    "message": "Inserted 1 row(s) into users"
  }
}
```

---

### SELECT

Query data from a table.

```json
{
  "sql": "SELECT * FROM users"
}
```

**Response**:
```json
{
  "success": true,
  "result": {
    "rows": [
      {"id": 1, "name": "John Doe"}
    ],
    "message": "Selected 1 rows from users"
  }
}
```

**With WHERE clause**:
```json
{
  "sql": "SELECT * FROM users WHERE id = 1"
}
```

---

### UPDATE

Update existing records.

```json
{
  "sql": "UPDATE users SET name = 'Jane Doe' WHERE id = 1"
}
```

**Response**:
```json
{
  "success": true,
  "result": {
    "rows_affected": 1,
    "message": "Updated 1 row(s) in users"
  }
}
```

---

### DELETE

Delete records from a table.

```json
{
  "sql": "DELETE FROM users WHERE id = 1"
}
```

**Response**:
```json
{
  "success": true,
  "result": {
    "rows_affected": 1,
    "message": "Deleted 1 row(s) from users"
  }
}
```

---

### DROP TABLE

Remove a table.

```json
{
  "sql": "DROP TABLE users"
}
```

**Response**:
```json
{
  "success": true,
  "result": {
    "message": "Dropped table users"
  }
}
```

---

## Error Handling

### Parse Errors

**Request**:
```json
{
  "sql": "SELEC * FROM users"
}
```

**Response** (400):
```json
{
  "error": "Parse error: unexpected token: SELEC"
}
```

---

### Execution Errors

**Request**:
```json
{
  "sql": "SELECT * FROM nonexistent_table"
}
```

**Response** (500):
```json
{
  "error": "Execution error: table nonexistent_table does not exist"
}
```

---

### Missing SQL

**Request**:
```json
{}
```

**Response** (400):
```json
{
  "error": "Key: 'sql' Error:Field validation for 'sql' failed on the 'required' tag"
}
```

---

## Client Examples

### cURL

```bash
# Health check
curl http://localhost:8080/health

# Execute query
curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{"sql": "SELECT * FROM users"}'
```

### Python

```python
import requests

# Execute query
response = requests.post('http://localhost:8080/query', json={
    'sql': 'SELECT * FROM users WHERE id = 1'
})

if response.status_code == 200:
    result = response.json()
    print(result['result']['rows'])
else:
    print(f"Error: {response.json()['error']}")
```

### JavaScript (Node.js)

```javascript
const axios = require('axios');

async function executeQuery(sql) {
  try {
    const response = await axios.post('http://localhost:8080/query', {
      sql: sql
    });
    return response.data.result;
  } catch (error) {
    console.error('Error:', error.response.data.error);
  }
}

executeQuery('SELECT * FROM users').then(result => {
  console.log(result.rows);
});
```

### Go

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type QueryRequest struct {
    SQL string `json:"sql"`
}

type QueryResponse struct {
    Success bool        `json:"success"`
    Result  interface{} `json:"result"`
    Error   string      `json:"error,omitempty"`
}

func executeQuery(sql string) (*QueryResponse, error) {
    reqBody, _ := json.Marshal(QueryRequest{SQL: sql})
    
    resp, err := http.Post(
        "http://localhost:8080/query",
        "application/json",
        bytes.NewBuffer(reqBody),
    )
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var result QueryResponse
    json.NewDecoder(resp.Body).Decode(&result)
    
    return &result, nil
}

func main() {
    result, _ := executeQuery("SELECT * FROM users")
    fmt.Printf("%+v\n", result)
}
```

---

## Rate Limiting

Currently, there is no rate limiting implemented. This is a future enhancement.

---

## Authentication

Currently, there is no authentication required. This is a future enhancement for production deployments.

---

## Monitoring Integration

The `/metrics` endpoint exposes Prometheus-compatible metrics that can be scraped by a Prometheus server.

**Metrics Available**:
- `minidb_requests_total`: Counter of total requests by method and endpoint
- `minidb_query_duration_seconds`: Histogram of query execution times by query type

**Prometheus Configuration**:
```yaml
scrape_configs:
  - job_name: 'minidb'
    static_configs:
      - targets: ['localhost:8080']
    metrics_path: '/metrics'
```

---

## Future Enhancements

- Batch query execution
- Prepared statements
- Transaction control (BEGIN, COMMIT, ROLLBACK)
- Authentication and authorization
- Query result pagination
- WebSocket support for streaming results
- GraphQL interface
