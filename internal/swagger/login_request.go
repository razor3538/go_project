package swagger

// LoginRequest user model for /login request route
type LoginRequest struct {
	Username string `json:"username" example:"admin@mail.ru"`
	Password string `json:"password" example:"123"`
}
