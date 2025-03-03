package main

import (
	"audio_converter/db"
	"audio_converter/handlers"
	"audio_converter/repository"
	"fmt"
	"log"
	"net/http"
	"os"
)

func startServer() {
	port := "0.0.0.0:8080"
	fmt.Println("Starting server on", port)

	database.InitDB()
	defer database.DB.Close()

	userRepo := repository.NewUserRepository(database.DB)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheckHandler)
	mux.HandleFunc("/login", handlers.LoginHandler(userRepo))

	// Запускаем сервер
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

func main() {
	// Если GO_TEST=true, сервер НЕ запускается
	if os.Getenv("GO_TEST") == "true" {
		fmt.Println("Skipping server start for tests")
		return
	}

	// Запускаем сервер
	startServer()
}
