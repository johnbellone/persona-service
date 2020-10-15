PROTOC_OPTS=-Iproto -Iinternal/proto

.PHONY: all proto chcek build tools

all: check tools build

tools:
	go generate -tags tools tools/tools.go

proto: tools
	@buf check lint
	@protoc $(PROTOC_OPTS) --experimental_allow_proto3_optional \
		--go_out=paths=source_relative:./proto \
		proto/persona/type/*.proto
	@protoc $(PROTOC_OPTS) \
		--micro_out=paths=source_relative:./proto \
		--go_out=paths=source_relative:./proto \
		proto/persona/api/*.proto

build:
	@go build -o bin/persona-service

check:
	@go mod tidy
	@golangci-lint run --skip-dirs proto/

release: all
	@zip --junk-paths bin/persona-service README.md
