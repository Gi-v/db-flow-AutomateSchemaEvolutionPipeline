# Mini Relational DBMS

A simplified relational database management system built from scratch, integrating core DBMS concepts with modern DevOps practices.

![Project Status](https://img.shields.io/badge/status-active-success.svg)
![Go Version](https://img.shields.io/badge/go-1.21-blue.svg)
![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)

## 🎯 Project Overview

This project is a semester-long educational implementation of a Mini Relational DBMS that combines:

- **Core DBMS Functionality**: SQL query processing, disk-backed storage, B+ tree indexing, transactions with WAL, and concurrency control
- **DevOps Integration**: Docker containerization, Kubernetes deployment, CI/CD pipeline, and monitoring with Prometheus/Grafana

## ✨ Features

### Database Engine
- ✅ **SQL Support**: SELECT, INSERT, UPDATE, DELETE operations
- ✅ **Storage Engine**: Disk-backed pages with buffer pool management
- ✅ **Indexing**: B+ tree implementation for fast data access
- ✅ **Transactions**: Write-Ahead Logging (WAL) for durability
- ✅ **Concurrency**: Lock manager for multi-user access

### DevOps Stack
- 🐳 **Docker**: Multi-stage builds for optimized images
- ☸️ **Kubernetes**: Production-ready deployment manifests
- 🔄 **CI/CD**: Automated testing and deployment pipelines
- 📊 **Monitoring**: Prometheus metrics and Grafana dashboards
- 🛠️ **Dev Environment**: Codespaces/Dev Container support

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or higher
- Docker and Docker Compose
- Make

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline.git
   cd db-flow-AutomateSchemaEvolutionPipeline
   ```

2. **Install dependencies**
   ```bash
   make deps
   make install-tools
   ```

3. **Build and run**
   ```bash
   make build
   make run
   ```

4. **Test the API**
   ```bash
   curl http://localhost:8080/health
   curl http://localhost:8080/api/info
   ```

### Using Docker

1. **Build Docker image**
   ```bash
   make docker-build
   ```

2. **Run with Docker Compose (includes Prometheus & Grafana)**
   ```bash
   make docker-compose-up
   ```

3. **Access services**
   - MiniDB API: http://localhost:8080
   - Prometheus: http://localhost:9090
   - Grafana: http://localhost:3000 (admin/admin)

### Using Kubernetes

1. **Deploy to cluster**
   ```bash
   make k8s-deploy
   ```

2. **Or use Helm**
   ```bash
   make helm-install
   ```

## 📖 API Documentation

### Query Endpoint
```bash
POST /query
Content-Type: application/json

{
  "sql": "SELECT * FROM users WHERE id = 1"
}
```

### Health Check
```bash
GET /health
```

### Metrics (Prometheus)
```bash
GET /metrics
```

See [docs/API.md](docs/API.md) for complete API documentation.

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────┐
│                     REST API Layer                       │
│                    (Gin Framework)                       │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                SQL Parser & Executor                     │
│         (Lexer → AST → Query Execution)                 │
└─────┬──────────────┬──────────────┬────────────┬────────┘
      │              │              │            │
┌─────▼──────┐ ┌────▼────────┐ ┌──▼──────┐ ┌───▼────────┐
│  Storage   │ │   Index     │ │   WAL   │ │   Lock     │
│  Manager   │ │  (B+Tree)   │ │ Manager │ │  Manager   │
└─────┬──────┘ └─────────────┘ └─────────┘ └────────────┘
      │
┌─────▼──────────────────────────────────────────────────┐
│              Buffer Pool Manager                        │
│           (LRU Eviction Policy)                        │
└─────┬──────────────────────────────────────────────────┘
      │
┌─────▼──────────────────────────────────────────────────┐
│                  Disk Storage                           │
│              (8KB Page-based)                          │
└────────────────────────────────────────────────────────┘
```

See [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) for detailed design documentation.

## 🧪 Testing

```bash
# Run all tests
make test-all

# Unit tests only
make test

# Integration tests only
make test-integration

# Generate coverage report
make coverage
```

## 📦 Project Structure

```
├── src/                   # Source code
│   ├── storage/          # Storage engine (pages, buffer pool, tables)
│   ├── parser/           # SQL parser (lexer, AST, parser)
│   ├── executor/         # Query executor
│   ├── index/            # B+ tree indexing
│   ├── transaction/      # WAL and lock manager
│   └── main.go          # REST API entry point
├── tests/
│   ├── unit/            # Unit tests
│   └── integration/     # Integration tests
├── deploy/
│   ├── docker/          # Docker configs
│   ├── k8s/             # Kubernetes manifests
│   ├── helm/            # Helm chart
│   └── monitoring/      # Prometheus & Grafana configs
├── docs/                # Documentation
├── .github/workflows/   # CI/CD pipelines
└── .devcontainer/       # Codespaces config
```

## 🗓️ Development Phases

This is a semester-long project divided into phases:

1. **Weeks 2-4**: Storage Layer (pages, buffer pool, tables)
2. **Weeks 5-6**: SQL Parser (lexer, AST, parser)
3. **Weeks 7-8**: B+ Tree Indexing
4. **Weeks 9-10**: Query Executor
5. **Weeks 11-12**: Transactions & WAL
6. **Weeks 13-14**: Performance optimization and final testing

See [docs/PHASES.md](docs/PHASES.md) for detailed phase breakdown.

## 🔧 Make Targets

```bash
make help              # Show all available targets
make build             # Build the binary
make run               # Build and run
make test              # Run tests
make lint              # Run linter
make docker-build      # Build Docker image
make k8s-deploy        # Deploy to Kubernetes
make clean             # Clean build artifacts
```

## 📊 Monitoring

Access Grafana dashboard at http://localhost:3000 to view:
- Request rates and latencies
- Query execution times
- Buffer pool hit rates
- Transaction throughput

## 🤝 Contributing

This is an educational project. Contributions and improvements are welcome!

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by database systems courses and open-source DBMS projects
- Built with modern Go ecosystem tools and best practices
- DevOps integration following cloud-native principles

## 📚 Resources

- [PostgreSQL Internals](https://www.postgresql.org/docs/)
- [Database Internals by Alex Petrov](https://www.databass.dev/)
- [CMU Database Systems Course](https://15445.courses.cs.cmu.edu/)

---

**Note**: This is an educational project for learning DBMS concepts and DevOps practices. It is not intended for production use.
