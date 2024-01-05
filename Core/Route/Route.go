package Route

import (
	HttpClientError "cartease-login-microservice/Core/HttpStatus/clientError"
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

func RoutePost(
	g *gin.Engine,
	url string,
	fn func(request Request) (interface{}, *Error),
	middleWare func(req Request) bool,
) {
	g.POST(url, func(ctx *gin.Context) {
		processRequest(ctx, middleWare, fn)
	})
}

func RoutePut(
	g *gin.Engine, url string,
	fn func(request Request) (interface{}, *Error),
	middleWare func(req Request) bool,
) {
	g.PUT(url, func(ctx *gin.Context) {
		processRequest(ctx, middleWare, fn)
	})
}

func RoutePath(
	g *gin.Engine, url string,
	fn func(request Request) (interface{}, *Error),
	middleWare func(req Request) bool,
) {
	g.PATCH(url, func(ctx *gin.Context) {
		processRequest(ctx, middleWare, fn)
	})
}

func RouteGet(
	g *gin.Engine,
	url string,
	fn func(request Request) (interface{}, *Error),
	middleWare func(req Request) bool,
) {
	g.GET(url, func(ctx *gin.Context) {
		processRequest(ctx, middleWare, fn)
	})
}

func RouteDelete(
	g *gin.Engine,
	url string,
	fn func(request Request) (interface{}, *Error),
	middleWare func(req Request) bool,
) {
	g.DELETE(url, func(ctx *gin.Context) {
		processRequest(ctx, middleWare, fn)
	})
}

func RouteGroup(
	prefix string,
	g *gin.Engine,
	routes []Route,
	middleWare func(req Request) bool,
) {
	for _, route := range routes {
		route.RouteMethod(g, prefix+"/"+route.Url, route.Action, middleWare)
	}
}

func verifyMiddleware(req Request, mw func(req Request) bool) (bool, Error) {

	if mw != nil {
		success := mw(req)
		return success, Error{
			ErrorCode: HttpClientError.Unauthorized,
			Message:   HttpClientError.GetStatusName(HttpClientError.Unauthorized),
		}
	}

	return true, Error{}
}

func processRequest(
	context *gin.Context,
	middleWare func(req Request) bool,
	fn func(request Request) (interface{}, *Error),
) {
	request := makeRequest(context)

	if success, err := verifyMiddleware(request, middleWare); success {
		if !success {
			context.JSON(err.ErrorCode, err.Message)
			return
		}
	}

	result, err := fn(request)

	if err == nil {
		context.JSON(HttpstatusSuccess.OK, result)
	} else {
		context.JSON(err.ErrorCode, err.Message)
	}
}
