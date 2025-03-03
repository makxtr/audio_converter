package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	// Загружаем переменные из .env (если файл существует)
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Файл .env не найден, используем переменные окружения")
	}

	// Читаем переменные окружения
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		log.Fatal("Не заданы переменные окружения для MySQL")
	}

	// Формируем DSN строку для MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, "3306", database)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к MySQL: %v", err)
	}

	// Проверяем соединение
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Ошибка соединения с MySQL: %v", err)
	}

	fmt.Println("✅ Успешное подключение к MySQL!")
}
