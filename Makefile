gen:
	protoc -I=proto/ --go_out=./ --go-grpc_out=./ proto/*.proto

clean:
	del pb\*.go

test:
	go test -cover ./...

server:
	go run ./cmd/server/main.go -port 8080

client:
	go run ./cmd/client/main.go -address 0.0.0.0:8080
