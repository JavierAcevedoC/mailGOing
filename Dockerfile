# Etapa 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o mailgoing cmd/main.go

# Etapa 2: Runtime
FROM alpine:latest

WORKDIR /app

COPY .env ./
COPY --from=builder /app/mailgoing .
EXPOSE 8090

CMD ["./mailgoing"]

