# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies (if any)
RUN go mod download || true

# Copy source code
COPY src/ ./src/

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dbms ./src/main.go

# Runtime stage - use scratch for minimal image
FROM alpine:3.18

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/dbms .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./dbms"]
