package controllers

import (
	"GoAir-Accounts/API/Users/application/services"
	usecases "GoAir-Accounts/API/Users/application/useCases"
	"GoAir-Accounts/API/Users/infrastructure"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	app *usecases.GetUserByEmail
	hashService *services.ValidatePassword
	jwtService *services.CreateJWT
}

func NewLoginUserController() *LoginUserController {
	postgres := infrastructure.GetPostgreSQL()
	bcrypt := infrastructure.GetBcrypt()
	app := usecases.NewGetUserByEmail(postgres)
	hashService := services.NewValidatePassword(bcrypt) 
	jwt := services.NewCreateJWT(bcrypt)
	return &LoginUserController{app: app, hashService: hashService, jwtService: jwt}
}

func (l_c *LoginUserController) Login(c *gin.Context) {
	var credentials struct {
		Email string
		Password string
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}
	fmt.Println(credentials)
	user := l_c.app.Run(credentials.Email)
	if !l_c.hashService.Run(credentials.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Contraseña incorrecta: ",
		})
		return
	}
	
	fmt.Println(user)

	token, err := l_c.jwtService.Run(user.Id_user, user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Error al generar el JWT: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/users/",
		},
		"token": token,
	})
		
}