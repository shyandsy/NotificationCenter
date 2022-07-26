server:
	go run main.go

test:
	go test -v -cover ./...

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
            --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
            proto/*.proto
.PHONY: server test proto