FROM golang:latest AS builder

WORKDIR /app

# Копируем файлы и загружаем зависимости
COPY go.mod ./
RUN go mod download

# Копируем исходный код
COPY . /app/

# Сборка бинарного файла
RUN CGO_ENABLED=0 go build -o main .

# Финальный контейнер
FROM golang:latest

WORKDIR /app

# Копируем бинарник из builder stage
COPY --from=builder /app/main .

# Копируем исходные файлы для тестирования
COPY . .

# Открываем порт приложения
EXPOSE 8080

# Устанавливаем команду по умолчанию
ENTRYPOINT ["./main"]