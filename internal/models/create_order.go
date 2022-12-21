package models

// CreateOrderRequest order model for /orders request route
type CreateOrderRequest struct {
	Number string `json:"number" example:"12345678903" binding:"required"`
} //@name CreateOrderRequest
