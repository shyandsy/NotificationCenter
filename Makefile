server:
	go run main.go

test:
	go test -v -cover ./...

lint:
	golangci-lint --version
	golangci-lint run --timeout 600s --verbose ./$*/...

proto:
	protoc --go_out=internal/pb internal/proto/*.proto
	protoc --go-grpc_out=internal/pb --go-grpc_opt=require_unimplemented_servers=false internal/proto/*.proto

.PHONY: server test proto
