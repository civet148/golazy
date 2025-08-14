#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
PROGRAM_NAME=golazy

install: build
	sudo cp golazy /usr/local/bin
.PHONY: install

build:
	rm -f ${PROGRAM_NAME}
	go mod tidy && go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o ${PROGRAM_NAME}

.PHONY: build
BINS+=${PROGRAM_NAME}
