package services

import (
	"errors"
	"example.com/m/internal/models"
	"example.com/m/internal/repository"
	"example.com/m/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthService struct
type AuthService struct{}

// NewAuthService return pointer to Auth struct
// with all methods
func NewAuthService() *AuthService {
	return &AuthService{}
}

var userRepository = repository.NewUserRepo()

func (as *AuthService) IsAuthenticated(c *gin.Context) (models.User, int, error) {
	claims, err := middleware.Passport().CheckIfTokenExpire(c)

	if err != nil {
		return models.User{}, http.StatusUnauthorized, err
	}
	if int64(claims["exp"].(float64)) < middleware.Passport().TimeFunc().Unix() {
		_, _, _ = middleware.Passport().RefreshToken(c)
	}

	id := claims[middleware.IdentityKeyID]
	result, err := userRepo.GetByID(id.(string))

	if err != nil {
		return models.User{}, http.StatusUnauthorized, errors.New("токен не действителен")
	}
	return result, 0, nil
}
