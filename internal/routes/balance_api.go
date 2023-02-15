package routes

import (
	"example.com/m/internal/repository"
	"example.com/m/internal/services"
	"example.com/m/tools"
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
// @Success  200      {object}  []domain.Balance
// @Failure  400      {object}  models.Error
// @Router   /api/user/balance [get]
func (b *Balance) Get(c *gin.Context) {
	id, err := tools.ExtractTokenID(c)

	if err != nil {
		tools.CreateError(100, err, c)
		return
	}

	user, err := repository.NewUserRepo().GetByID(id)

	if err != nil {
		tools.CreateError(101, err, c)
		return
	}

	orderModel, err := balanceService.Get(user.ID.String())

	if err != nil {
		tools.CreateError(102, err, c)
		return
	}

	c.JSON(http.StatusCreated, orderModel)
}
