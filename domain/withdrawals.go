package domain

import "github.com/gofrs/uuid"

// Withdrawals model
type Withdrawals struct {
	Base
	Sum    float64   `gorm:"type:float" json:"sum"`
	Order  string    `gorm:"type:text;unique" json:"order"`
	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
}
