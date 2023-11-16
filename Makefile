.PHONY: all deps test lint fmt

all: deps lint fmt test
ci: lint fmt test

deps:
	@echo "Installing dependencies..."
	go mod tidy

test:
	@echo "Running tests..."
	go test ./...

lint:
	@echo "Running linter..."
	golint ./...

fmt:
	@echo "Running go fmt..."
	go fmt ./...
