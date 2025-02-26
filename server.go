package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// healthCheckHandler - обработчик для тестового эндпоинта
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running!"))
}

func startServer() {
	port := "0.0.0.0:8080"
	fmt.Println("Starting server on", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheckHandler)

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
