#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
PROGRAM_NAME=golazy


build:
	rm -f ${PROGRAM_NAME}
	go mod tidy && go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o ${PROGRAM_NAME}

gen: install
	rm -rf example && golazy api go -f example.api -o example && cp example.api example

test: gen
	cd example && go mod tidy && go run .

install: build
	sudo cp golazy ${GOPATH}/bin

.PHONY: build install start gen

BINS+=${PROGRAM_NAME}
