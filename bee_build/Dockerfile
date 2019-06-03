FROM golang:1.12.1-alpine
RUN apk add --no-cache git
RUN apk add --no-cache build-base

RUN set -x \
    # go get bee
    && go get -u github.com/astaxie/beego \
    && go get -u github.com/beego/bee 
WORKDIR /go/src