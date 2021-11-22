BINARY_PATH=bin/
BINARY=portal

.PHONY: all build run gotool clean help

default: help

help:  ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

all:  ## all
	gotool
	checktool
	build

build:  ## compile Go src, generate bin file
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_PATH}${BINARY} ./cmd
	@echo "[OK] Server was build!"

run:  ## execute Go src
	@go run ./

gotool:  ## execute Go tool 'fmt' and 'vet'
	go mod tidy
	go fmt ./...
	go vet ./...
	go clean -testcache
	go test -v -cover -covermode=atomic ./...

checktool:  ## execute check tool
	golangci-lint run ./...

clean:  ## rm bin file and vim swap files
	rm -rf ${BINARY_PATH}

swagger:  ## swagger
	swag init

start:  ## execute bin file
	./${BINARY_PATH}${BINARY}