.PHONY: build

build:
	cd src/lambda/hello-function && GOARCH=amd64 GOOS=linux go build -o bin/main
	cd src/lambda/dlq-function && GOARCH=amd64 GOOS=linux go build -o bin/main