package controllers

import (
	usecases "GoAir-Accounts/API/Places/application/useCases"
	"GoAir-Accounts/API/Places/domain"
	"GoAir-Accounts/API/Places/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePalceController struct {
	app *usecases.CreatePalce
}

func NewCreatePalceController() *CreatePalceController {
	postgres := infrastructure.GetPostgreSQL()
	app := usecases.NewCreatePlace(postgres)
	return &CreatePalceController{app: app}
}

func (cp_c *CreatePalceController) AddPlace(c *gin.Context) {
	var place domain.Place

	if err := c.ShouldBindJSON(&place); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	id, err := cp_c.app.Run(place)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo guardar el recurso " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "place",
			"id_place": id,
			"attributes": gin.H{
				"name": place.Name,
				"create_at": place.Timestamp,
			},
		},
	})

}