proto:
	protoc pkg/pg/*.proto --go_out=. --go-grpc-out=.
server:
	go run cmd/main.go