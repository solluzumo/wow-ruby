FROM golang:1.24.5-alpine AS builder

# Корневая директория для сборки
WORKDIR /app

# 1. Копируем общие модули
COPY pkg ./pkg
COPY proto ./proto

# 2. Копируем код webapi сервиса
COPY webapi ./webapi

# 3. Переходим в директорию webapi — там go.mod
WORKDIR /app/webapi

# 4. Тянем зависимости
RUN go mod tidy
RUN go mod download

# 5. Собираем бинарник
RUN go build -o /app/webapi-service ./cmd/main.go

# ---------------------------
# Финальный легковесный образ
# ---------------------------
FROM debian:bookworm-slim

COPY --from=builder /app/webapi-service /bin/webapi-service

CMD ["/bin/webapi-service"]
