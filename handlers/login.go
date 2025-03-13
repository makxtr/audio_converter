package handlers

import (
	"audio_converter/models"
	"audio_converter/utils"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

func LoginHandler(
	userRepo models.UserRepository,
	accessRepo models.AccessRepository,
) http.HandlerFunc {
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

		token := utils.GenToken()
		// Сохраняем токен в базу данных
		expiresAt := time.Now().Add(24 * time.Hour) // Токен действителен 24 часа
		userAccess := &models.Access{
			UserID:    user.ID,
			Token:     token,
			ExpiresAt: expiresAt,
		}

		if err := accessRepo.CreateAccess(userAccess); err != nil {
			http.Error(w, "Ошибка при создании токена", http.StatusInternalServerError)
			return
		}

		// Успешный вход, отправляем ID
		resp := LoginResponse{ID: user.ID, Token: token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
