package controllers

import (
	"GoAir-Accounts/API/Users/application/services"
	usecases "GoAir-Accounts/API/Users/application/useCases"
	"GoAir-Accounts/API/Users/domain"
	"GoAir-Accounts/API/Users/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	app *usecases.CreateUser
	hashService *services.HashPassword
}

func NewCreateUserController() *CreateUserController {
	postgres := infrastructure.GetPostgreSQL()
	app := usecases.NewCreateUser(postgres)
	bcryp := infrastructure.GetBcrypt()
	hashService := services.NewHashPassword(bcryp)
	return &CreateUserController{app: app, hashService: hashService}
}

func (cuc *CreateUserController) AddUser(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	pass, err := cuc.hashService.Run(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Error al hashear la contraseña" + err.Error(),
		})
		return
	}

	user.Password = pass
	id, err := cuc.app.Run(user)
	
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
			"type": "user",
			"id_user": id,
			"attributes": gin.H{
				"first_name": user.First_name,
				"last_name": user.Last_name,
				"email": user.Email,
				"password": user.Password,
			},
		},
	})

}