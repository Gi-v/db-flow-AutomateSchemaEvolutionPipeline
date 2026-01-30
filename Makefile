.PHONY: build test run docker-build docker-run clean lint coverage help install-tools

# Variables
BINARY_NAME=minidb
DOCKER_IMAGE=minidb:latest
MAIN_PATH=./src/main.go
GOFLAGS=-v

# Help target
help: ## Display this help message
	@echo "Mini Relational DBMS - Makefile targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Build targets
build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	go build $(GOFLAGS) -o $(BINARY_NAME) $(MAIN_PATH)

build-race: ## Build with race detector
	@echo "Building with race detector..."
	go build $(GOFLAGS) -race -o $(BINARY_NAME) $(MAIN_PATH)

# Run targets
run: build ## Build and run the application
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

run-dev: ## Run in development mode with auto-reload (requires air)
	@which air > /dev/null || (echo "Installing air..." && go install github.com/cosmtrek/air@latest)
	air

# Test targets
test: ## Run unit tests
	@echo "Running unit tests..."
	go test -v ./src/... ./tests/unit/...

test-integration: ## Run integration tests
	@echo "Running integration tests..."
	go test -v ./tests/integration/...

test-all: ## Run all tests
	@echo "Running all tests..."
	go test -v ./...

coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Linting targets
lint: ## Run linter
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

fmt: ## Format code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

# Docker targets
docker-build: ## Build Docker image
	@echo "Building Docker image $(DOCKER_IMAGE)..."
	docker build -f deploy/docker/Dockerfile -t $(DOCKER_IMAGE) .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE)

docker-compose-up: ## Start docker-compose stack
	docker compose -f deploy/docker/docker-compose.yml up -d

docker-compose-down: ## Stop docker-compose stack
	docker compose -f deploy/docker/docker-compose.yml down

# Kubernetes targets
k8s-deploy: ## Deploy to Kubernetes
	kubectl apply -f deploy/k8s/

k8s-delete: ## Delete Kubernetes resources
	kubectl delete -f deploy/k8s/

helm-install: ## Install Helm chart
	helm install minidb deploy/helm/minidb

helm-uninstall: ## Uninstall Helm chart
	helm uninstall minidb

# Development targets
install-tools: ## Install development tools
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cosmtrek/air@latest
	go install golang.org/x/tools/cmd/goimports@latest

deps: ## Download dependencies
	go mod download
	go mod tidy

# Clean targets
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html
	rm -rf dist/
	go clean

clean-all: clean ## Clean all including Docker images
	docker rmi -f $(DOCKER_IMAGE) 2>/dev/null || true

# Database targets
db-start: ## Start local database (PostgreSQL for development)
	docker run -d --name minidb-postgres -p 5432:5432 -e POSTGRES_PASSWORD=minidb -e POSTGRES_DB=minidb postgres:16

db-stop: ## Stop local database
	docker stop minidb-postgres || true
	docker rm minidb-postgres || true

# CI targets
ci-test: lint test ## Run CI tests (lint + test)

# Default target
.DEFAULT_GOAL := help
