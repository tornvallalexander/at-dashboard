FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main


FROM alpine:latest AS production

ENV PORT 8080
ENV GIN_MODE release

COPY --from=builder /app .

CMD ["./main"]