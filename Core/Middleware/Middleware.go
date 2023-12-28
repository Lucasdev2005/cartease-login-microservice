package middleware

import (
	HttpClientError "cartease-login-microservice/Core/HttpStatus/clientError"
	"cartease-login-microservice/Core/Route"
)

func VerifyMiddleware(req Route.Request, mw func(Route.Request) bool) (bool, Route.Error) {
	success := mw(req)
	return success, Route.Error{
		ErrorCode: HttpClientError.Unauthorized,
		Message:   HttpClientError.GetStatusName(HttpClientError.Unauthorized),
	}
}
