package Core

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Body   map[string]interface{}
	Params map[string]interface{}
}

func makeRequest(ctx *gin.Context) *Request {
	var bodyJSON map[string]interface{}
	params := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&bodyJSON); err != nil {
		log.Println("Erro ao vincular JSON:", err.Error())
	}

	for _, param := range ctx.Params {
		if param.Key != "" {
			params[param.Key] = param.Value
		}
	}

	return &Request{
		Body:   bodyJSON,
		Params: params,
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
