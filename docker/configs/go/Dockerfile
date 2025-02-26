# Используем Go-образ
FROM golang:latest

WORKDIR /app

# Копируем файлы go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Открываем порт
EXPOSE 8080

# По умолчанию запускаем сервер в режиме разработки
CMD ["go", "run", "server.go"]