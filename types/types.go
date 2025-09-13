package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id string) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"userName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RegisterUserPayload struct {
	Name     string `json:"name"`
	Username string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
