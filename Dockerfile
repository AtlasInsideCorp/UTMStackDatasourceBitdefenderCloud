FROM golang:alpine

RUN apk update 

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy
RUN go build
