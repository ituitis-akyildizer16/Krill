.PHONY: build test lint completions clean

build:
	go build -o bin/krill ./cmd/krill

test:
	go test ./... -race -count=1

lint:
