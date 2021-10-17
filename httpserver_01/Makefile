BINARY_NAME="httpserver"

build:
	echo "building httpserver"
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux /usr/local/go/bin/go build -o ${BINARY_NAME} .

release: build
	echo "building httpserver image"
	docker build -t geektrainingcamp/httpserver:v1 .

