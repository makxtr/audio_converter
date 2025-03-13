package repository

import (
	"audio_converter/models"
	"database/sql"
	"errors"
)

type AccessRepositorySQL struct {
	db *sql.DB
}

func NewAccessRepository(db *sql.DB) *AccessRepositorySQL {
	return &AccessRepositorySQL{db: db}
}

func (r *AccessRepositorySQL) FindAccessByToken(token string) (*models.Access, error) {
	query := `SELECT user_id, token, expires_at FROM user_access WHERE token = ?`
	access := &models.Access{}
	err := r.db.QueryRow(query, token).Scan(&access.UserID, &access.Token, &access.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return access, nil
}

func (r *AccessRepositorySQL) CreateAccess(access *models.Access) error {
	_, err := r.db.Exec(
		"INSERT INTO user_access (user_id, token, expires_at) VALUES (?, ?, ?)",
		access.UserID, access.Token, access.ExpiresAt,
	)
	if err != nil {
		return errors.New("Access not created: " + err.Error())
	}
	return nil
}
