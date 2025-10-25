package main

import (
	"audio_converter/config"
	database "audio_converter/db"
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
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	port := config.App.Server.Addr
	fmt.Println("Starting server on", port)

	userRepo := repository.NewUserRepository(db)
	accessRepo := repository.NewAccessRepository(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheckHandler)
	mux.HandleFunc("/login", handlers.LoginHandler(auth.NewAuthUseCase(userRepo, accessRepo, &models.DefaultTokenGenerator{})))

	securityHandler := http.HandlerFunc(handlers.SecurityHandler)
	mux.Handle("/security", middleware.AuthMiddleware(accessRepo)(securityHandler))

	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

func main() {
	if os.Getenv("GO_TEST") == "true" {
		fmt.Println("Skipping server start for tests")
		return
	}

	startServer()
}
