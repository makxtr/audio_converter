package handlers

import (
	"net/http"
)

// HealthCheckHandler healthCheckHandler - обработчик для тестового эндпоинта
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running!"))
}
