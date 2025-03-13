package repository

import (
	"audio_converter/models"
	"database/sql"
	"errors"
)

type UserRepositorySQL struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositorySQL {
	return &UserRepositorySQL{db: db}
}

func (r *UserRepositorySQL) CreateUser(user *models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.Password,
	)
	if err != nil {
		return errors.New("user not created: " + err.Error())
	}
	return nil
}

func (r *UserRepositorySQL) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepositorySQL) FindUserAccessByToken(token string) (*models.UserAccess, error) {
	query := `SELECT user_id, token, expires_at FROM user_access WHERE token = ?`
	userAccess := &models.UserAccess{}
	err := r.db.QueryRow(query, token).Scan(&userAccess.UserID, &userAccess.Token, &userAccess.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return userAccess, nil
}

func (r *UserRepositorySQL) CreateUserAccess(access *models.UserAccess) error {
	_, err := r.db.Exec(
		"INSERT INTO user_access (user_id, token, expires_at) VALUES (?, ?, ?)",
		access.UserID, access.Token, access.ExpiresAt,
	)
	if err != nil {
		return errors.New("Access not created: " + err.Error())
	}
	return nil
}
