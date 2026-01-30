.PHONY: help build test run clean docker-build docker-run lint coverage install-tools

# Default target
help:
	@echo "Mini Relational DBMS - Makefile targets:"
	@echo "  make build         - Build the DBMS binary"
	@echo "  make test          - Run all tests"
	@echo "  make run           - Run the DBMS server"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make lint          - Run golangci-lint"
	@echo "  make coverage      - Run tests with coverage"
	@echo "  make install-tools - Install development tools"

# Build the binary
build:
	@echo "Building Mini DBMS..."
	@mkdir -p bin
	go build -o bin/dbms ./src/main.go

# Run all tests
test:
	@echo "Running tests..."
	go test ./... -v

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	@mkdir -p coverage
	go test ./... -coverprofile=coverage/coverage.out
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Coverage report: coverage/coverage.html"

# Run the DBMS server
run: build
	@echo "Starting Mini DBMS server..."
	./bin/dbms

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -rf coverage/
	go clean

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t mini-dbms:latest -f deploy/Dockerfile .

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 mini-dbms:latest

# Run linter
lint:
	@echo "Running golangci-lint..."
	golangci-lint run ./...

# Install development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Tidy dependencies
tidy:
	@echo "Tidying Go modules..."
	go mod tidy

# Run all checks (test + lint)
check: test lint
	@echo "All checks passed!"
