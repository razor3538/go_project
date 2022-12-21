package domain

import "github.com/gofrs/uuid"

// Balance model
type Balance struct {
	Base
	Current   float64   `gorm:"type:int" json:"current"`
	Withdrawn float64   `gorm:"type:int" json:"withdrawn"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
}
