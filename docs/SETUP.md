# Development Setup Guide

This guide will help you set up your development environment for the Mini Relational DBMS project.

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [GitHub Codespaces Setup](#github-codespaces-setup)
3. [Local Development Setup](#local-development-setup)
4. [Development Workflow](#development-workflow)
5. [Testing](#testing)
6. [Debugging](#debugging)
7. [Troubleshooting](#troubleshooting)

## Prerequisites

### Required Tools
- **Go**: Version 1.21 or higher
  - Download from [golang.org](https://golang.org/dl/)
  - Verify: `go version`
- **Git**: For version control
  - Verify: `git --version`
- **Make**: Build automation
  - Linux/macOS: Usually pre-installed
  - Windows: Install via Chocolatey or WSL

### Optional Tools
- **Docker**: For containerized development
  - Download from [docker.com](https://www.docker.com/get-started)
  - Verify: `docker --version`
- **kubectl**: For Kubernetes deployment
  - Install via `brew install kubectl` (macOS) or package manager
- **golangci-lint**: For code linting
  - Install: `make install-tools`

## GitHub Codespaces Setup

**Recommended for quickest setup!**

GitHub Codespaces provides a pre-configured cloud development environment.

### Steps:

1. **Open in Codespaces**
   - Navigate to the repository on GitHub
   - Click "Code" → "Codespaces" → "Create codespace on main"
   - Wait for the environment to initialize (~2-3 minutes)

2. **Verify Setup**
   ```bash
   # Check Go installation
   go version
   
   # Check Docker
   docker --version
   
   # Check kubectl
   kubectl version --client
   ```

3. **Build and Test**
   ```bash
   # Download dependencies
   go mod download
   
   # Build the project
   make build
   
   # Run tests
   make test
   ```

4. **Start Development**
   - All tools are pre-installed via `.devcontainer/devcontainer.json`
   - Port 8080 is automatically forwarded
   - VS Code extensions for Go are pre-configured

### Codespaces Benefits
- ✅ Zero local setup required
- ✅ Consistent environment across team
- ✅ Pre-installed tools (Go, Docker, kubectl)
- ✅ VS Code integration with Go extensions
- ✅ Automatic port forwarding

## Local Development Setup

### Step 1: Clone Repository

```bash
git clone https://github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline.git
cd db-flow-AutomateSchemaEvolutionPipeline
```

### Step 2: Install Go

**macOS (Homebrew)**:
```bash
brew install go@1.21
```

**Linux (Ubuntu/Debian)**:
```bash
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

**Windows**:
- Download installer from [golang.org](https://golang.org/dl/)
- Add Go to PATH

### Step 3: Install Development Tools

```bash
# Install golangci-lint and other tools
make install-tools

# Or manually:
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Step 4: Configure Go Environment

Add to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.):

```bash
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Reload your shell:
```bash
source ~/.bashrc  # or ~/.zshrc
```

### Step 5: Verify Installation

```bash
# Check Go
go version

# Check dependencies
go mod download

# Build project
make build

# Run tests
make test
```

## Development Workflow

### 1. Start Development Server

```bash
# Build and run
make run

# Or run directly
go run src/main.go
```

The server will start on `http://localhost:8080`

### 2. Test Your Changes

```bash
# Run all tests
make test

# Run specific package tests
go test ./src/storage/... -v

# Run with coverage
make coverage
```

### 3. Check Code Quality

```bash
# Format code
make fmt

# Run linter
make lint

# Run all checks (test + lint)
make check
```

### 4. Build for Production

```bash
# Build optimized binary
make build

# The binary will be in bin/dbms
./bin/dbms
```

### 5. Docker Development

```bash
# Build Docker image
make docker-build

# Run container
make docker-run

# Or use docker-compose
docker-compose up mini-dbms
```

## Testing

### Unit Tests

Located in the same package as the code (e.g., `src/storage/manager_test.go`)

```bash
# Run all unit tests
go test ./... -v

# Run specific test
go test ./src/storage -run TestStorageManager -v

# Run with coverage
go test ./... -cover
```

### Integration Tests

Located in `tests/integration/`

```bash
# Run integration tests (when implemented)
go test ./tests/integration/... -v
```

### Acceptance Tests

Located in `tests/acceptance/`

```bash
# Run end-to-end tests (when implemented)
go test ./tests/acceptance/... -v
```

### Test Coverage

```bash
# Generate coverage report
make coverage

# View HTML report
open coverage/coverage.html
```

## Debugging

### VS Code Debugging

1. Install Go extension for VS Code
2. Use provided launch configuration (`.vscode/launch.json`)
3. Set breakpoints in code
4. Press F5 to start debugging

### Command-Line Debugging with Delve

```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug main program
dlv debug ./src/main.go

# Debug tests
dlv test ./src/storage
```

### Logging

Add debug logging to your code:

```go
import "log"

log.Printf("Debug: variable value = %v", myVar)
```

## IDE Setup

### VS Code (Recommended)

**Extensions**:
- Go (golang.Go)
- Docker (ms-azuretools.vscode-docker)
- YAML (redhat.vscode-yaml)
- GitLens (eamodio.gitlens)

**Settings** (`.vscode/settings.json`):
```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "package",
  "editor.formatOnSave": true
}
```

### GoLand / IntelliJ IDEA

1. Open project directory
2. Configure Go SDK (File → Settings → Go → GOROOT)
3. Enable Go modules (automatically detected)
4. Use built-in test runner and debugger

## Troubleshooting

### Go Module Issues

**Problem**: `go: cannot find main module`
```bash
# Solution: Initialize module
go mod init github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline
```

**Problem**: Dependency conflicts
```bash
# Solution: Clean and re-download
go clean -modcache
go mod download
```

### Build Issues

**Problem**: `package not found`
```bash
# Solution: Ensure module path is correct
go mod tidy
```

**Problem**: Import cycle
```bash
# Solution: Refactor code to break circular dependency
# Check import statements in affected packages
```

### Test Failures

**Problem**: Tests fail unexpectedly
```bash
# Solution: Clean build cache and re-run
go clean -testcache
go test ./... -v
```

### Docker Issues

**Problem**: Port 8080 already in use
```bash
# Solution: Find and kill process using port
# Linux/macOS:
lsof -ti:8080 | xargs kill -9

# Or change port in docker-compose.yml
```

**Problem**: Docker build fails
```bash
# Solution: Clean Docker cache
docker system prune -a
make docker-build
```

### Performance Issues

**Problem**: Slow tests
```bash
# Solution: Run tests in parallel
go test ./... -parallel 4

# Or skip long-running tests
go test ./... -short
```

## Environment Variables

```bash
# Set Go module proxy (optional)
export GOPROXY=https://proxy.golang.org,direct

# Enable Go module mode
export GO111MODULE=on

# Set log level
export LOG_LEVEL=debug
```

## Useful Commands

```bash
# Check all dependencies
go list -m all

# Update all dependencies
go get -u ./...

# Find unused dependencies
go mod tidy

# View Go environment
go env

# Clean all build artifacts
make clean

# Format all Go files
gofmt -w .

# Check for race conditions
go test -race ./...
```

## Contributing

1. Create a feature branch
   ```bash
   git checkout -b feature/my-feature
   ```

2. Make changes and test
   ```bash
   make test
   make lint
   ```

3. Commit with descriptive message
   ```bash
   git commit -m "Add feature X"
   ```

4. Push and create PR
   ```bash
   git push origin feature/my-feature
   ```

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Testing Guide](https://golang.org/pkg/testing/)
- [Docker Documentation](https://docs.docker.com/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)

## Getting Help

- Check existing [Issues](https://github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/issues)
- Review [ARCHITECTURE.md](ARCHITECTURE.md) for design details
- Ask in team discussion forum

---

**Happy Coding! 🚀**