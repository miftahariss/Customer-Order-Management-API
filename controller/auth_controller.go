// controller/auth_controller.go
package controller

import (
	"kstyle-hub/service"
	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{Service: service}
}

func (a *AuthController) Register(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	err := a.Service.Register(username, password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "User registered successfully")
}

func (a *AuthController) Login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	user, err := a.Service.Login(username, password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "Invalid username or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Could not generate token")
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
