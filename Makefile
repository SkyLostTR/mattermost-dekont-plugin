# Variables
PLUGIN_ID = mattermost-dekont-plugin
PLUGIN_VERSION = 1.0.0
BUNDLE_NAME = $(PLUGIN_ID)-$(PLUGIN_VERSION)-linux.tar.gz
GOOS = linux
GOARCH = amd64

# Default target
.PHONY: all
all: build

# Build the plugin for Linux
.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o plugin

# Clean build artifacts
.PHONY: clean
clean:
	rm -f plugin
	rm -f *.tar.gz
	rm -rf dist/

# Create plugin bundle for deployment
.PHONY: bundle
bundle: build
	mkdir -p dist/server
	cp plugin dist/server/plugin-linux-amd64
	chmod +x dist/server/plugin-linux-amd64
	cp plugin.json dist/
	cd dist && tar -czf ../$(BUNDLE_NAME) *

# Install dependencies
.PHONY: deps
deps:
	go mod tidy
	go mod download

# Run tests (when added)
.PHONY: test
test:
	go test ./...

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
.PHONY: lint
lint:
	golangci-lint run

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  build  - Build the plugin executable"
	@echo "  bundle - Create a plugin bundle for deployment"
	@echo "  clean  - Clean build artifacts"
	@echo "  deps   - Install dependencies"
	@echo "  test   - Run tests"
	@echo "  fmt    - Format code"
	@echo "  lint   - Lint code"
	@echo "  help   - Show this help message"
