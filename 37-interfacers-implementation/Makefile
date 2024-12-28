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

mockgen:
	mockgen -source=shape/shape.go -destination=shape/mocks/mock_ishape.go -package=mocks

.PHONY: build clean test all mockgen