#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
PROGRAM_NAME=golazy

install: build
	sudo cp golazy /usr/local/bin

build:
	rm -f ${PROGRAM_NAME}
	go mod tidy && go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o ${PROGRAM_NAME}

gen: build
	rm -rf example && golazy api go -f example.api -o example

start:
	cd example && go mod tidy && go run .

.PHONY: build install start gen

BINS+=${PROGRAM_NAME}
