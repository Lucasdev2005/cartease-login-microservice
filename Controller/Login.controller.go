package Controller

import (
	"cartease-login-microservice/Core"
	"cartease-login-microservice/Entity"
	"log"
)

type LoginParams struct {
	USR_Username string
	USR_Password string
}

func Login(request *Core.Request) (interface{}, error) {

	log.Println(request.Body["USR_Username"])
	user := &Entity.User{
		USR_Username:   request.Body["USR_Username"].(string),
		USR_Complement: request.Params["id"].(string),
	}
	return user, nil
}
