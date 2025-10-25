package database

import (
	"audio_converter/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.App.DB.User,
		config.App.DB.Password,
		config.App.DB.Host,
		"3306", config.App.DB.Name,
	)

	var err error
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("MySQL connection error: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("MySQL Ping error: %w", err)
	}

	fmt.Println("✅ Successful connection to MySQL!")
	return db, nil
}
