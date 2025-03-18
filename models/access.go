package models

import "time"

type Access struct {
	ID        int
	UserID    int
	Token     Token
	CreatedAt time.Time
	ExpiresAt time.Time
}

type AccessRepository interface {
	CreateAccess(access *Access) error
	FindAccessByToken(token string) (*Access, error)
}
