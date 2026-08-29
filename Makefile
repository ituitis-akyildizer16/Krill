.PHONY: build test lint completions clean

build:
	go build -o bin/krill ./cmd/krill

test:
	go test ./... -race -count=1

lint:
	golangci-lint run ./...

completions:
	./scripts/completions.sh

clean:
	rm -rf bin
