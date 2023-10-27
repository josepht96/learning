FROM golang:1.18-alpine

WORKDIR /scout

COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go build -o /scout

EXPOSE 8080

CMD [ "/scout" ]