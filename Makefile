all:

fmt:
	find src/ -name '*.go' -exec go fmt {} ';'

build: fmt
	gb build

serve: build
	PORT=8791 ./bin/server

.PHONY: run fmt clean
