package handlers

import (
	"audio_converter/models"
	"audio_converter/utils"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID int `json:"id"`
}

func LoginHandler(userRepo models.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
			return
		}

		// Найти пользователя в БД
		user, err := userRepo.FindByEmail(req.Email)
		if err != nil {
			http.Error(w, "Неверный email или пароль", http.StatusUnauthorized)
			return
		}

		passHash, err := hex.DecodeString(user.Password)
		if err != nil {
			http.Error(w, "Ошибка при декодировании пароля", http.StatusInternalServerError)
			return
		}

		// Проверка пароля
		if !utils.CheckPass(passHash, req.Password) {
			http.Error(w, "Неверный email или пароль", http.StatusUnauthorized)
			return
		}

		// Успешный вход, отправляем ID
		resp := LoginResponse{ID: user.ID}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
