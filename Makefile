.PHONY: start build

# Go parameters
# NOW = $(shell date -u '+%Y%m%d%I%M%S')
# GO_PATH=$(shell which go)
MAIN_PATH=./main.go
BUILD_PATH=./build/package/
CONFIG_FILE=./configs/config.toml
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_RUN=$(GO_CMD) run
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
# BINARY_NAME=jobob-amd64
BINARY_NAME=app
BINARY_LINUX=linux-amd64
BINARY_WINDOWS=windows-amd64.exe
BINARY_DARWIN=darwin-amd64
BINARY_ARM=app-arm
BINARY_UNIX=$(BINARY_NAME)_unix


all: clean build build-linux build-windows

# -ldflags "-w -s"
build:
	go build -o $(BUILD_PATH)$(BINARY_NAME) $(MAIN_PATH)

#swagger:
#  bash swag init --generalInfo cmd/main/main.go --exclude cmd/main/test
#       swag init -g ./internal/app/routers/swagger.go -o ./docs/swagger
#

test:
	$(GO_TEST) -cover -race ./...

run:
	$(GO_RUN) $(MAIN_PATH) -c $(CONFIG_FILE)

start:
	$(GO_BUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_LINUX)
	rm -f $(BINARY_WINDOWS)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BUILD_PATH)$(BINARY_LINUX) $(MAIN_PATH)
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO_BUILD) -o $(BUILD_PATH)$(BINARY_WINDOWS) $(MAIN_PATH)
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o $(BUILD_PATH)$(BINARY_DARWIN) $(MAIN_PATH)
build-android:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GO_BUILD) -o $(BUILD_PATH)$(BINARY_ARM) $(MAIN_PATH)

build-docker:
	docker run --rm -it -v "$(GO_PATH)":/go -w /t/a/b/c golang:latest go build -o $(BINARY_NAME) $(MAIN_PATH)

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"