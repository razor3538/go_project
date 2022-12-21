package services

import (
	"errors"
	"example.com/m/domain"
	"example.com/m/internal/models"
	"example.com/m/internal/repository"
	"example.com/m/tools"
	"github.com/jinzhu/gorm"
)

// UserService struct
type UserService struct{}

// NewUserService return pointer to
// user struct with all methods
func NewUserService() *UserService {
	return &UserService{}
}

var userRepo = repository.NewUserRepo()

func (u *UserService) Add(user domain.User) (models.User, error) {
	user.Password = tools.HashString(user.Password)

	_, err := userRepo.GetByKey("email", user.Email)

	if gorm.IsRecordNotFoundError(err) {
		userResult, userErr := userRepo.Add(user)

		if userErr != nil {
			return models.User{}, err
		}

		userModel, err := userRepo.GetByID(userResult.ID.String())

		return userModel, err
	}
	return models.User{}, errors.New("пользователь с такой почтой уже существует")
}
