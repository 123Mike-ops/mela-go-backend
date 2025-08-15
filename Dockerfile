FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


COPY . .

RUN go build -o auth-sso

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/auth-sso .

EXPOSE 8080

CMD ["./auth-sso"]
