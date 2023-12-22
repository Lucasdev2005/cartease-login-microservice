package main

import (
	"cartease-login-microservice/Controller"
	"cartease-login-microservice/Core"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	Core.RoutePost(g, "user/:id", Controller.Login)
	Core.RouteGet(g, "createAccount", Controller.CreateAccount)

	g.Run(":3000")
	defer Core.DatabaseConnection.Close()
}
