package main

import (
	iAdmin "GoAir-Accounts/API/Admin/infrastructure"
	Aroutes "GoAir-Accounts/API/Admin/infrastructure/routes"
	iUsers "GoAir-Accounts/API/Users/infrastructure"
	Broutes "GoAir-Accounts/API/Users/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	iAdmin.GoDependences()
	iUsers.GoDependences()
	
	r := gin.Default()
	r.Use(cors.Default())

	Aroutes.RegisterRouter(r)
	Broutes.RegisterRouter(r)
	
	r.Run(":8010")
}