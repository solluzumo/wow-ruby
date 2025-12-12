FROM golang:1.24.5-alpine AS builder

# Корневая директория для сборки
WORKDIR /app

# 1. Копируем общие модули
COPY pkg ./pkg
COPY proto ./proto

# 2. Копируем код auth сервиса
COPY auth ./auth

# 3. Переходим в директорию auth — там go.mod
WORKDIR /app/auth

# 4. Тянем зависимости
RUN go mod tidy
RUN go mod download

# 5. Собираем бинарник
RUN go build -o /app/auth-service ./cmd/main.go

# ---------------------------
# Финальный легковесный образ
# ---------------------------
FROM debian:bookworm-slim

COPY --from=builder /app/auth-service /bin/auth-service

CMD ["/bin/auth-service"]
