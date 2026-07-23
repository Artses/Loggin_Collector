.PHONY: all build run clean test

BINARY_NAME=log-collector
CMD_DIR=./cmd/main.go

all: build

build:
	@echo "Building native binary for Linux/macOS..."
	go build -o $(BINARY_NAME) $(CMD_DIR)
	@echo "Build complete: ./$(BINARY_NAME)"

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "Cleanup complete."

test:
	@echo "Running tests..."
	go test ./...
