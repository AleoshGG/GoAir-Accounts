package routes

import (
	"GoAir-Accounts/API/Users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controllers.NewCreateUserController().AddUser)
		userRoutes.DELETE("/:id", controllers.NewDeleteUser().DeleteUser)
		userRoutes.POST("/login", controllers.NewLoginUserController().Login)
	}
}