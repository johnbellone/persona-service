PROTOC_GATEWAY_OPTS=--grpc-gateway_opt=allow_repeated_fields_in_body=true,generate_unbound_methods=true,request_context=true

.PHONY: all proto check build tools

all: build

clean:
	@rm -fr tmp internal/gen bin/persona-service

tools:
	@mkdir -p internal/gen bin/
	@go mod tidy
	@go generate -tags tools tools/tools.go
	@minica -domains localhost -ip-addresses 127.0.0.1 -ca-cert server.crt -ca-key server.key
	@rm -fr localhost

proto: tools
	@buf check lint
	@protoc --experimental_allow_proto3_optional \
		-Iinternal/proto --go_out=paths=source_relative:./internal/gen \
		internal/proto/persona/type/*.proto internal/proto/persona/rpc/*.proto
	@protoc -Iinternal/proto --go_out=plugins=grpc,paths=source_relative:./internal/gen \
		internal/proto/persona/api/v1/*.proto

build:
	@go build -o bin/persona-service

check:
	@golangci-lint run --skip-dirs internal/proto/google

release: all
	@zip --junk-paths bin/persona-service README.md
