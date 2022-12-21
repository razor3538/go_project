package models

import "github.com/gofrs/uuid"

// GetOrdersByUser order model for login user
type GetOrdersByUser struct {
	UserID uuid.UUID `json:"user_id" example:"" binding:"required"`
} //@name GetOrdersByUser
