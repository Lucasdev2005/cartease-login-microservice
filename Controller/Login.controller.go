package Controller

import (
	"cartease-login-microservice/Core/Route"
	"cartease-login-microservice/Entity"
)

type LoginParams struct {
	USR_Username string
	USR_Password string
}

func Login(request Route.Request) (interface{}, *Route.Error) {
	user := Entity.User{}
	var err *Route.Error

	return user, err
}
