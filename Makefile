all:

run:
	go run server.go

fmt:
	go fmt server.go

.PHONY: run fmt clean
