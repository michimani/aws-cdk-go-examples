.PHONY: build

build:
	cd src/lambda/simple-response \
	&& GOARCH=amd64 GOOS=linux go build -o bin/default/main -ldflags "-X \"main.message=Default Response\""