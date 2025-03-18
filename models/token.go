package models

import (
	"audio_converter/utils"
	"time"
)

// TokenGenerator — интерфейс для генерации токенов
type TokenGenerator interface {
	GenerateToken() *Token
}

type Token struct {
	Value     string
	ExpiresAt time.Time
}

func (t *Token) GenerateToken() *Token {
	return &Token{
		Value:     utils.GenToken(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
}
