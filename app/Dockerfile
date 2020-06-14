FROM golang:1.12-alpine as build

WORKDIR /go/app

COPY . .

RUN apk add --no-cache git \
  && go build -o app \
  && go get github.com/oxequa/realize
  
FROM alpine