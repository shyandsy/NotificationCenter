server:
	go run main.go

test:
	go test -v -cover ./...

proto:
	protoc --go_out=pb proto/*.proto
	protoc --go-grpc_out=pb --go-grpc_opt=require_unimplemented_servers=false proto/*.proto

.PHONY: server test proto
