package controllers

import (
	usecases "GoAir-Accounts/API/Users/application/useCases"
	"GoAir-Accounts/API/Users/infrastructure"
	"net/http"
	"strconv"

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
	id := c.Param("id")
	id_user, _ := strconv.ParseInt(id, 10, 64)

	rowsAffected, _ := du_c.app.Run(int(id_user))

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