package middleware

import (
	"audio_converter/models"
	"context"
	"net/http"
	"time"
)

func AuthMiddleware(accessRepo models.AccessRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Token not found", http.StatusUnauthorized)
				return
			}

			access, err := accessRepo.FindAccessByToken(token)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			
			if time.Now().After(access.Token.ExpiresAt) {
				http.Error(w, "Token expire", http.StatusUnauthorized)
				return
			}

			// Добавляем информацию о пользователе в контекст запроса
			ctx := context.WithValue(r.Context(), "userID", access.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
