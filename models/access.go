package models

import "time"

type Access struct {
	ID        int
	UserID    int
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type AccessRepository interface {
	CreateAccess(access *Access) error
	FindAccessByToken(token string) (*Access, error)
}
