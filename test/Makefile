#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
MANAGER_DIR=${PWD}
FRONTEND_CODE=/tmp/web-frontend
PROGRAM_NAME=test
IMAGE_NAME=${PROGRAM_NAME}

build:
	rm -f ${PROGRAM_NAME}
	go mod tidy && go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}'" -o ${PROGRAM_NAME}
.PHONY: build
BINS+=${PROGRAM_NAME}

docker:
	docker build --build-arg GIT_USER=${GIT_USER} --build-arg GIT_PASSWORD=${GIT_PASSWORD} --tag ${IMAGE_NAME} -f Dockerfile .
.PHONY: docker

nodejs:
	curl -sL https://deb.nodesource.com/setup_14.x | sudo -E bash - && sudo apt update && sudo apt install -y nodejs build-essential && sudo npm install -g yarn
.PHONY: nodejs

web:
	rm -rf ${FRONTEND_CODE} && git clone -b master https://git.your-enterprise.com/web-frontend.git ${FRONTEND_CODE}
	cd ${FRONTEND_CODE} && git log -3 && npm install && npm run build:prod
.PHONY: web

# 检查环境变量
env-%:
	@ if [ "${${*}}" = "" ]; then \
	    echo "Environment variable $* not set"; \
	    exit 1; \
	fi

db2go:
	go install github.com/civet148/db2go@latest
.PHONY: db2go

clean:
	rm -rf $(CLEAN) $(BINS)
.PHONY: clean

