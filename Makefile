BINARY_NAME=fs

build:
	@go build -o bin/$(BINARY_NAME)

run: build
	@./bin/$(BINARY_NAME)

test:
	@go test -v ./...
