.PHONY: build test lint clean

BINARY_NAME=md-to-mediawiki-plus

build:
	go build -o $(BINARY_NAME)

test:
	go test ./...

lint:
	golangci-lint run

clean:
	go clean
	rm -f $(BINARY_NAME)
