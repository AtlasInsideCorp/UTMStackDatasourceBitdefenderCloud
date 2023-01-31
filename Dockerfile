FROM golang:alpine

RUN apk update 

WORKDIR /usr/src/app

COPY certs .
COPY auth.go .
COPY certs.go .
COPY config.go .
COPY envVar.go .
COPY epsSyslogHerlper.go .
COPY go.mod .
COPY go.sum .
COPY main.go .
COPY server.go .

RUN go mod tidy
ENTRYPOINT go run .
