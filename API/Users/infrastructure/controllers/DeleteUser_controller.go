package controllers

import (
	usecases "GoAir-Accounts/API/Users/application/useCases"
	"GoAir-Accounts/API/Users/infrastructure"
	"GoAir-Accounts/API/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteUser struct {
	app *usecases.DeleteUser
}

func NewDeleteUser() *DeleteUser {
	postgres := infrastructure.GetPostgreSQL()
	app := usecases.NewDeleteUser(postgres)
	return &DeleteUser{app: app}
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

	claims, err := core.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		return
	}

	rowsAffected, _ := du_c.app.Run(int(claims.Id_user))

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