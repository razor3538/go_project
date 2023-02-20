package models

// CreateUserRequest user model for /login request route
type CreateUserRequest struct {
	Email    string `json:"login" example:"admin@mail.ru" binding:"required"`
	Password string `json:"password" example:"123" binding:"required"`
} //@name CreateUserRequest
