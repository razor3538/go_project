package repository

import (
	"example.com/m/config"
	"example.com/m/domain"
	"example.com/m/internal/models"
)

// UserRepo struct
type UserRepo struct{}

// NewUserRepo return pointer to user repository
// with all methods
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// GetByKey return user by key
// and error if exist
func (ur UserRepo) GetByKey(key, value string) (domain.User, error) {
	var user domain.User
	err := config.DB.
		Unscoped().
		Where(key+" = ?", value).
		First(&user).Error

	return user, err
}

func (ur UserRepo) GetByID(id string) (models.User, error) {
	var user models.User
	err := config.DB.
		Table("users as u").
		Select("u.*").
		Where("u.id = ?", id).
		Scan(&user).
		Error

	return user, err
}

func (ur UserRepo) Add(user domain.User) (domain.User, error) {
	if err := config.DB.
		Create(&user).
		Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}
