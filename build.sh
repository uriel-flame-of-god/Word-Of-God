#!/bin/bash
# build.sh - Build script for Word of God Gabriel Compiler (Linux/macOS)

set -e

BINARY_NAME="gabriel"
MAIN_FILE="main.go"
BUILD_DIR="build"
VERSION="1.0.0"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

print_help() {
    echo "Word of God - Gabriel Compiler Build Script"
    echo ""
    echo "Usage: ./build.sh [OPTION]"
    echo ""
    echo "Options:"
    echo "  linux       Build for Linux (64-bit)"
    echo "  macos       Build for macOS (64-bit)"
    echo "  windows     Build for Windows (64-bit)"
    echo "  all         Build for all platforms"
    echo "  clean       Remove build artifacts"
    echo "  install     Build and install to system"
    echo "  help        Show this help message"
    echo ""
    echo "Examples:"
    echo "  ./build.sh              # Build for current OS"
    echo "  ./build.sh all          # Build for all platforms"
    echo "  ./build.sh install      # Install to system"
}

build_linux() {
    echo -e "${YELLOW}Building Gabriel for Linux...${NC}"
    mkdir -p "$BUILD_DIR"
    GOOS=linux GOARCH=amd64 go build -v -o "$BUILD_DIR/${BINARY_NAME}-linux" "$MAIN_FILE"
    echo -e "${GREEN}Build complete: $BUILD_DIR/${BINARY_NAME}-linux${NC}"
}

build_macos() {
    echo -e "${YELLOW}Building Gabriel for macOS...${NC}"
    mkdir -p "$BUILD_DIR"
    GOOS=darwin GOARCH=amd64 go build -v -o "$BUILD_DIR/${BINARY_NAME}-macos" "$MAIN_FILE"
    echo -e "${GREEN}Build complete: $BUILD_DIR/${BINARY_NAME}-macos${NC}"
}

build_windows() {
    echo -e "${YELLOW}Building Gabriel for Windows...${NC}"
    mkdir -p "$BUILD_DIR"
    GOOS=windows GOARCH=amd64 go build -v -o "$BUILD_DIR/${BINARY_NAME}.exe" "$MAIN_FILE"
    echo -e "${GREEN}Build complete: $BUILD_DIR/${BINARY_NAME}.exe${NC}"
}

build_current() {
    echo -e "${YELLOW}Building Gabriel for current OS...${NC}"
    mkdir -p "$BUILD_DIR"
    go build -v -o "$BUILD_DIR/$BINARY_NAME" "$MAIN_FILE"
    chmod +x "$BUILD_DIR/$BINARY_NAME"
    echo -e "${GREEN}Build complete: $BUILD_DIR/$BINARY_NAME${NC}"
}

build_all() {
    echo -e "${YELLOW}Building Gabriel for all platforms...${NC}"
    build_linux
    build_macos
    build_windows
    echo -e "${GREEN}All platform builds complete${NC}"
}

clean() {
    echo -e "${YELLOW}Cleaning build artifacts...${NC}"
    rm -rf "$BUILD_DIR"
    go clean
    echo -e "${GREEN}Clean complete${NC}"
}

install() {
    build_current
    echo -e "${YELLOW}Installing Gabriel to system...${NC}"
    mkdir -p "$HOME/.local/bin"
    cp "$BUILD_DIR/$BINARY_NAME" "$HOME/.local/bin/"
    chmod +x "$HOME/.local/bin/$BINARY_NAME"
    echo -e "${GREEN}Gabriel installed to $HOME/.local/bin/$BINARY_NAME${NC}"
    echo -e "${YELLOW}Add $HOME/.local/bin to your PATH if not already present${NC}"
}

# Main script logic
case "${1:-}" in
    linux)
        build_linux
        ;;
    macos)
        build_macos
        ;;
    windows)
        build_windows
        ;;
    all)
        build_all
        ;;
    clean)
        clean
        ;;
    install)
        install
        ;;
    help|--help|-h)
        print_help
        ;;
    *)
        build_current
        ;;
esac

echo ""
echo -e "${GREEN}Done!${NC}"