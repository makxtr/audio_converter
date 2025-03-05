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
		log.Println("File .env not found, using environment variables")
	}

	App = Config{
		DBHost:     os.Getenv("MYSQL_HOST"),
		DBPort:     os.Getenv("MYSQL_PORT"),
		DBUser:     os.Getenv("MYSQL_USER"),
		DBPassword: os.Getenv("MYSQL_PASSWORD"),
		DBName:     os.Getenv("MYSQL_DATABASE"),
		ServAddr:   os.Getenv("SERVER_ADDR"),
	}
	
	if App.DBHost == "" || App.DBPort == "" || App.DBUser == "" || App.DBPassword == "" || App.DBName == "" {
		log.Fatal("Environment variables not set for MySQL")
	}
}
