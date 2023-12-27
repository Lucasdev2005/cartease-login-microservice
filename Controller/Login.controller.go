package Controller

import (
	"cartease-login-microservice/Core"
	"cartease-login-microservice/Entity"
)

type LoginParams struct {
	USR_Username string
	USR_Password string
}

func Login(request Core.Request) (interface{}, *Core.Error) {
	user := Entity.User{}
	var err *Core.Error

	return user, err
}
