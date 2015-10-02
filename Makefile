all:

run:
	go run server.go

fmt:
	go fmt server.go

get:
	go get -u golang.org/x/crypto/bcrypt
	rm -rf src/golang.org/x/crypto/bcrypt/.git
	go get -u github.com/gorilla/mux
	rm -rf src/github.com/gorilla/mux/.git

build:
	go build server.go

.PHONY: run fmt clean
