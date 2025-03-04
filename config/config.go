package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServAddr   string
}

var App Config

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Файл .env не найден, используем переменные окружения")
	}

	App = Config{
		DBHost:     os.Getenv("MYSQL_HOST"),
		DBPort:     os.Getenv("MYSQL_PORT"),
		DBUser:     os.Getenv("MYSQL_USER"),
		DBPassword: os.Getenv("MYSQL_PASSWORD"),
		DBName:     os.Getenv("MYSQL_DATABASE"),
		ServAddr:   os.Getenv("SERVER_ADDR"),
	}

	// Проверяем, что все переменные заданы
	if App.DBHost == "" || App.DBPort == "" || App.DBUser == "" || App.DBPassword == "" || App.DBName == "" {
		log.Fatal("Не заданы переменные окружения для MySQL")
	}
}
