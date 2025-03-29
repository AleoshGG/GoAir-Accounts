package controllers

import (
	"GoAir-Accounts/API/Users/infrastructure"
	"GoAir-Accounts/API/Users/application/services"
	usecases "GoAir-Accounts/API/Users/application/useCases"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetPlacesUserController struct {
	app *usecases.GetPlaces
	auth *services.Auth
}

func NewGetPlacesUserController() *GetPlacesUserController {
	postgres := infrastructure.GetPostgreSQL()
	jwt := infrastructure.GetBcrypt()
	app := usecases.NewGetPlaces(postgres)
	auth := services.NewAuth(jwt)
	return &GetPlacesUserController{app: app, auth: auth}
}

func (gpu_c *GetPlacesUserController) GetPlacesUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := gpu_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	places := gpu_c.app.Run(claims.Id_user)
	fmt.Print(places)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/users/",
		},
		"places": places,
	})
}