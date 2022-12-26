package routes

import (
	"example.com/m/internal/repository"
	"example.com/m/internal/services"
	"example.com/m/middleware"
	"example.com/m/tools"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Balance struct
type Balance struct{}

// NewBalance return pointer to
// balance struct with all methods
func NewBalance() *Balance {
	return &Balance{}
}

var balanceService = services.NewBalanceService()

// Get return balance
// @Summary  return error or 200 status code if success
// @Produce  json
// @Accept   json
// @Tags     balance
// @Success  200      {object}  []domain.Balance  false    "Order"
// @Failure  400      {object}  swagger.Error          "Error"
// @Router   /api/user/balance [get]
func (b *Balance) Get(c *gin.Context) {
	token, _ := c.Cookie("jwt")

	value, _ := middleware.Passport().ParseTokenString(token)

	id := jwt.ExtractClaimsFromToken(value)["id"]

	user, err := repository.NewUserRepo().GetByID(id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	orderModel, err := balanceService.Get(user.ID.String())

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, orderModel)
}
