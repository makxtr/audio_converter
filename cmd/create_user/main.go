package main

import (
	"audio_converter/config"
	database "audio_converter/db"
	"audio_converter/models"
	"audio_converter/repository"
	"audio_converter/utils"
	_ "database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	config.Init()
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	if len(os.Args) != 4 {
		fmt.Println("Use: go run cmd/create_user/main.go <name> <email> <password>")
		os.Exit(1)
	}
	name := os.Args[1]
	email := os.Args[2]
	password := os.Args[3]

	salt := utils.GetSalt()
	hashedPassword := utils.HashPass(salt, password)
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hex.EncodeToString(hashedPassword),
	}
	if err := repository.NewUserRepository(db).CreateUser(user); err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	fmt.Println("✅ User successfully created!")
	fmt.Println("🔑 Generated salt (hex):", hex.EncodeToString(salt))
}
