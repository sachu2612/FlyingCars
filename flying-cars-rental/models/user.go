package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        int64      `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// NewUser creates a new User instance with the given email and password
func NewUser(email, password string) *User {
	now := time.Now()
	return &User{
		Email:     email,
		Password:  password,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
