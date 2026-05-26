# syntax=docker/dockerfile:1

FROM golang:1.21.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./*.go ./
COPY ./test-workflow-manual.bpmn ./

RUN go build -o /zeebe-worker
EXPOSE 8080

CMD [ "/zeebe-worker" ]