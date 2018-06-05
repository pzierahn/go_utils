# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
SOURCE_DIR=src/github.com/patrickz98/GoGo/
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
	cd $(SOURCE_DIR) && CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o $(BINARY_NAME)
	mkdir -p $(BINARY_DIR)
	mv $(SOURCE_DIR)/$(BINARY_NAME) $(BINARY_DIR)

build-linux:
	cd $(SOURCE_DIR) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)
	mkdir -p $(BINARY_DIR)
	mv $(SOURCE_DIR)/$(BINARY_NAME) $(BINARY_DIR)

docker-build:
	docker build -f Dockerfile .