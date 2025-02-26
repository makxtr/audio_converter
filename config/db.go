package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

// InitDB - подключение к БД и применение миграций
func InitDB() {
	_ = godotenv.Load() // Загружаем .env файл (если есть)

	// Формируем строку подключения
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	// Открываем подключение к БД
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %s", err)
	}

	// Проверяем соединение
	if err = DB.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping database: %s", err)
	}

	fmt.Println("✅ Database connected successfully")

	// Применяем миграции
	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatalf("❌ Goose dialect error: %s", err)
	}
	if err := goose.Up(DB, "db/migrations"); err != nil {
		log.Fatalf("❌ Goose migration failed: %s", err)
	}

	fmt.Println("✅ Migrations applied successfully")
}
