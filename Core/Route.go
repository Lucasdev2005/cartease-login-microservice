package Core

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Body   interface{}
	Params []interface{}
}

func makeRequest(ctx *gin.Context) *Request {

	var body interface{}
	body = ctx.BindJSON(&body)

	log.Println(ctx.Param("id"))

	return &Request{
		Body:   body,
		Params: nil,
	}
}

func RoutePost(g *gin.Engine, url string, fn func(request *Request) (interface{}, error)) {
	g.POST(url, func(ctx *gin.Context) {
		req := makeRequest(ctx)
		result, err := fn(req)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, result)
	})
}

func RoutePut(g *gin.Engine, url string, fn func(request *Request) (interface{}, error)) {
	g.PUT(url, func(ctx *gin.Context) {
		req := makeRequest(ctx)
		result, err := fn(req)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, result)
	})
}

func RoutePath(g *gin.Engine, fn func(request *Request) (interface{}, error)) {
	g.PUT("/", func(ctx *gin.Context) {
		req := makeRequest(ctx)
		result, err := fn(req)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, result)
	})
}

func RouteGet(g *gin.Engine, fn func(request *Request) (interface{}, error)) {
	g.PUT("/", func(ctx *gin.Context) {
		req := makeRequest(ctx)
		result, err := fn(req)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, result)
	})
}

func RouteDelete(g *gin.Engine, fn func(request *Request) (interface{}, error)) {
	g.PUT("/", func(ctx *gin.Context) {
		req := makeRequest(ctx)
		result, err := fn(req)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, result)
	})
}
