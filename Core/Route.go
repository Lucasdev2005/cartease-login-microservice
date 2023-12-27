package Core

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
	Body          interface{}
	GetParam      func(key string) string
	GetQueryParam func(key string) string
}

type Error struct {
	ErrorCode int
	Message   string
}

type Route struct {
	RouteMethod func(g *gin.Engine, url string, fn func(request Request) (interface{}, *Error))
	Url         string
	Action      func(request Request) (interface{}, *Error)
}

func makeRequest(ctx *gin.Context) Request {

	var body map[string]interface{}
	ctx.BindJSON(&body)

	params := make(map[string]interface{})

	for _, param := range ctx.Params {
		params[param.Key] = param.Value
	}

	return Request{
		Body:          body,
		GetParam:      ctx.Param,
		GetQueryParam: ctx.Query,
	}
}

func RoutePost(g *gin.Engine, url string, fn func(request Request) (interface{}, *Error)) {
	g.POST(url, func(ctx *gin.Context) {
		request := makeRequest(ctx)
		result, err := fn(request)
		if err == nil {
			ctx.JSON(200, result)
			return
		}

		ctx.JSON(err.ErrorCode, err.Message)
	})
}

func RoutePut(g *gin.Engine, url string, fn func(request Request) (interface{}, *Error)) {
	g.PUT(url, func(ctx *gin.Context) {
		request := makeRequest(ctx)
		result, err := fn(request)
		if err == nil {
			ctx.JSON(200, result)
			return
		}

		ctx.JSON(err.ErrorCode, err.Message)
	})
}

func RoutePath(g *gin.Engine, url string, fn func(request Request) (interface{}, *Error)) {
	g.PATCH(url, func(ctx *gin.Context) {
		request := makeRequest(ctx)
		result, err := fn(request)
		if err == nil {
			ctx.JSON(200, result)
			return
		}

		ctx.JSON(err.ErrorCode, err.Message)
	})
}

func RouteGet(g *gin.Engine, url string, fn func(request Request) (interface{}, *Error)) {
	g.GET(url, func(ctx *gin.Context) {
		request := makeRequest(ctx)
		result, err := fn(request)
		if err == nil {
			ctx.JSON(200, result)
			return
		}

		ctx.JSON(err.ErrorCode, err.Message)
	})
}

func RouteDelete(g *gin.Engine, url string, fn func(request Request) (interface{}, *Error)) {
	g.DELETE(url, func(ctx *gin.Context) {
		request := makeRequest(ctx)
		result, err := fn(request)
		if err == nil {
			ctx.JSON(200, result)
			return
		}

		ctx.JSON(err.ErrorCode, err.Message)
	})
}

func RouteGroup(
	prefix string,
	g *gin.Engine,
	routes []Route) {
	for _, route := range routes {
		route.RouteMethod(g, prefix+"/"+route.Url, route.Action)
	}
}
