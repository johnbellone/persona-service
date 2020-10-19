PROTOC_OPTS=-Iinternal/proto --go_out=paths=source_relative:./proto
PROTOC_GATEWAY_OPTS=--grpc-gateway_opt=allow_repeated_fields_in_body=true,generate_unbound_methods=true,request_context=true

.PHONY: all proto chcek build tools

all: check tools build

tools:
	go generate -tags tools tools/tools.go

proto: tools
	@buf check lint
	@mkdir -p proto
	@protoc $(PROTOC_OPTS) --experimental_allow_proto3_optional \
		internal/proto/persona/type/*.proto
	@protoc $(PROTOC_OPTS) --grpc-gateway_out=proto \
		$(PROTOC_GATEWAY_OPTS) --grpc-gateway_opt=paths=source_relative  \
		internal/proto/persona/api/v1/*.proto

build:
	@go build -o bin/persona-service

check:
	@go mod tidy
	@golangci-lint run --skip-dirs proto/

release: all
	@zip --junk-paths bin/persona-service README.md
