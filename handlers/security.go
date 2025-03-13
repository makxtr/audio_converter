package handlers

import (
	"encoding/json"
	"net/http"
)

func SecurityHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем userID из контекста
	userID := r.Context().Value("userID").(int)

	// Возвращаем ответ
	response := map[string]interface{}{
		"message": "OK",
		"userID":  userID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
