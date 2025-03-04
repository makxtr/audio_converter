package database

import (
	"database/sql"
	"fmt"
	"log"

	"audio_converter/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.App.DBUser, config.App.DBPassword, config.App.DBHost, "3306", config.App.DBName)

	var err error
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
