all: check build

build:
	@go build -o bin/persona-service

check:
	@go mod tidy
	@golangci-lint run --skip-dirs proto/

release: all
	@zip --junk-paths bin/persona-service README.md

generate:
	@go generate
