package repository

import (
	"audio_converter/models"
	"errors"
)

type MockAccessRepository struct {
	Access *models.Access
	Err    error
}

func (m *MockAccessRepository) CreateAccess(access *models.Access) error {
	return m.Err
}

func (m *MockAccessRepository) FindAccessByToken(token string) (*models.Access, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	if m.Access != nil && m.Access.Token == token {
		return m.Access, nil
	}
	return nil, errors.New("token not found")
}
