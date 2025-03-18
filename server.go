package main

import (
	"audio_converter/config"
	"audio_converter/db"
	"audio_converter/handlers"
	"audio_converter/middleware"
	"audio_converter/models"
	"audio_converter/repository"
	"audio_converter/usecases/auth"
	"fmt"
	"log"
	"net/http"
	"os"
)

func startServer() {
	config.Init()
	database.InitDB()
	defer database.DB.Close()

	port := config.App.Server.Addr
	fmt.Println("Starting server on", port)

	userRepo := repository.NewUserRepository(database.DB)
	accessRepo := repository.NewAccessRepository(database.DB)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheckHandler)
	mux.HandleFunc("/login", handlers.LoginHandler(auth.NewAuthUseCase(userRepo, accessRepo, &models.Token{})))

	securityHandler := http.HandlerFunc(handlers.SecurityHandler)
	mux.Handle("/security", middleware.AuthMiddleware(accessRepo)(securityHandler))

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

	startServer()
}
