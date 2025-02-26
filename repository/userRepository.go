package repository

import (
	"audio_converter/models"
	"database/sql"
	"errors"
)

type UserRepositorySQL struct {
	db *sql.DB
}

// Конструктор репозитория
func NewUserRepository(db *sql.DB) *UserRepositorySQL {
	return &UserRepositorySQL{db: db}
}

func (r *UserRepositorySQL) CreateUser(user *models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.Password,
	)
	if err != nil {
		return errors.New("не удалось создать пользователя: " + err.Error())
	}
	return nil
}
