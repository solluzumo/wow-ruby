# ---- Build stage ----
FROM golang:1.24-alpine AS builder
WORKDIR /app/myapp

# Копируем зависимости
COPY myapp/go.mod myapp/go.sum ./
RUN go mod download

# Копируем весь код
COPY myapp .

# Собираем бинарник
RUN go build -o main ./cmd

# ---- Final stage ----
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/myapp/main ./main
RUN apk add --no-cache ca-certificates

EXPOSE 8080
CMD ["./main"]
