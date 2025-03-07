package repository

import (
	"audio_converter/models"
	"errors"
)

type MockUserRepository struct {
	User *models.User
	Err  error
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	return m.Err
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	if m.User != nil && m.User.Email == email {
		return m.User, nil
	}
	return nil, errors.New("user not found")
}
