all:

run:
	go run server.go

fmt:
	go fmt server.go

get:
	go get -u golang.org/x/crypto/bcrypt
	go get -u github.com/gorilla/mux

.PHONY: run fmt clean
