#变量
APP_NAME=web3-practice
GOOS=darwin
GOARCH=arm64

.PHONY: build run test query

build:
	@echo "Building $(APP_NAME)..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./bin/$(APP_NAME)

run: build
	@./bin/$(APP_NAME)

test:
	@echo "Running tests..."
	@go test -v ./tests

query: 
	@echo "test Querying..."
	@go test -v ./tests/query_testing/

wallet: 
	@echo "test Wallet..."
	@go test -v ./tests/wallet_testing/
