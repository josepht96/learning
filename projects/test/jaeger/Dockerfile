FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY *.go ./
# RUN go get github.com/josepht96/learning/projects/test/jaeger
RUN go build -o /jaeger-test

RUN chmod +x /jaeger-test

CMD [ "/jaeger-test" ]

FROM golang:1.21-alpine
RUN apk add curl
COPY --from=build /jaeger-test /bin/jaeger-test
EXPOSE 8080
CMD ["/bin/jaeger-test"]