FROM golang:latest

WORKDIR /app

# Устанавливаем goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

EXPOSE 8080

# По умолчанию запускаем сервер в режиме разработки
CMD ["go", "run", "server.go"]