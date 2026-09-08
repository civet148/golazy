#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
PROGRAM_NAME=golazy
BIN_DIR=bin

# 创建bin目录
$(shell mkdir -p $(BIN_DIR))

build:
	rm -f ${PROGRAM_NAME}
	go mod tidy && go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o ${PROGRAM_NAME}

# Linux AMD64 编译
linux:
	@echo "Building for Linux AMD64..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o ${BIN_DIR}/${PROGRAM_NAME}-linux-amd64
	@echo "Build complete: ${BIN_DIR}/${PROGRAM_NAME}-linux-amd64"

# Mac AArch64 (Apple Silicon) 编译
mac:
	@echo "Building for Mac AArch64..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o ${BIN_DIR}/${PROGRAM_NAME}-darwin-arm64
	@echo "Build complete: ${BIN_DIR}/${PROGRAM_NAME}-darwin-arm64"

# 编译所有平台
all: linux mac
	@echo "All builds complete!"

gen-api: install
	rm -rf example && golazy api go -f example.api -o example && cp example.api example

test: gen
	cd example && go mod tidy && go run .

install: build
	sudo mv golazy ${GOPATH}/bin

.PHONY: build install start gen linux mac all

BINS+=${PROGRAM_NAME}