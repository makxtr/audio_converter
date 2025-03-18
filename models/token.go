package models

import (
	"audio_converter/utils"
	"time"
)

type TokenGenerator interface {
	GenerateToken() *Token
}

type Token struct {
	Value     string
	ExpiresAt time.Time
}

type DefaultTokenGenerator struct{}

func (g *DefaultTokenGenerator) GenerateToken() *Token {
	return &Token{
		Value:     utils.GenToken(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
}
