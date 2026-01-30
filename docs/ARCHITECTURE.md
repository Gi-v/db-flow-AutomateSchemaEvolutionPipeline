# Mini Relational DBMS - Architecture Documentation

## Overview

This document describes the architecture and design decisions for the Mini Relational DBMS project. The system is designed as a simplified relational database management system with fundamental DBMS capabilities including SQL processing, disk-backed storage, indexing, and transaction support.

## System Architecture

### High-Level Components

The system is organized into distinct layers, each responsible for specific functionality:

```
┌─────────────────────────────────────────────────────────┐
│                      REST API Layer                      │
│  - HTTP endpoints for database operations                │
│  - Request validation and response formatting            │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                    SQL Parser Layer                      │
│  - Lexical analysis (tokenization)                       │
│  - Syntactic analysis (AST construction)                 │
│  - Semantic validation                                   │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                 Query Executor Layer                     │
│  - Query planning and optimization                       │
│  - Operator execution (scan, join, filter)               │
│  - Result set construction                               │
└────────────────────┬────────────────────────────────────┘
                     │
          ┌──────────┴──────────┐
          │                     │
┌─────────▼────────┐   ┌───────▼──────────┐
│   Index Layer    │   │   Storage Layer   │
│  - B+ Tree impl. │   │  - Page manager   │
│  - Index scan    │   │  - Tuple format   │
│  - Key lookups   │   │  - Buffer pool    │
└─────────┬────────┘   └───────┬──────────┘
          │                     │
          └──────────┬──────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                  Transaction Layer                       │
│  - Transaction management (Begin/Commit/Abort)           │
│  - Write-Ahead Logging (WAL)                             │
│  - Lock management and concurrency control               │
│  - ACID guarantees and crash recovery                    │
└─────────────────────────────────────────────────────────┘
```

## Component Details

### 1. REST API Layer (`src/api/`)

**Purpose**: Provides HTTP interface for database operations

**Key Components**:
- `router.go`: HTTP request routing and handling
- Health check endpoint (`/health`)
- Query endpoint (`/api/query`)

**Design Decisions**:
- Uses standard Go `net/http` for simplicity and minimal dependencies
- Future: May integrate gin or chi for advanced routing
- RESTful design for easy client integration

### 2. SQL Parser (`src/parser/`)

**Purpose**: Converts SQL text into structured representation

**Key Components**:
- `parser.go`: Main parser logic
- Statement types: SELECT, INSERT, UPDATE, DELETE

**Implementation Strategy** (To Be Implemented):
1. **Lexer**: Tokenize SQL input
2. **Parser**: Build Abstract Syntax Tree (AST)
3. **Validator**: Semantic validation of queries

**Design Decisions**:
- Custom parser for educational purposes
- Alternative: Could use existing parser libraries
- AST representation allows for query optimization

### 3. Query Executor (`src/executor/`)

**Purpose**: Executes parsed SQL statements

**Key Components**:
- `executor.go`: Query execution engine

**Execution Model** (To Be Implemented):
- **Volcano Model**: Iterator-based execution
- Operators: TableScan, IndexScan, Filter, Join, Project
- Pipeline execution for efficiency

**Design Decisions**:
- Operator-based design for modularity
- Support for multiple execution strategies
- Integration with storage and index layers

### 4. Storage Layer (`src/storage/`)

**Purpose**: Manages disk-backed data storage

**Key Components**:
- `page.go`: Page abstraction (4KB pages)
- `manager.go`: Storage manager for page allocation/retrieval

**Storage Model**:
```
┌──────────────────────────────────────┐
│              Page (4KB)               │
├──────────────────────────────────────┤
│  Header (metadata, free space, etc)  │
├──────────────────────────────────────┤
│  Tuple 1                             │
│  Tuple 2                             │
│  ...                                 │
│  Tuple N                             │
├──────────────────────────────────────┤
│  Free Space                          │
└──────────────────────────────────────┘
```

**Design Decisions**:
- Fixed 4KB page size (industry standard)
- Slotted page design for variable-length tuples
- In-memory page cache (buffer pool) for performance
- Future: LRU eviction policy for buffer pool

### 5. Index Layer (`src/index/`)

**Purpose**: Provides fast data lookup via B+ tree

**Key Components**:
- `btree.go`: B+ tree implementation

**B+ Tree Properties**:
- **Order**: Configurable (default: 4)
- **Structure**: All data in leaves, internal nodes for routing
- **Operations**: Insert, Search, Delete, Range scan

**Design Decisions**:
- B+ tree chosen for:
  - Efficient range queries
  - Balanced structure (O(log n) operations)
  - High fan-out reduces tree height
- Leaf nodes linked for sequential access

### 6. Transaction Layer (`src/transaction/`)

**Purpose**: Ensures ACID properties and crash recovery

**Key Components**:
- `transaction.go`: Transaction manager
- `wal.go`: Write-Ahead Logging

#### Transaction Manager

**Transaction States**:
- Active: Transaction in progress
- Committed: Successfully committed
- Aborted: Rolled back

**Operations**:
- `Begin()`: Start new transaction
- `Commit()`: Finalize changes
- `Abort()`: Rollback changes

#### Write-Ahead Logging (WAL)

**Log Record Types**:
- BEGIN: Transaction start
- UPDATE: Data modification
- COMMIT: Transaction commit
- ABORT: Transaction abort

**Recovery Protocol**:
1. **Redo Phase**: Replay committed transactions
2. **Undo Phase**: Rollback uncommitted transactions

**Design Decisions**:
- WAL for crash recovery (ARIES-style)
- Log records persisted before data pages
- Checkpointing for recovery optimization

#### Concurrency Control

**Locking Strategy** (To Be Implemented):
- Two-Phase Locking (2PL)
- Lock types: Shared (S), Exclusive (X)
- Lock granularity: Tuple-level locks
- Deadlock detection via wait-for graph

## Data Flow

### Query Execution Flow

```
SQL Query
   │
   ▼
[Parser] → AST
   │
   ▼
[Executor] → Execution Plan
   │
   ▼
[Storage/Index] → Data Access
   │
   ▼
[Transaction] → ACID Guarantees
   │
   ▼
Result Set
```

### Write Operation Flow

```
INSERT Statement
   │
   ▼
[Parser] → Parsed Statement
   │
   ▼
[Transaction] → Begin Transaction
   │
   ▼
[WAL] → Log Record Written
   │
   ▼
[Storage] → Page Updated
   │
   ▼
[Index] → Index Updated
   │
   ▼
[Transaction] → Commit
```

## Storage Format

### Page Layout

```
Offset 0:    Page ID (4 bytes)
Offset 4:    LSN (8 bytes) - Log Sequence Number
Offset 12:   Free Space Pointer (2 bytes)
Offset 14:   Tuple Count (2 bytes)
Offset 16:   Slot Array (variable)
Offset N:    Tuples (variable, from end)
```

### Tuple Format

```
┌──────────────────────────────────────┐
│  Header (flags, null bitmap)         │
├──────────────────────────────────────┤
│  Field 1 (type-specific encoding)    │
│  Field 2 (type-specific encoding)    │
│  ...                                 │
│  Field N (type-specific encoding)    │
└──────────────────────────────────────┘
```

## Scalability Considerations

### Current Design (Single-Node)
- In-memory buffer pool
- Single-threaded execution (initial)
- File-based storage

### Future Enhancements
- Multi-threaded query execution
- Distributed transactions
- Replication and sharding
- Query result caching

## Performance Optimizations

### Planned Optimizations
1. **Buffer Pool**: LRU cache for hot pages
2. **Index Selection**: Cost-based query optimizer
3. **Parallel Execution**: Multi-threaded operators
4. **Compression**: Page-level compression
5. **Statistics**: Table/index statistics for planning

## Monitoring & Observability

### Metrics (Planned)
- Query latency (p50, p95, p99)
- Transaction throughput
- Buffer pool hit ratio
- Lock contention
- WAL write rate

### Integration
- Prometheus for metrics collection
- Grafana for visualization
- Structured logging for debugging

## Testing Strategy

### Unit Tests
- Per-component testing
- Mock dependencies
- Code coverage > 80%

### Integration Tests
- Multi-component workflows
- Transaction scenarios
- Crash recovery testing

### Acceptance Tests
- End-to-end SQL queries
- Performance benchmarks
- Correctness validation

## Security Considerations

### Current
- No authentication (development only)
- Local-only deployment

### Production Requirements
- Authentication and authorization
- SQL injection prevention (parameterized queries)
- Encrypted connections (TLS)
- Audit logging

## Technology Stack

- **Language**: Go 1.21+
- **Testing**: Go testing framework
- **Containerization**: Docker
- **Orchestration**: Kubernetes
- **CI/CD**: GitHub Actions
- **Monitoring**: Prometheus + Grafana (planned)

## References

- Database System Concepts (Silberschatz, Korth, Sudarshan)
- ARIES Recovery Algorithm
- B+ Tree Indexing
- Two-Phase Locking Protocol

---

**Last Updated**: 2026-01-30  
**Status**: Phase 1 - Storage Layer Implementation