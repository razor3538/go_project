package routes

import (
	"example.com/m/domain"
	"example.com/m/internal/models"
	"example.com/m/internal/repository"
	"example.com/m/internal/services"
	"example.com/m/middleware"
	"example.com/m/tools"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
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
	var body models.CreateOrderRequest
	var order domain.Order

	if err := tools.RequestBinderBody(&body, c); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	if err := copier.Copy(&order, &body); err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	token, _ := c.Cookie("jwt")

	value, _ := middleware.Passport().ParseTokenString(token)

	id := jwt.ExtractClaimsFromToken(value)["id"]

	user, err := repository.NewUserRepo().GetByID(id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	order.UserID = user.ID

	//number, err := strconv.Atoi(order.Number)
	//if err != nil {
	//	tools.CreateError(http.StatusBadRequest, errors.New("неверный формат запроса"), c)
	//	return
	//}

	//ok := tools.Valid(number)
	//if !ok {
	//	tools.CreateError(422, errors.New("неверный формат номера заказа"), c)
	//	return
	//}

	orderModel, err := orderService.Add(order)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, orderModel)
}

// Get return orders
// @Summary  return error or 200 status code if success
// @Produce  json
// @Accept   json
// @Tags     order
// @Success  200      {object}  []domain.Order
// @Failure  400      {object}  models.Error
// @Router   /api/user/get-orders [get]
func (o *Order) Get(c *gin.Context) {
	token, _ := c.Cookie("jwt")

	value, _ := middleware.Passport().ParseTokenString(token)

	id := jwt.ExtractClaimsFromToken(value)["id"]

	user, err := repository.NewUserRepo().GetByID(id.(string))

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	orderModel, err := orderService.GetAllByUser(user.ID.String())

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, orderModel)
}
