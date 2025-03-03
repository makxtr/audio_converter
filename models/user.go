package models

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	CreateUser(user *User) error
	FindByEmail(email string) (*User, error)
}
