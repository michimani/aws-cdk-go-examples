.PHONY: build

build:
	cd src/lambda/simple-response \
	&& GOARCH=amd64 GOOS=linux go build -o bin/default/main -ldflags "-X \"main.message=Default Response\"" \
	&& GOARCH=amd64 GOOS=linux go build -o bin/hello/main -ldflags "-X \"main.message=Hello Lambda!\"" \
	&& GOARCH=amd64 GOOS=linux go build -o bin/bye/main -ldflags "-X \"main.message=Goodbye Lambda!\"" \