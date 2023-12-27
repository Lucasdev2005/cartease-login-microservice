package main

import (
	"cartease-login-microservice/Controller"
	"cartease-login-microservice/Core"

	"github.com/gin-gonic/gin"
)

func main() {
	// if err := godotenv.Load(); err != nil{
	// }
	g := gin.Default()

	apiRoutes := []Core.Route{
		{
			RouteMethod: Core.RoutePost,
			Url:         "user/:id",
			Action:      Controller.Login,
		},
	}
	Core.RouteGroup("api", g, apiRoutes)
	g.Run(":3000")
}
