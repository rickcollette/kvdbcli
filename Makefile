# Variables
APP_NAME := kvdbcli
KVDBGENKEY_PATH := tools/kvdbgenkey
BIN_DIR := bin
VERSION_FILE := VERSION
VERSION_GO := cmd/version.go

# Read the current version from the VERSION file
CURRENT_VERSION := $(shell cat $(VERSION_FILE))
VERSION_MAJOR := $(word 1, $(subst ., ,$(CURRENT_VERSION)))
VERSION_MINOR := $(word 2, $(subst ., ,$(CURRENT_VERSION)))
VERSION_PATCH := $(word 3, $(subst ., ,$(CURRENT_VERSION)))

# Default target
all: build build-genkey

# Build the main application
build: 
	@echo "Building $(APP_NAME)..."
	@go build -o $(BIN_DIR)/$(APP_NAME) main.go
	@echo "$(APP_NAME) built successfully."

# Build the kvdbgenkey tool
build-genkey:
	@echo "Building kvdbgenkey tool..."
	@go build -o $(BIN_DIR)/kvdbgenkey $(KVDBGENKEY_PATH)/main.go
	@echo "kvdbgenkey tool built successfully."

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean up binaries
clean:
	@echo "Cleaning up binaries..."
	@rm -rf $(BIN_DIR)
	@echo "Cleaned up successfully."

# Lint the code (requires golangci-lint to be installed)
lint:
	@echo "Linting code..."
	@golangci-lint run

# Push version to GitHub (for use with your pushversion.sh script)
push:
	@echo "Pushing new version to GitHub..."
	@./githubBuild/pushversion.sh

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	@./$(BIN_DIR)/$(APP_NAME)

# Help message
help:
	@echo "Makefile commands:"
	@echo "  make build            - Build the main application"
	@echo "  make build-genkey     - Build the kvdbgenkey tool"
	@echo "  make test             - Run tests"
	@echo "  make clean            - Clean up binaries"
	@echo "  make lint             - Lint the code"
	@echo "  make push             - Push new version to GitHub"
	@echo "  make deps             - Install dependencies"
	@echo "  make run              - Run the main application"
	@echo "  make increment-minor  - Increment the minor version number"
	@echo "  make increment-major  - Increment the major version number"
	@echo "  make increment-patch  - Increment the patch version number"

# Function to update cmd/version.go
update-version-go:
	@sed -i 's/const Version string = "v[0-9]\+\.[0-9]\+\.[0-9]\+"/const Version string = "v$(NEW_VERSION)"/' $(VERSION_GO)
	@echo "Updated version in $(VERSION_GO)"

# Increment version numbers
increment-patch:
	@echo "Current version: $(CURRENT_VERSION)"
	@NEW_VERSION=$(VERSION_MAJOR).$(VERSION_MINOR).$$(( $(VERSION_PATCH) + 1 )) && \
	echo $$NEW_VERSION > $(VERSION_FILE) && \
	echo "Updated version to $$NEW_VERSION" && \
	$(MAKE) update-version-go NEW_VERSION=$$NEW_VERSION

increment-minor:
	@echo "Current version: $(CURRENT_VERSION)"
	@NEW_VERSION=$(VERSION_MAJOR).$$(( $(VERSION_MINOR) + 1 )).0 && \
	echo $$NEW_VERSION > $(VERSION_FILE) && \
	echo "Updated version to $$NEW_VERSION" && \
	$(MAKE) update-version-go NEW_VERSION=$$NEW_VERSION

increment-major:
	@echo "Current version: $(CURRENT_VERSION)"
	@NEW_VERSION=$$(( $(VERSION_MAJOR) + 1 )).0.0 && \
	echo $$NEW_VERSION > $(VERSION_FILE) && \
	echo "Updated version to $$NEW_VERSION" && \
	$(MAKE) update-version-go NEW_VERSION=$$NEW_VERSION

# Run both the build and kvdbgenkey build
build-all: build build-genkey

.PHONY: all build build-genkey test clean lint push deps run help build-all increment-patch increment-minor increment-major
