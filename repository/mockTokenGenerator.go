package repository

import (
	"audio_converter/models"
)

type MockTokenGenerator struct {
	Token *models.Token
}

func (m *MockTokenGenerator) GenerateToken() *models.Token {
	return m.Token
}
