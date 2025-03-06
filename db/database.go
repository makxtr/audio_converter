package database

import (
	"audio_converter/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.App.DB.User, config.App.DB.Password, config.App.DB.Host, "3306", config.App.DB.Name)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("MySQL connection error: %v", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)

	err = DB.Ping()
	if err != nil {
		log.Fatalf("MySQL Ping error: %v", err)
	}

	fmt.Println("✅ Successful connection to MySQL!")
}
