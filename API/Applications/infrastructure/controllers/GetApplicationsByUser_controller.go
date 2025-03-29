package controllers

import (
	"GoAir-Accounts/API/Applications/application/services"
	usecases "GoAir-Accounts/API/Applications/application/useCases"
	"GoAir-Accounts/API/Applications/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetApplicationByUserController struct {
	app *usecases.GetApplicationByUser
	auth *services.Auth
}

func NewGetApplicationByUser() *GetApplicationByUserController {
	postgres := infrastructure.GetPostgreSQL()
	jwt      := infrastructure.GetJWT()
	app := usecases.NewGetApplicationByUser(postgres)
	auth := services.NewAuth(jwt)
	return &GetApplicationByUserController{app: app, auth: auth}
} 

func (gabu_c *GetApplicationByUserController) GetApplicationByUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := gabu_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	apps := gabu_c.app.Run(claims.Id_user)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/admin/",
		},
		"data": apps,
	})
}

