## We specify the base image we need for our
## go application
FROM golang:1.16.7-alpine3.14
## We create an /app directory within our
## image that will hold our application source
## files

RUN mkdir /app

## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app/cmd/registryService
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .

EXPOSE 3001
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/cmd/registryService/main"]