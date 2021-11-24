BINARY_PATH=bin/
BINARY_FILE=portal

.PHONY: all build run gotool clean help

all: gotool checktool build

help:  ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

build:  ## compile Go src, generate bin file
	go build -race -ldflags -s -o ${BINARY_PATH}${BINARY_FILE} ./cmd
	@echo "[OK] Server was build!"

gotool:  ## execute Go tool 'fmt' and 'vet'
	go mod tidy
	go fmt ./...
	go vet ./...
	go clean -testcache
	go test -v -cover -covermode=atomic ./...

checktool:  ## execute check tool
	gocritic check ./...
	golangci-lint run ./...

clean:  ## rm bin file and vim swap files
	rm -rf ${BINARY_PATH}

swagger:  ## swagger
	swag init

start:  ## execute bin file
	./${BINARY_PATH}${BINARY_FILE}