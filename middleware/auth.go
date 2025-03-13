package middleware

import (
	"audio_converter/models"
	"context"
	"net/http"
	"time"
)

func AuthMiddleware(userRepo models.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем токен из заголовка
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Token not found", http.StatusUnauthorized)
				return
			}

			// Ищем пользователя по токену
			userAccess, err := userRepo.FindUserAccessByToken(token)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Проверяем срок действия токена
			if time.Now().After(userAccess.ExpiresAt) {
				http.Error(w, "Token expire", http.StatusUnauthorized)
				return
			}

			// Добавляем информацию о пользователе в контекст запроса
			ctx := context.WithValue(r.Context(), "userID", userAccess.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
