package domain

import "github.com/gofrs/uuid"

// Order model
type Order struct {
	Base
	Number string    `gorm:"type:text;unique" json:"number"`
	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
}
