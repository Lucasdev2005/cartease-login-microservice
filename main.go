package main

import (
	"cartease-login-microservice/Controller"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	var requestBody *Controller.LoginParams

	g.POST("/", func(ctx *gin.Context) {
		if err := ctx.BindJSON(&requestBody); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, Controller.Login(requestBody))
	})

	g.Run(":3000")
}
