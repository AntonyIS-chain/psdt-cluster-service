# Project variables
APP_NAME := psdt-cluster-service
MAIN_FILE := main.go

.PHONY: all build run test lint fmt clean

# Run the application
run:
	go run .

# Build the binary
build:
	go build -o bin/$(APP_NAME) $(MAIN_FILE)

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run ./...

# Clean up
clean:
	rm -rf bin/
