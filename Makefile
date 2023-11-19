.PHONY: install-deps generate-mocks proto test

PROTO_FILE = ./internal/proto/tickets.proto

install-deps:
	@which protoc > /dev/null || (echo "protoc not found, installing..." && brew install protobuf) 
	@which protoc-gen-go-grpc > /dev/null || (echo "protoc-gen-go-grpc not found, installing..." && brew install protoc-gen-go-grpc)
	@which mockgen > /dev/null || (echo "mockgen not found, installing..." && go install github.com/golang/mock/mockgen@v1.6.0)

generate-mocks: install-deps
	go generate ./...

proto: install-deps
	protoc -I=. --go_out ./internal/proto --go-grpc_out=./internal/proto $(PROTO_FILE)

test:
	go test ./...
