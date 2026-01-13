.PHONY: help build run clean test install uninstall lint fmt

BINARY_NAME=gabriel
MAIN_FILE=main.go
BUILD_DIR=build
GO=go
GOFLAGS=-v
GOARCH=amd64

help:
	@echo "Word of God - Gabriel Compiler"
	@echo ""
	@echo "Available targets:"
	@echo "  make build              Build the Gabriel compiler"
	@echo "  make build-windows      Build for Windows (64-bit)"
	@echo "  make build-linux        Build for Linux (64-bit)"
	@echo "  make build-macos        Build for macOS (64-bit)"
	@echo "  make build-all          Build for all platforms"
	@echo "  make run FILE=prog.wog  Run a Word of God program"
	@echo "  make install            Install Gabriel to system"
	@echo "  make uninstall          Remove Gabriel from system"
	@echo "  make test               Run Go tests"
	@echo "  make lint               Run linter"
	@echo "  make fmt                Format code"
	@echo "  make clean              Remove build artifacts"
	@echo "  make setup-vscode       Install VS Code extension"
	@echo "  make help               Show this help message"

build: clean
	@echo "Building Gabriel compiler..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

build-windows:
	@echo "Building Gabriel for Windows..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=$(GOARCH) $(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME).exe $(MAIN_FILE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME).exe"

build-linux:
	@echo "Building Gabriel for Linux..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=$(GOARCH) $(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux $(MAIN_FILE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)-linux"

build-macos:
	@echo "Building Gabriel for macOS..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=$(GOARCH) $(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-macos $(MAIN_FILE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)-macos"

build-all: build-windows build-linux build-macos
	@echo "All platform builds complete"

run: build
	@if [ -z "$(FILE)" ]; then \
		echo "Error: Please specify FILE=script.wog"; \
		exit 1; \
	fi
	@./$(BUILD_DIR)/$(BINARY_NAME) $(FILE)

install: build
	@echo "Installing Gabriel to system..."
	@mkdir -p $(HOME)/.local/bin
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(HOME)/.local/bin/
	@chmod +x $(HOME)/.local/bin/$(BINARY_NAME)
	@echo "Gabriel installed to $(HOME)/.local/bin/$(BINARY_NAME)"
	@echo "Add $(HOME)/.local/bin to your PATH if not already present"

uninstall:
	@echo "Uninstalling Gabriel..."
	@rm -f $(HOME)/.local/bin/$(BINARY_NAME)
	@echo "Gabriel removed"

test:
	@echo "Running tests..."
	@$(GO) test $(GOFLAGS) ./...

lint:
	@echo "Running golangci-lint..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not found. Install: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; exit 1)
	@golangci-lint run ./...

fmt:
	@echo "Formatting code..."
	@$(GO) fmt ./...
	@goimports -w .

setup-vscode:
	@echo "Setting up VS Code extension..."
	@mkdir -p $(HOME)/.vscode/extensions/
	@cp -r .vscode/extensions/word-of-god-syntax $(HOME)/.vscode/extensions/
	@echo "VS Code extension installed"
	@echo "Restart VS Code to apply syntax highlighting"

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@$(GO) clean
	@echo "Clean complete"

.DEFAULT_GOAL := help