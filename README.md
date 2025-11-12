## Клонирование репозитория

Сначала клонируем проект и переключаемся на ветку `dev`:

```bash
git clone git@gitlab.com:plebstomsk/wow-ruby.server.git
cd wow-ruby.server
git checkout dev
git pull origin dev
```

В myapp обновляем сваггер
```
swag init -g cmd/main.go
```

Собираем Docker образы сервера и базы данных:
```
docker compose -f testing.docker-compose.yml build --no-cache
```

Поднимаем контейнеры:
```
docker compose -f testing.docker-compose.yml up -d
```