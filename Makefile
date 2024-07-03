TARGET = ./build/report-generator

.PHONY: build

# Build project and put the output binary in build/
build:
	@go build -o $(TARGET) ./cmd/csv_report_generator.go

# Build project under linux and put the output binary in build/
build-linux:
	@GOOS=linux GOARCH=amd64 go build -o $(TARGET)-linux-amd64 ./cmd/csv_report_generator.go

# Build and Run project
run: build
	@$(TARGET) config/config.json

# Run the tests of the project
test:
	@go test ./...

# Docker:
image: build-linux ## Use the dockerfile to build the container
	@docker build . -t report-generator

# Create and start project container
compose-up: image
	 @docker compose up --detach

# Stop and remove project container
compose-down:
	 @docker compose down