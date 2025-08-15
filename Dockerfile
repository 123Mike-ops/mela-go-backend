# Build stage
FROM golang:1.25 AS builder

WORKDIR /app/cmd

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . ./

# Build the binary
RUN go build -o auth-sso ./

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the built binary
COPY --from=builder /app/auth-sso .

EXPOSE 8080

CMD ["./auth-sso"]
