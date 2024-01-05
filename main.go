package main

import (
	"cartease-login-microservice/Controller"
	Route "cartease-login-microservice/Core/Route"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error on load enviroment")
	}

	g := gin.Default()

	apiRoutes := []Route.Route{
		{
			RouteMethod: Route.RoutePost,
			Url:         "user/:id",
			Action:      Controller.Login,
		},
	}
	Route.RouteGroup("api", g, apiRoutes, nil)
	g.Run(":" + os.Getenv("PORT"))
}
