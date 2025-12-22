.DEFAULT_GOAL := build

# Install tools used in the Makefile
tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

# Format the code
fmt:
	go fmt ./...
.PHONY: fmt

# Run linters
lint: fmt
	~/go/bin/golangci-lint run
.PHONY: lint

# Run static analysis
vet: fmt
	go vet ./...
	~/go/bin/staticcheck ./...
.PHONY: vet

# Build the application
build: vet
	go build -o json2bin-parser ./cmd/json2bin-parser/
.PHONY: build

# Run tests
test:
	go test -v -race ./...
.PHONY: test

# Clean up build artifacts
clean:
	rm -f json2bin-parser
.PHONY: clean
