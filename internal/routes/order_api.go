package routes

import (
	"errors"
	"example.com/m/domain"
	"example.com/m/internal/repository"
	"example.com/m/internal/services"
	"example.com/m/tools"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io"
	"net/http"
	"strconv"
)

// Order struct
type Order struct{}

// NewOrder return pointer to
// user struct with all methods
func NewOrder() *Order {
	return &Order{}
}

var orderService = services.NewOrderService()

// Add create order
// @Summary  return error or 201 status code if success
// @Produce  json
// @Accept   json
// @Tags     order
// @Param body body swagger.CreateOrder false "order"
// @Success  201      {object}  domain.Order
// @Failure  400      {object}  models.Error
// @Router   /api/user/orders [post]
func (o *Order) Add(c *gin.Context) {
	var order domain.Order
	var reader = c.Request.Body

	b, err := io.ReadAll(reader)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	body := string(b)

	id, err := tools.ExtractTokenID(c)

	if err != nil {
		tools.CreateError(http.StatusUnauthorized, err, c)
		return
	}

	user, err := repository.NewUserRepo().GetByID(id)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	order.Number = body
	order.UserID = user.ID
	order.Status = "NEW"

	number, err := strconv.Atoi(order.Number)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, errors.New("неверный формат запроса"), c)
		return
	}

	ok := tools.Valid(number)
	if !ok {
		tools.CreateError(422, errors.New("неверный формат номера заказа"), c)
		return
	}

	orderModel, err := orderService.Add(order)

	if err != nil {
		tools.CreateError(http.StatusOK, err, c)
		return
	}

	c.JSON(http.StatusAccepted, orderModel)
}

// Get return orders
// @Summary  return error or 200 status code if success
// @Produce  json
// @Accept   json
// @Tags     order
// @Success  200      {object}  []domain.Order
// @Failure  400      {object}  models.Error
// @Router   /api/user/orders [get]
func (o *Order) Get(c *gin.Context) {
	id, err := tools.ExtractTokenID(c)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	orderModel, err := orderService.GetAllByUser(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		print("good")
		c.Status(http.StatusOK)
		return
	} else {
		print("bad")
	}

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusOK, orderModel)
}
