FROM golang:1.25.0-alpine3.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src/ ./src/

RUN go build -o main ./src/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]