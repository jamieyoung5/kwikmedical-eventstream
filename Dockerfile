# Stage 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /run-app ./cmd/main.go

# Stage 2: Use a compatible runtime image
FROM debian:bookworm-slim

COPY --from=builder /run-app /run-app

ENV PORT 4343

ENTRYPOINT ["/run-app"]
