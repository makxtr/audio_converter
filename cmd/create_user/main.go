package main

import (
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
	// 1️⃣ Инициализируем подключение к БД
	database.InitDB()
	defer database.DB.Close() // Закрываем соединение в конце

	// 2️⃣ Создаём репозиторий пользователя
	repo := repository.NewUserRepository(database.DB)

	// 3️⃣ Читаем аргументы командной строки
	if len(os.Args) != 4 {
		fmt.Println("Использование: go run cmd/create_user/main.go <name> <email> <password>")
		os.Exit(1)
	}
	name := os.Args[1]
	email := os.Args[2]
	password := os.Args[3]

	// 4️⃣ Генерируем соль и хешируем пароль
	salt := utils.GetSalt()
	hashedPassword := utils.HashPass(salt, password)

	// 5️⃣ Создаём пользователя в базе данных
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hex.EncodeToString(hashedPassword),
	}
	if err := repo.CreateUser(user); err != nil {
		log.Fatalf("Ошибка создания пользователя: %v", err)
	}

	fmt.Println("✅ Пользователь успешно создан!")
	fmt.Println("🔑 Сгенерированная соль (hex):", hex.EncodeToString(salt))
}
