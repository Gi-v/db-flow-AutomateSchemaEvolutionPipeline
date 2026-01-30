# Mini Relational DBMS

A semester-long project implementing a simplified relational database management system with SQL support, disk-backed storage, B+ tree indexing, transactions, and DevOps integration.

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker (for containerized deployment)
- Make (for build automation)

### Local Development

1. **Clone the repository**
```bash
git clone https://github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline.git
cd db-flow-AutomateSchemaEvolutionPipeline
```

2. **Build the project**
```bash
make build
```

3. **Run tests**
```bash
make test
```

4. **Start the DBMS server**
```bash
make run
```

The server will start on `http://localhost:8080`

### Using Docker

```bash
# Build and run with Docker Compose
docker-compose up mini-dbms

# Or build Docker image manually
make docker-build
make docker-run
```

### Using GitHub Codespaces

This project is optimized for GitHub Codespaces with pre-configured development environment:

1. Open the repository in Codespaces
2. All tools (Go, Docker, kubectl) are pre-installed
3. Run `make build && make test` to verify setup

## 📋 Project Overview

This Mini Relational DBMS implements:

### Core Features
- **SQL Engine**: Basic SQL operations (SELECT, INSERT, UPDATE, DELETE)
- **Storage Layer**: Disk-backed storage with pages and tuples
- **Indexing**: B+ tree implementation for efficient queries
- **Transactions**: ACID guarantees with proper isolation
- **Write-Ahead Logging (WAL)**: Crash recovery mechanism
- **Concurrency Control**: Simple locking for multi-user access
- **REST API**: HTTP endpoints for database operations

### DevOps & Infrastructure
- **Containerization**: Docker and Docker Compose
- **Orchestration**: Kubernetes manifests
- **CI/CD**: Automated testing and deployment with GitHub Actions
- **Monitoring**: Prometheus and Grafana integration (planned)
- **Development**: GitHub Codespaces ready

## 🏗️ Architecture

```
┌─────────────────┐
│   REST API      │  ← HTTP endpoints
└────────┬────────┘
         │
┌────────▼────────┐
│   SQL Parser    │  ← Parse SQL statements
└────────┬────────┘
         │
┌────────▼────────┐
│ Query Executor  │  ← Execute queries
└────────┬────────┘
         │
    ┌────┴────┐
    │         │
┌───▼──┐  ┌──▼────┐
│Index │  │Storage│  ← B+ Tree & Page Manager
└───┬──┘  └──┬────┘
    │        │
┌───▼────────▼───┐
│  Transaction   │  ← ACID & WAL
└────────────────┘
```

See [ARCHITECTURE.md](docs/ARCHITECTURE.md) for detailed design.

## 📁 Project Structure

```
db-flow-AutomateSchemaEvolutionPipeline/
├── src/
│   ├── storage/          # Storage engine, page management
│   ├── parser/           # SQL parser
│   ├── executor/         # Query execution engine
│   ├── index/            # B+ tree implementation
│   ├── transaction/      # Transaction manager, WAL, locking
│   ├── main.go           # Entry point
│   └── api/              # REST API endpoints
├── tests/
│   ├── unit/             # Unit tests
│   ├── integration/      # Integration tests
│   └── acceptance/       # End-to-end scenarios
├── deploy/
│   ├── Dockerfile        # Container image
│   ├── docker-compose.yml # Local development
│   ├── k8s/              # Kubernetes manifests
│   └── helm/             # Helm charts
├── .github/workflows/    # CI/CD pipelines
├── docs/                 # Documentation
├── Makefile              # Build automation
└── README.md             # This file
```

## 🛠️ Development

### Available Make Targets

```bash
make build         # Build the DBMS binary
make test          # Run all tests
make run           # Run the DBMS server
make clean         # Clean build artifacts
make docker-build  # Build Docker image
make docker-run    # Run Docker container
make lint          # Run golangci-lint
make coverage      # Run tests with coverage report
make install-tools # Install development tools
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run specific package tests
go test ./src/storage/... -v
```

### Code Quality

```bash
# Run linter
make lint

# Format code
make fmt

# Run all checks
make check
```

## 🧪 API Endpoints

### Health Check
```bash
curl http://localhost:8080/health
```

### Query Endpoint (planned)
```bash
curl -X POST http://localhost:8080/api/query \
  -H "Content-Type: application/json" \
  -d '{"sql": "SELECT * FROM users"}'
```

## 🚦 CI/CD Pipeline

The project uses GitHub Actions for continuous integration:

- ✅ Automated testing on push/PR
- ✅ Code linting with golangci-lint
- ✅ Coverage reporting
- ✅ Docker image building
- 🔄 Kubernetes deployment (planned)

## 📚 Documentation

- [ARCHITECTURE.md](docs/ARCHITECTURE.md) - System architecture and design decisions
- [SETUP.md](docs/SETUP.md) - Detailed setup guide for developers

## 🗺️ Roadmap

### Phase 1: Storage Layer (Current)
- [x] Page management
- [x] Basic storage manager
- [ ] Tuple format
- [ ] Disk persistence

### Phase 2: SQL Parser
- [ ] Lexer/tokenizer
- [ ] AST construction
- [ ] Query validation

### Phase 3: Query Execution
- [ ] Table scan
- [ ] Index scan
- [ ] Join algorithms

### Phase 4: Indexing
- [ ] B+ tree implementation
- [ ] Index maintenance

### Phase 5: Transactions & Recovery
- [ ] Transaction manager
- [ ] Write-ahead logging
- [ ] Crash recovery

### Phase 6: Concurrency
- [ ] Lock manager
- [ ] Deadlock detection

### Phase 7: Production Ready
- [ ] Performance optimization
- [ ] Monitoring integration
- [ ] Production deployment

## 🤝 Contributing

This is a semester project. Contributions are managed through:
1. Feature branches
2. Pull requests with CI checks
3. Code review process

## 📝 License

See [LICENSE](LICENSE) file for details.

## 👥 Team

Developed as part of a semester-long database systems course.

---

**Status**: 🟢 Active Development | Phase 1 (Storage Layer)
