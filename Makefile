all:

run:
	go run server.go

fmt:
	go fmt server.go

get:
	go get -u golang.org/x/crypto/bcrypt

.PHONY: run fmt clean
