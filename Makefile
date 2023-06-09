.PHONY: build install run build-linux build-darwin build-windows build-arm

BUILD_DIR := ./build
CLI_NAME := opencamp
GOPATH := $(shell go env GOPATH)

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(CLI_NAME) .

install: build
	cp $(BUILD_DIR)/$(CLI_NAME) $(GOPATH)/bin

run: build
	$(BUILD_DIR)/$(CLI_NAME)

build-linux:
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(CLI_NAME)-linux-amd64 .

build-darwin:
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(CLI_NAME)-darwin-amd64 .

build-windows:
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(CLI_NAME)-windows-amd64.exe .

build-arm:
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(CLI_NAME)-linux-arm64 .

release: build-linux build-darwin build-windows build-arm
