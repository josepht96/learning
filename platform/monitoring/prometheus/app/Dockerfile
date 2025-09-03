# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
COPY *.go ./

RUN go build -o /hello-server

EXPOSE 8081

CMD [ "/hello-server" ]