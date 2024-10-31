# Variables
APP_NAME = task-tracker-cli
SRC = ./cmd/main.go

all: build run test

# Build the binary
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) $(SRC)

# Run the project directly with go run
run:
	@echo "Running $(APP_NAME)..."
	@go run $(SRC) $(ARGS)

# Test the project
test:
	@echo "Testing $(APP_NAME)..."
	@go test ./...

# Ignore add and other task-related commands as targets
%:
	@:

# Phony targets (so that they are always run even if a file with that name exists)
.PHONY: all build run test
