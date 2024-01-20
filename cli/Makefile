BINARY_NAME=homework

all: test clean build

build:
	go build -o $(BINARY_NAME) -v

test: 
	go test -v ./...

clean: 
	go clean
	rm -vf $(BINARY_NAME)

build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)_linux_amd64 -v
