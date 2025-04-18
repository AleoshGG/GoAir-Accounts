package routes

import (
	"GoAir-Accounts/API/Users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controllers.NewCreateUserController().AddUser)
		userRoutes.DELETE("/", controllers.NewDeleteUser().DeleteUser)
		userRoutes.POST("/login", controllers.NewLoginUserController().Login)
		userRoutes.GET("/places", controllers.NewGetPlacesUserController().GetPlacesUser)
		userRoutes.GET("/token", controllers.NewValidateTokenController().ValidateToken)
	}
}