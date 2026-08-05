APP_NAME := go-design-patterns
GO_VERSION := 1.26

.PHONY: all build test race lint fmt vet clean tidy help

all: test

## Build the project
build:
	go build ./...

## Run unit tests
test:
	go test ./...

## Run tests with race detector
race:
	go test -race ./...

## Run gofmt check
fmt:
	@echo "Checking formatting..."
	@test -z "$$(gofmt -l .)" || (echo "Files are not formatted. Run 'make format'"; exit 1)

## Format all Go files
format:
	gofmt -w .

## Run go vet
vet:
	go vet ./...

## Run golangci-lint
lint:
	golangci-lint run

## Download and tidy dependencies
tidy:
	go mod tidy

## Clean build artifacts
clean:
	go clean
	rm -rf bin/

## Run all quality checks locally
check: fmt vet test race lint

## Display available commands
help:
	@echo "Available commands:"
	@echo ""
	@echo "  make build     - Build all packages"
	@echo "  make test      - Run unit tests"
	@echo "  make race      - Run tests with race detector"
	@echo "  make fmt       - Verify go formatting"
	@echo "  make format    - Format Go files"
	@echo "  make vet       - Run go vet"
	@echo "  make lint      - Run golangci-lint"
	@echo "  make tidy      - Run go mod tidy"
	@echo "  make clean     - Remove build artifacts"
	@echo "  make check     - Run all validation checks"