FROM mcr.microsoft.com/devcontainers/go:1.21

# Install additional tools
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    postgresql-client \
    vim \
    curl \
    wget \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*

# Install Go tools
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install github.com/cosmtrek/air@latest \
    && go install golang.org/x/tools/cmd/goimports@latest

USER vscode
