# Project Phases - Week-by-Week Guide

This document outlines the semester-long development plan for the Mini Relational DBMS project, broken down into manageable weekly phases.

## Overview

**Duration**: 14 weeks  
**Approach**: Incremental development with testing at each phase  
**DevOps**: CI/CD runs on every commit from day one

---

## Phase 0: Project Setup (Week 1)

**Goal**: Establish project infrastructure

### Tasks
- [x] Set up Go module and dependencies
- [x] Create directory structure
- [x] Configure Makefile with build/test/run targets
- [x] Set up Docker and docker-compose
- [x] Create Kubernetes manifests
- [x] Configure CI/CD pipelines (GitHub Actions)
- [x] Set up monitoring (Prometheus + Grafana)
- [x] Create dev container for Codespaces
- [x] Write initial documentation

### Deliverables
- Working build system
- CI pipeline passing
- Docker image builds successfully
- Documentation ready

### Success Criteria
- `make build` succeeds
- `make test` runs stub tests
- Docker container starts and responds to health check
- Codespaces environment works

---

## Phase 1: Storage Layer - Pages (Weeks 2-3)

**Goal**: Implement disk-backed page storage

### Week 2: Page Abstraction

**Tasks**:
1. Implement Page struct with 8KB fixed size
2. Add Pin/Unpin reference counting
3. Implement dirty flag tracking
4. Add Read/Write methods for page data
5. Write comprehensive unit tests

**Files**:
- `src/storage/page.go` (enhance stub)
- `tests/unit/page_test.go`

**Tests**:
```go
TestPagePinUnpin()
TestPageDirtyFlag()
TestPageReadWrite()
TestPageBoundaryConditions()
```

### Week 3: Buffer Pool Manager

**Tasks**:
1. Implement buffer pool with configurable size
2. Add LRU eviction policy
3. Implement page fetching and eviction
4. Add flush operations (single page, all pages)
5. Thread-safety with mutexes
6. Write unit and integration tests

**Files**:
- `src/storage/buffer_pool.go` (enhance stub)
- `tests/unit/buffer_pool_test.go`
- `tests/integration/storage_integration_test.go`

**Tests**:
```go
TestBufferPoolFetch()
TestBufferPoolEviction()
TestBufferPoolConcurrency()
TestBufferPoolFlush()
```

### Deliverables
- Fully functional page abstraction
- Working buffer pool with LRU
- 90%+ test coverage
- CI passing

---

## Phase 2: Storage Layer - Tables (Week 4)

**Goal**: Implement table management

### Tasks
1. Enhance Schema and Column definitions
2. Implement TableManager for CRUD on tables
3. Add tuple serialization/deserialization
4. Implement table scanning
5. Add integration with buffer pool
6. Write comprehensive tests

**Files**:
- `src/storage/table.go` (enhance stub)
- `src/storage/tuple.go` (new)
- `src/storage/schema.go` (new)
- `tests/unit/table_test.go`

**Tests**:
```go
TestTableCreation()
TestTableInsert()
TestTableScan()
TestTableDrop()
TestSchemaValidation()
```

### Deliverables
- Table creation and management
- Tuple insertion and scanning
- Schema validation
- Integration with buffer pool

---

## Phase 3: SQL Parser - Lexer (Week 5)

**Goal**: Implement lexical analysis

### Tasks
1. Enhance lexer for all SQL keywords
2. Add support for operators and literals
3. Implement string and number parsing
4. Add error reporting with line numbers
5. Write comprehensive tokenization tests

**Files**:
- `src/parser/lexer.go` (enhance stub)
- `tests/unit/lexer_test.go`

**Tests**:
```go
TestLexerKeywords()
TestLexerIdentifiers()
TestLexerNumbers()
TestLexerStrings()
TestLexerOperators()
TestLexerErrors()
```

### Deliverables
- Complete lexical analyzer
- Support for all SQL tokens
- Error reporting
- 95%+ test coverage

---

## Phase 4: SQL Parser - Parser & AST (Week 6)

**Goal**: Implement syntax analysis and AST generation

### Tasks
1. Enhance parser for all SQL statements
2. Build proper AST for each statement type
3. Add WHERE clause parsing
4. Implement error recovery
5. Write parser tests for all statement types

**Files**:
- `src/parser/parser.go` (enhance stub)
- `src/parser/ast.go` (enhance stub)
- `tests/unit/parser_test.go`

**Tests**:
```go
TestParseSelect()
TestParseInsert()
TestParseUpdate()
TestParseDelete()
TestParseCreateTable()
TestParseWhereClause()
TestParserErrors()
```

### Deliverables
- Complete SQL parser
- AST for all supported statements
- Error messages with context
- End-to-end parsing tests

---

## Phase 5: B+ Tree Indexing (Weeks 7-8)

**Goal**: Implement B+ tree for fast data access

### Week 7: B+ Tree Structure

**Tasks**:
1. Implement B+ tree node structure
2. Add insert operation with node splitting
3. Implement search operation
4. Add range scan functionality
5. Write unit tests

**Files**:
- `src/index/btree.go` (enhance stub)
- `src/index/btree_node.go` (new)
- `tests/unit/btree_test.go`

**Tests**:
```go
TestBTreeInsert()
TestBTreeSearch()
TestBTreeSplit()
TestBTreeRangeScan()
```

### Week 8: Index Integration

**Tasks**:
1. Integrate B+ tree with table manager
2. Implement index creation on table columns
3. Add index usage in query execution
4. Implement delete operation
5. Write integration tests

**Files**:
- `src/index/index_manager.go`
- `tests/integration/index_test.go`

**Tests**:
```go
TestIndexCreation()
TestIndexedQuery()
TestIndexMaintenance()
TestIndexPerformance()
```

### Deliverables
- Working B+ tree implementation
- Index integration with storage
- Performance improvement verification

---

## Phase 6: Query Executor (Weeks 9-10)

**Goal**: Execute parsed queries

### Week 9: Basic Execution

**Tasks**:
1. Enhance executor for SELECT queries
2. Implement INSERT execution
3. Add UPDATE execution
4. Implement DELETE execution
5. Write execution tests

**Files**:
- `src/executor/executor.go` (enhance stub)
- `src/executor/select.go` (new)
- `src/executor/insert.go` (new)
- `tests/unit/executor_test.go`

**Tests**:
```go
TestExecuteSelect()
TestExecuteInsert()
TestExecuteUpdate()
TestExecuteDelete()
```

### Week 10: Advanced Execution

**Tasks**:
1. Implement WHERE clause evaluation
2. Add projection (column selection)
3. Optimize table scans with indexes
4. Add query result formatting
5. Write end-to-end tests

**Files**:
- `src/executor/filter.go` (new)
- `src/executor/projection.go` (new)
- `tests/integration/query_test.go`

**Tests**:
```go
TestWhereClauseEvaluation()
TestProjection()
TestIndexOptimization()
TestEndToEndQueries()
```

### Deliverables
- Complete query executor
- All SQL operations working
- End-to-end query execution
- Performance benchmarks

---

## Phase 7: Transactions & WAL (Weeks 11-12)

**Goal**: Implement durability and recovery

### Week 11: Write-Ahead Logging

**Tasks**:
1. Enhance WAL manager
2. Implement log record writing
3. Add log flushing on commit
4. Implement checkpoint mechanism
5. Write WAL tests

**Files**:
- `src/transaction/wal.go` (enhance stub)
- `src/transaction/log_record.go` (new)
- `tests/unit/wal_test.go`

**Tests**:
```go
TestWALWrite()
TestWALFlush()
TestCheckpoint()
TestLogRecords()
```

### Week 12: Recovery

**Tasks**:
1. Implement ARIES recovery algorithm
2. Add analysis phase
3. Implement redo phase
4. Add undo phase
5. Test crash recovery scenarios

**Files**:
- `src/transaction/recovery.go` (new)
- `tests/integration/recovery_test.go`

**Tests**:
```go
TestRecoveryAfterCrash()
TestRedoPhase()
TestUndoPhase()
TestTransactionRollback()
```

### Deliverables
- Working WAL system
- Crash recovery
- Transaction durability
- Recovery tests passing

---

## Phase 8: Concurrency Control (Week 13)

**Goal**: Support concurrent transactions

### Tasks
1. Enhance lock manager
2. Implement two-phase locking (2PL)
3. Add deadlock detection
4. Implement lock upgrading
5. Write concurrency tests

**Files**:
- `src/transaction/lock_manager.go` (enhance stub)
- `src/transaction/deadlock.go` (new)
- `tests/integration/concurrency_test.go`

**Tests**:
```go
TestConcurrentReads()
TestConcurrentWrites()
TestDeadlockDetection()
TestLockUpgrade()
TestSerializability()
```

### Deliverables
- Working concurrency control
- Deadlock detection and prevention
- Serializable transactions
- Concurrency benchmarks

---

## Phase 9: Optimization & Polish (Week 14)

**Goal**: Performance tuning and final testing

### Tasks
1. Performance profiling and optimization
2. Add query optimizer (basic cost model)
3. Improve error messages
4. Complete documentation
5. Final integration testing
6. Load testing

**Files**:
- `src/optimizer/optimizer.go` (new)
- `docs/PERFORMANCE.md` (new)
- `tests/load/load_test.go` (new)

**Deliverables**:
- Performance benchmarks
- Optimization report
- Complete documentation
- Production-ready system

---

## Continuous Activities (All Weeks)

### DevOps
- Push code daily
- CI/CD validates every commit
- Monitor test coverage
- Update documentation

### Testing
- Write tests before/with code
- Maintain >85% coverage
- Fix failing tests immediately
- Add integration tests

### Documentation
- Update README as features complete
- Document API changes
- Add code comments
- Create examples

---

## Success Metrics

### Code Quality
- Test coverage > 85%
- All linters passing
- No critical security issues
- Code reviewed

### Functionality
- All SQL operations work
- Transactions are ACID
- Recovery works correctly
- Performance acceptable

### DevOps
- CI/CD pipeline green
- Docker builds successfully
- Kubernetes deployment works
- Monitoring operational

---

## Milestones

### Milestone 1 (Week 4)
✅ Storage layer complete  
✅ Tables and tuples working  
✅ Buffer pool functional

### Milestone 2 (Week 6)
✅ SQL parser complete  
✅ All statements supported  
✅ AST generation working

### Milestone 3 (Week 8)
✅ B+ tree indexing  
✅ Index integration  
✅ Query optimization

### Milestone 4 (Week 10)
✅ Query executor complete  
✅ All SQL ops working  
✅ End-to-end tests passing

### Milestone 5 (Week 12)
✅ Transactions working  
✅ WAL and recovery  
✅ Durability guaranteed

### Final Milestone (Week 14)
✅ Production ready  
✅ All features complete  
✅ Documentation done  
✅ Performance validated

---

## Tips for Success

1. **Start early**: Don't wait until the deadline
2. **Test continuously**: Write tests as you code
3. **Commit often**: Push to GitHub daily
4. **Ask questions**: Use issues for clarification
5. **Read documentation**: Study DBMS papers and books
6. **Profile performance**: Use Go's profiler
7. **Review code**: Learn from feedback
8. **Stay organized**: Follow the phase plan

---

## Resources

### Books
- "Database Internals" by Alex Petrov
- "Database System Concepts" by Silberschatz

### Courses
- CMU 15-445: Database Systems
- Stanford CS346: Database System Implementation

### Papers
- ARIES Recovery Algorithm
- B+ Tree Implementation
- Two-Phase Locking

### Tools
- Go pprof for profiling
- golangci-lint for code quality
- Prometheus for monitoring
