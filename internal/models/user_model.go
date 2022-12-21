package models

import "github.com/gofrs/uuid"

// User struct for users
type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
} //@name UserResponse
