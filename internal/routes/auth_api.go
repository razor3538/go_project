package routes

import (
	"example.com/m/internal/services"
	"example.com/m/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth struct
type Auth struct{}

// NewAuth return pointer to Auth struct
// with all methods
func NewAuth() *Auth {
	return &Auth{}
}

var authService = services.NewAuthService()

// IsAuthenticated check if user is authorized and
// if user exists
// @Summary  return user info if authorized
// @Produce  json
// @Accept   json
// @Tags     auth
// @Success  200  {object}  swagger.UserResponse
// @Failure  401  {object}  swagger.Error
// @Failure  404  {object}  swagger.Error
// @Router   /api/user/is-authenticated [get]
func (a Auth) IsAuthenticated(c *gin.Context) {
	user, code, err := authService.IsAuthenticated(c)

	if err != nil {
		tools.CreateError(code, err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary  Returns info about user
// @Produce  json
// @Accept   json
// @Tags     auth
// @Param    payload  body      swagger.LoginRequest  false  "User Credentials"
// @Success  200      {object}  swagger.UserResponse
// @Failure  401      {object}  swagger.Error  "Error"
// @Router   /api/user/login [post]
// nolint:deadcode, unused
func login() {}

// @Summary  Removes cookie if set
// @Produce  json
// @Accept   json
// @Tags     auth
// @Success  200  {object}  swagger.Error  "Success"
// @Failure  401  {object}  swagger.Error  "Error"
// @Router   /api/user/logout [get]
// nolint:deadcode, unused
func logout() {}
