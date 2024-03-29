package routes

import (
	"example.com/m/domain"
	"example.com/m/internal/models"
	"example.com/m/internal/repository"
	"example.com/m/internal/services"
	"example.com/m/tools"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

// Withdraw struct
type Withdraw struct{}

// NewWithdraw return pointer to
// withdraw struct with all methods
func NewWithdraw() *Withdraw {
	return &Withdraw{}
}

var withdrawService = services.NewWithdrawService()

// Pay return withdrawal
// @Summary  return error or 200 status code if success
// @Produce  json
// @Accept   json
// @Tags     withdraw
// @Param    payload  body      swagger.CreateWithdrawals    false  "Order"
// @Success  200      {object}  domain.Withdrawals
// @Failure  400      {object}  models.Error
// @Router   /api/user/balance/withdraw [post]
func (w *Withdraw) Pay(c *gin.Context) {
	var body models.CreateWithdrawals
	var order domain.Withdrawals

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	if err := copier.Copy(&order, &body); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	id, err := tools.ExtractTokenID(c)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	user, err := repository.NewUserRepo().GetByID(id)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	order.UserID = user.ID

	withdrawalsModel, err := withdrawService.Pay(order)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, withdrawalsModel)
}

// Get return withdrawal
// @Summary  return error or 200 status code if success
// @Produce  json
// @Accept   json
// @Tags     withdraw
// @Success  200      {object}  []domain.Withdrawals
// @Failure  400      {object}  models.Error
// @Router   /api/user/balance/withdrawals [get]
func (w *Withdraw) Get(c *gin.Context) {
	id, err := tools.ExtractTokenID(c)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	
	user, err := repository.NewUserRepo().GetByID(id)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	withdrawalsModel, err := withdrawService.GetAllByUser(user.ID.String())

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, withdrawalsModel)
}
