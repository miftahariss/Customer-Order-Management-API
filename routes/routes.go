// routes/routes.go
package routes

import (
	"kstyle-hub/controller"
	"kstyle-hub/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, cc *controller.CustomerController, oc *controller.OrderController, ac *controller.AuthController) {
	e.POST("/auth/register", ac.Register)
	e.POST("/auth/login", ac.Login)

	customerGroup := e.Group("/customers")
	customerGroup.Use(middleware.JWTMiddleware())
	customerGroup.GET("", cc.GetCustomers)
	customerGroup.GET("/:id", cc.GetCustomer)
	customerGroup.POST("", cc.CreateCustomer)
	customerGroup.PUT("/:id", cc.UpdateCustomer)
	customerGroup.DELETE("/:id", cc.DeleteCustomer)
	customerGroup.GET("/search", cc.SearchCustomers)

	orderGroup := e.Group("/orders")
	orderGroup.Use(middleware.JWTMiddleware())
	orderGroup.GET("", oc.GetOrders)
	orderGroup.GET("/:id", oc.GetOrder)
	orderGroup.POST("", oc.CreateOrder)
	orderGroup.PUT("/:id", oc.UpdateOrder)
	orderGroup.DELETE("/:id", oc.DeleteOrder)
	orderGroup.GET("/search", oc.SearchOrders)
}
