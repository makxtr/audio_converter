package models

import "time"

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserAccess struct {
	ID        int
	UserID    int
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type UserRepository interface {
	CreateUser(user *User) error
	FindByEmail(email string) (*User, error)
	CreateUserAccess(access *UserAccess) error
}
