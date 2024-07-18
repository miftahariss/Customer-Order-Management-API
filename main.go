package main

import (
	"kstyle-hub/config"
	"kstyle-hub/controller"
	"kstyle-hub/repository"
	"kstyle-hub/routes"
	"kstyle-hub/service"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Initialize DB
	config.InitDB()
	db := config.DB

	// Initialize Repositories
	customerRepo := repository.NewCustomerRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize Services
	customerService := service.NewCustomerService(customerRepo)
	orderService := service.NewOrderService(orderRepo)
	authService := service.NewAuthService(userRepo)

	// Initialize Controllers
	customerController := controller.NewCustomerController(customerService)
	orderController := controller.NewOrderController(orderService)
	authController := controller.NewAuthController(authService)

	// Set up routes
	routes.InitRoutes(e, customerController, orderController, authController)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
