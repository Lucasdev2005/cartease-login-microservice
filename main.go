package main

import (
	"github.com/gin-gonic/gin"
	"cartease-login-microservice/controller"
)

func main() {
	g := gin.Default()

	var requestBody *controller.LoginParams

	g.POST("/", func(ctx *gin.Context) {
		if err := ctx.BindJSON(&requestBody); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	
		ctx.JSON(200, controller.Login(requestBody))
		// ctx.JSON(200, requestBody)
	})

	g.Run(":3000")
}