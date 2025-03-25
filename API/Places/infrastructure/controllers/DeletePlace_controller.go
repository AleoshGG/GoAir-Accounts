package controllers

import (
	usecases "GoAir-Accounts/API/Places/application/useCases"
	"GoAir-Accounts/API/Places/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeletePlace struct {
	app *usecases.DeletePlace
}

func NewDeletePlace() *DeletePlace {
	postgres := infrastructure.GetPostgreSQL()
	app := usecases.NewDeletePlace(postgres)
	return &DeletePlace{app: app}
}

func (dp_c *DeletePlace) DeletePlace(c *gin.Context) {
	id := c.Param("id")
	id_place, _ := strconv.ParseInt(id, 10, 64)

	rowsAffected, _ := dp_c.app.Run(int(id_place))

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