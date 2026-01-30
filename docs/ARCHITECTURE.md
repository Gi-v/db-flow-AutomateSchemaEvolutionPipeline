# System Architecture

## Overview

The Mini Relational DBMS is designed as a layered architecture, with each layer responsible for specific functionality. This document describes the system architecture, component interactions, and design decisions.

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────┐
│                     Client Layer                         │
│           (REST API, HTTP Requests)                     │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                  API Layer                               │
│          - Request handling (Gin)                       │
│          - Authentication (future)                      │
│          - Metrics collection                           │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                Parser Layer                              │
│          - Lexical analysis (Lexer)                     │
│          - Syntax analysis (Parser)                     │
│          - AST generation                               │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│              Execution Layer                             │
│          - Query optimization (future)                  │
│          - Execution planning                           │
│          - Result generation                            │
└─────┬──────────┬──────────┬──────────┬─────────────────┘
      │          │          │          │
      ▼          ▼          ▼          ▼
┌──────────┐ ┌────────┐ ┌─────────┐ ┌──────────┐
│ Storage  │ │ Index  │ │   WAL   │ │   Lock   │
│ Manager  │ │Manager │ │ Manager │ │ Manager  │
└────┬─────┘ └────────┘ └─────────┘ └──────────┘
     │
┌────▼──────────────────────────────────────────────────┐
│              Buffer Pool Manager                       │
│          - Page caching                                │
│          - LRU eviction                                │
│          - Pin/unpin management                        │
└────┬──────────────────────────────────────────────────┘
     │
┌────▼──────────────────────────────────────────────────┐
│              Disk Manager                              │
│          - Page I/O                                    │
│          - File management                             │
└───────────────────────────────────────────────────────┘
```

## Component Details

### 1. API Layer

**Responsibility**: Handle HTTP requests and responses

**Components**:
- Gin router for REST endpoints
- Request validation
- Response formatting
- Prometheus metrics instrumentation

**Endpoints**:
- `POST /query` - Execute SQL queries
- `GET /health` - Health check
- `GET /metrics` - Prometheus metrics
- `GET /api/info` - API information

### 2. Parser Layer

**Responsibility**: Convert SQL text to executable AST

**Components**:
- **Lexer**: Tokenizes SQL input
  - Keywords (SELECT, INSERT, etc.)
  - Identifiers (table names, column names)
  - Operators (=, <, >, etc.)
  - Literals (strings, numbers)

- **Parser**: Builds Abstract Syntax Tree
  - Recursive descent parsing
  - Syntax validation
  - Error reporting

- **AST Nodes**:
  - SelectStatement
  - InsertStatement
  - UpdateStatement
  - DeleteStatement
  - CreateTableStatement
  - DropTableStatement

### 3. Execution Layer

**Responsibility**: Execute queries and return results

**Components**:
- Query executor
- Transaction coordinator
- Result builder

**Operations**:
1. Acquire locks (via Lock Manager)
2. Log operation (via WAL Manager)
3. Execute query (via Storage/Index Manager)
4. Build result set
5. Release locks

### 4. Storage Manager

**Responsibility**: Manage tables and tuples

**Components**:
- **TableManager**: Create/drop tables, maintain schemas
- **Table**: Store and retrieve tuples
- **Schema**: Define column types and constraints
- **Tuple**: Represent rows of data

**Features**:
- Table creation/deletion
- Tuple insertion
- Table scanning
- Schema management

### 5. Buffer Pool Manager

**Responsibility**: Cache pages in memory

**Components**:
- Page cache (map of PageID → Page)
- LRU eviction policy
- Pin/unpin reference counting
- Dirty page tracking

**Page Management**:
- 8KB page size
- LRU eviction when full
- Pin counting prevents eviction
- Flush dirty pages to disk

**Algorithm**:
```
FetchPage(pageID):
  if page in cache:
    pin page
    update LRU
    return page
  if cache is full:
    evict unpinned page
  load page from disk
  pin page
  add to cache
  return page
```

### 6. Index Manager (B+ Tree)

**Responsibility**: Provide fast data access via indexing

**Components**:
- BTree structure (order 4)
- Internal nodes (keys + child pointers)
- Leaf nodes (keys + values + next pointer)
- Index catalog

**Features**:
- Point queries: O(log n)
- Range scans: O(log n + k)
- Insert/Delete: O(log n)
- Leaf node linking for range queries

### 7. WAL Manager

**Responsibility**: Ensure durability and recovery

**Components**:
- Write-Ahead Log entries
- Log buffer
- LSN (Log Sequence Number) tracking

**Log Entry Types**:
- BEGIN - Transaction start
- INSERT - Insert operation
- UPDATE - Update operation
- DELETE - Delete operation
- COMMIT - Transaction commit
- ABORT - Transaction rollback

**Protocol**:
1. Log changes before applying them
2. Flush log on commit
3. Checkpoint periodically
4. Replay log on recovery

### 8. Lock Manager

**Responsibility**: Coordinate concurrent access

**Components**:
- Lock table (resource → locks)
- Wait queue for blocked transactions
- Deadlock detection (future)

**Lock Types**:
- Read locks (shared)
- Write locks (exclusive)

**Protocol**:
- Multiple read locks allowed
- Write lock is exclusive
- Release locks on commit/abort

## Data Flow

### SELECT Query Flow

```
1. Client sends POST /query with SQL
2. API layer validates request
3. Lexer tokenizes SQL string
4. Parser builds SelectStatement AST
5. Executor acquires read lock on table
6. Executor scans table via Storage Manager
7. Storage Manager fetches pages from Buffer Pool
8. Buffer Pool loads pages from disk if needed
9. Results returned through layers
10. Lock released
```

### INSERT Query Flow

```
1. Client sends POST /query with INSERT SQL
2. API layer validates request
3. Lexer/Parser build InsertStatement AST
4. Executor acquires write lock on table
5. WAL Manager logs INSERT operation
6. Storage Manager inserts tuple into page
7. Buffer Pool marks page as dirty
8. Index Manager updates indexes (future)
9. Result returned to client
10. Lock released
11. Background: dirty pages flushed to disk
```

## Design Decisions

### 1. Page-Based Storage
- **Why**: Standard in DBMS, enables efficient caching
- **Trade-off**: Fixed 8KB size may waste space for small records

### 2. LRU Eviction
- **Why**: Simple, effective for most workloads
- **Trade-off**: Not optimal for sequential scans

### 3. REST API
- **Why**: Simple, language-agnostic, HTTP-based
- **Trade-off**: Higher overhead than binary protocol

### 4. Write-Ahead Logging
- **Why**: Industry standard for durability
- **Trade-off**: Double I/O (log + data)

### 5. Go Language
- **Why**: Strong concurrency, good performance, simple deployment
- **Trade-off**: GC pauses (mitigated by tuning)

## Scalability Considerations

### Current Limitations
- Single-node only
- Limited query optimizer
- No query parallelism
- No distributed transactions

### Future Enhancements
- Query optimizer with cost-based planning
- Parallel query execution
- Distributed storage (sharding)
- Replication for high availability
- Advanced indexing (hash, GiST)

## Performance Characteristics

### Buffer Pool
- Hit rate: Critical for performance
- Size: Configurable (default 100 pages = 800KB)
- Eviction: O(n) for LRU update (can optimize with doubly-linked list)

### B+ Tree
- Search: O(log n)
- Insert: O(log n)
- Range scan: O(log n + k) where k = result size

### Concurrency
- Lock granularity: Table-level (can optimize to row-level)
- Deadlock detection: Not yet implemented

## Monitoring

### Key Metrics
- `minidb_requests_total`: Total requests by endpoint
- `minidb_query_duration_seconds`: Query latency histogram
- Buffer pool hit rate (future)
- Transaction throughput (future)

### Observability
- Prometheus for metrics collection
- Grafana for visualization
- Structured logging (future)
- Distributed tracing (future)

## Security

### Current State
- No authentication/authorization
- No encryption at rest
- No encryption in transit (HTTP)

### Future Enhancements
- JWT-based authentication
- Role-based access control
- TLS for API
- Encryption at rest
- SQL injection prevention (parameterized queries)

## Testing Strategy

### Unit Tests
- Component-level testing
- Mock dependencies
- Edge case coverage

### Integration Tests
- End-to-end API testing
- Multi-component interaction
- Concurrency testing

### Performance Tests (Future)
- Benchmark queries
- Load testing
- Stress testing

## Deployment Architecture

### Local Development
```
Developer → Docker Container → MiniDB
```

### Kubernetes Production
```
Client → LoadBalancer → K8s Service → MiniDB Pods (3 replicas)
                                    ↓
                              Prometheus (metrics)
                                    ↓
                              Grafana (dashboards)
```

## References

- [PostgreSQL Internals](https://www.postgresql.org/docs/current/internals.html)
- [Database Internals Book](https://www.databass.dev/)
- [CMU 15-445 Database Systems](https://15445.courses.cs.cmu.edu/)
