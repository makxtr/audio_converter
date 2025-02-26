FROM golang:latest AS builder

WORKDIR /app

# Устанавливаем Goose и Atlas
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install ariga.io/atlas/cmd/atlas@latest

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
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /go/bin/atlas /usr/local/bin/atlas

# Открываем порт приложения
EXPOSE 8080

# Устанавливаем команду по умолчанию
ENTRYPOINT ["./main"]