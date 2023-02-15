package routes

import (
	"example.com/m/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter setting up gin router and config
func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := NewAuth()
	user := NewUser()
	order := NewOrder()
	balance := NewBalance()
	withdraw := NewWithdraw()

	r.POST("/api/user/register", user.Add)
	r.GET("/api/user/is-authenticated", auth.IsAuthenticated)
	r.POST("/api/user/login", middleware.Passport().LoginHandler)
	r.GET("/api/user/logout", middleware.Passport().LogoutHandler)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRequired := r.Group("/")
	authRequired.Use(middleware.JwtAuthMiddleware())
	{
		authRequired.POST("/api/user/orders", order.Add)
		authRequired.GET("/api/user/orders", order.Get)
		authRequired.GET("/api/user/balance", balance.Get)
		authRequired.POST("/api/user/balance/withdraw", withdraw.Pay)
		authRequired.GET("/api/user/withdrawals", withdraw.Get)
	}
	return r
}
