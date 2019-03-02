# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
SOURCE_DIR=./GoGo/
BINARY_DIR=build
BINARY_NAME=main

all: build

clean:
	rm -rf $(BINARY_DIR)

build: clean
	cd $(SOURCE_DIR) && $(GOBUILD) -o $(BINARY_NAME)
	mkdir -p $(BINARY_DIR)
	mv $(SOURCE_DIR)/$(BINARY_NAME) $(BINARY_DIR)

run: build
	./$(BINARY_DIR)/$(BINARY_NAME)

# Cross compilation
build-odroid:
	cd $(SOURCE_DIR) && CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o $(BINARY_NAME)_odroid.exec
	mkdir -p $(BINARY_DIR)
	mv $(SOURCE_DIR)/$(BINARY_NAME)_odroid.exec $(BINARY_DIR)

build-linux:
	cd $(SOURCE_DIR) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_linux.exec
	mkdir -p $(BINARY_DIR)
	mv $(SOURCE_DIR)/$(BINARY_NAME)_linux.exec $(BINARY_DIR)

build-windows:
	cd $(SOURCE_DIR) && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_windows.exec
	mkdir -p $(BINARY_DIR)
	mv $(SOURCE_DIR)/$(BINARY_NAME)_windows.exec $(BINARY_DIR)

build-all: build build-odroid build-linux build-windows

docker-build:
	docker build -f Dockerfile .

deps:
	go get github.com/SSSaaS/sssa-golang
