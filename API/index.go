package main

import (
	iUsers "GoAir-Accounts/API/Users/infrastructure"
	Broutes "GoAir-Accounts/API/Users/infrastructure/routes"

	iApplication "GoAir-Accounts/API/Applications/infrastructure"
	Aroutes "GoAir-Accounts/API/Applications/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	iUsers.GoDependences()
	iApplication.GoDependences()
	
	r := gin.Default()
	r.Use(cors.Default())

	Broutes.RegisterRouter(r)
	Aroutes.RegisterRouter(r)

	r.Run(":8010")
}