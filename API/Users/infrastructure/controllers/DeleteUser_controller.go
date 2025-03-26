package controllers

import (
	"GoAir-Accounts/API/Users/application/services"
	usecases "GoAir-Accounts/API/Users/application/useCases"
	"GoAir-Accounts/API/Users/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteUser struct {
	app *usecases.DeleteUser
	auth *services.Auth
}

func NewDeleteUser() *DeleteUser {
	postgres := infrastructure.GetPostgreSQL()
	app := usecases.NewDeleteUser(postgres)
	auth := infrastructure.GetBcrypt()
	auths := services.NewAuth(auth)
	return &DeleteUser{app: app, auth: auths}
}

func (du_c *DeleteUser) DeleteUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se proporcionó token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, err := du_c.auth.Run(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	rowsAffected, _ := du_c.app.Run(claims.Id_user)

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "No se pudo eliminar: No se entontró la referencia o ocurrió algo más",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Recurso eliminado",
	})
}