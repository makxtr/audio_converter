package handlers

import (
	"audio_converter/usecases/auth"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

func LoginHandler(authUC *auth.AuthUseCase) http.HandlerFunc {
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

		user, token, err := authUC.Login(req.Email, req.Password)
		if err != nil {
			switch err {
			case auth.ErrInvalidCredentials:
				http.Error(w, "Неверный email или пароль", http.StatusUnauthorized)
			case auth.ErrTokenCreation:
				http.Error(w, "Ошибка при создании токена", http.StatusInternalServerError)
			default:
				http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			}
			return
		}

		// Успешный вход, отправляем ответ
		resp := LoginResponse{ID: user.ID, Token: token.Value}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
