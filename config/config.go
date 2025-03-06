package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Server struct {
		Addr string
	}
}

var App Config

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("File .env not found, using environment variables")
	}

	App = Config{
		DB: struct {
			Host     string
			Port     string
			User     string
			Password string
			Name     string
		}{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Name:     os.Getenv("MYSQL_DATABASE"),
		},
		Server: struct {
			Addr string
		}{
			Addr: os.Getenv("SERVER_ADDR"),
		},
	}

	if App.DB.Host == "" || App.DB.Port == "" || App.DB.User == "" || App.DB.Password == "" || App.DB.Name == "" {
		log.Fatal("Environment variables not set for MySQL")
	}
}
