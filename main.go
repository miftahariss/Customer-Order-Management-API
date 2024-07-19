package main

import (
	"kstyle-hub/config"
	"kstyle-hub/controller"
	"kstyle-hub/repository"
	"kstyle-hub/routes"
	"kstyle-hub/service"

	_ "kstyle-hub/docs" // Import generated docs

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

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

	// Swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
