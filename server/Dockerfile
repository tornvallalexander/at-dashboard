FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

ENV PORT 8080
ENV GIN_MODE release

RUN go build -o /golang-typing-api

CMD ["/golang-typing-api"]