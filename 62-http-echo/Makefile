# Variables
GO = go
APP_NAME = app
BUILD_DIR= ./bin
TEST = $(GO) test -v
FMT= gofmt
BUILD_FLAGS = -ldflags="-w -s"
DOCKER_FILE = Dockerfile-slim
DOCKER_TAG = jpalaparthi/primesoft-demo:v0.3
DOCKER= docker


all: build

build: clean test
	mkdir bin
	$(GO) mod tidy
	$(GO) build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(APP_NAME) .

clean:
	rm -rf bin

test: 
	$(TEST) ./...

docker-build:
	$(DOCKER) build . -f $(DOCKER_FILE) -t $(DOCKER_TAG)


.PHONY: build clean test all docker-build