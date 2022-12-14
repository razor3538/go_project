package routes

import (
	"errors"
	"example.com/m/domain"
	"example.com/m/internal/models"
	"example.com/m/internal/services"
	"example.com/m/tools"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

// User struct
type User struct{}

// NewUser return pointer to
// user struct with all methods
func NewUser() *User {
	return &User{}
}

var userService = services.NewUserService()

// Add create user
// @Summary  return error or 201 status code if success
// @Produce  json
// @Accept   json
// @Tags     user
// @Param    payload  body      swagger.CreateUser    false  "User"
// @Success  201      {object}  swagger.UserResponse  false  "User"
// @Failure  400      {object}  swagger.Error          "Error"
// @Router   /api/user/register [post]
func (u *User) Add(c *gin.Context) {
	var body models.CreateUserRequest
	var user domain.User

	if err := tools.RequestBinderBody(&body, c); err != nil {
		return
	}

	if err := copier.Copy(&user, &body); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	if ok := tools.IsEmailValid(user.Email); !ok {
		tools.CreateError(http.StatusBadRequest, errors.New("не валидная почта"), c)
		return
	}

	userModel, err := userService.Add(user)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	c.JSON(http.StatusCreated, userModel)
}
