package Route

import (
	HttpstatusSuccess "cartease-login-microservice/Core/HttpStatus/success"

	"github.com/gin-gonic/gin"
)

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
		GetHeader:     ctx.Request.Header.Get,
	}
}

func RoutePost(g *gin.Engine, url string, fn func(request Request) (interface{}, *Error)) {
	g.POST(url, func(ctx *gin.Context) {
		request := makeRequest(ctx)
		result, err := fn(request)
		if err == nil {
			ctx.JSON(HttpstatusSuccess.OK, result)
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
			ctx.JSON(HttpstatusSuccess.OK, result)
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
			ctx.JSON(HttpstatusSuccess.OK, result)
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
			ctx.JSON(HttpstatusSuccess.OK, result)
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
