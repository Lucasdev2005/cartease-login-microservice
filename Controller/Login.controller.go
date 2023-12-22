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
	user := &Entity.User{
		USR_Username:   request.Body["USR_Username"].(string),
		USR_Complement: request.Params["id"].(string),
	}
	return user, nil
}

func CreateAccount(request *Core.Request) (interface{}, error) {
	return nil, nil
}
