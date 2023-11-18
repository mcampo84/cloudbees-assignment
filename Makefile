.PHONY: install-deps rebuild-proto

PROTO_FILE = ./internal/proto/tickets.proto

install-deps:
	@which protoc > /dev/null || (echo "protoc not found, installing..." && brew install protobuf) 
	@which protoc-gen-go-grpc > /dev/null || (echo "protoc-gen-go-grpc not found, installing..." && brew install protoc-gen-go-grpc)

rebuild-proto: install-deps
	protoc -I=. --go_out ./internal/proto --go-grpc_out=./internal/proto $(PROTO_FILE)
