package controllers

import (
	"GoAir-Accounts/API/Applications/application/services"
	usecases "GoAir-Accounts/API/Applications/application/useCases"
	"GoAir-Accounts/API/Applications/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateApplicationController struct {
	app *usecases.CreateApplication
	auth *services.Auth
	broker *services.SendRequestPlace
}

func NewCreCreateApplicationController() *CreateApplicationController {
	postgres := infrastructure.GetPostgreSQL()
	jwt      := infrastructure.GetJWT()
	rabbitmq := infrastructure.GetRabbitMQ()
	app 	 := usecases.NewCreateApplication(postgres)
	auth     := services.NewAuth(jwt)
	broker   := services.NewSendRequestPlace(rabbitmq)
	return &CreateApplicationController{app: app, auth: auth, broker: broker}
}

func (ca_c *CreateApplicationController) CreateApplication(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := ca_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	msg, err := ca_c.app.Run(claims.Id_user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo obtener de manera correcta el mesaje"})
		return
	}

	ca_c.broker.Run(msg)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/admin/",
		},
		"status_application": msg.Status_application,
	})
}