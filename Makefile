# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
SOURCE_DIR=src/github.com/patrickz98/GoGo/
BINARY_DIR=bin
BINARY_NAME=binary

all: build

build:
	cd $(SOURCE_DIR) && $(GOBUILD) -o $(BINARY_NAME)

clean:
	rm -f $(BINARY_DIR)

run: build
	./$(BINARY_DIR)/$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o $(BINARY)

docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY)" -v
