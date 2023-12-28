package Route

import "github.com/gin-gonic/gin"

type Route struct {
	RouteMethod func(g *gin.Engine, url string, fn func(request Request) (interface{}, *Error))
	Url         string
	Action      func(request Request) (interface{}, *Error)
}
