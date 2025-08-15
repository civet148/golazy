FROM golang:latest AS builder
MAINTAINER lory

# 放开下面的注释允许docker通过git账户和私有仓库代码构建镜像
#ARG GIT_USER
#ARG GIT_PASSWORD
#RUN echo "machine git.your-company.com" >> ~/.netrc
#RUN echo "login $GIT_USER" >> ~/.netrc
#RUN echo "password $GIT_PASSWORD" >> ~/.netrc
#RUN cat ~/.netrc

RUN apt-get update && apt-get install -y ca-certificates make
ENV SRC_DIR /code
RUN set -x \
  && cd /tmp

RUN go env -w GOPROXY=https://goproxy.io

COPY . $SRC_DIR
RUN cd $SRC_DIR && export GIT_SSL_NO_VERIFY=true && git config --global http.sslVerify "false" && make

FROM ubuntu:22.04

RUN ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone && apt-get update && apt-get install -y tzdata
ENV TZ Asia/Shanghai
ENV SRC_DIR /code


# 管理系统主程序
COPY --from=builder $SRC_DIR/example /usr/local/bin/example
COPY --from=builder /etc/ssl/certs /etc/ssl/certs


ENV HOME_PATH /data

VOLUME $HOME_PATH

