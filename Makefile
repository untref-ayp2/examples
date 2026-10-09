.PHONY: build test test-v fmt lint clean

build:
	go build ./...

test:
	go test ./...

test-v:
	go test -v ./...

fmt:
	gofmt -w .

lint:
	golangci-lint run

clean:
	go clean -testcache
