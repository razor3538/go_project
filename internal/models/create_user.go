package models

// CreateUserRequest user model for /login request route
type CreateUserRequest struct {
	Email    string `json:"email" example:"admin@mail.ru" binding:"required"`
	Password string `json:"password" example:"123" binding:"required"`
	Name     string `json:"name" example:"Eduard" binding:"required"`
	LastName string `json:"last_name" example:"Gaifutdinov" binding:"required"`
} //@name CreateUserRequest
