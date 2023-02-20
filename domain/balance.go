package domain

import "github.com/gofrs/uuid"

// Balance model
type Balance struct {
	Base
	Current   float64   `gorm:"type:float" json:"current"`
	Withdrawn float64   `gorm:"type:float" json:"withdrawn"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
}
