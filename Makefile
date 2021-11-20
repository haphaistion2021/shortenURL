BINARY_PATH=bin/
BINARY=portal

.PHONY: all build run gotool clean help

all: gotool checktool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_PATH}${BINARY} ./cmd
	@echo "[OK] Server was build!"

run:
	@go run ./

gotool:
	go mod tidy
	go fmt ./...
	go vet ./...
	go clean -testcache
	go test -v -cover -covermode=atomic ./...

checktool:
	golangci-lint run ./...

clean:
	rm -rf ${BINARY_PATH}

swagger:
	swag init

help:
	@echo "make - format Go src, then compile as bin file"
	@echo "make build - compile Go src, generate bin file"
	@echo "make run - execute Go src"
	@echo "make clean - rm bin file and vim swap files"
	@echo "make gotool - execute Go tool 'fmt' and 'vet'"
	@echo "make start - execute bin file"
	@echo "make checktool - execute check tool"

start:
	./${BINARY_PATH}${BINARY}