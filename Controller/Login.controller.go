package Controller

import (
	"cartease-login-microservice/Core"
	"cartease-login-microservice/Entity"
)

type LoginParams struct {
	USR_Username string
	USR_Password string
}

func Login(request *Core.Request) (interface{}, error) {

	user := &Entity.User{}
	return user, nil
}
