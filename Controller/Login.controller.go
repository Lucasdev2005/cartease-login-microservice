package controller

import (
	"cartease-login-microservice/entity"
	"log"
)

type LoginParams struct {
	USR_Username string
	USR_Password string
}

func Login(loginParams *LoginParams) *entity.User {

	log.Println(loginParams)


	user := &entity.User{
        USR_Username:      loginParams.USR_Username,
        USR_Password:      loginParams.USR_Password,
        USR_Neighborhood:  "neighborhood",
        USR_CEP:           "cep",
        USR_City:          "city",
        USR_UF:            "uf",
        USR_Complement:    "complement",
    }

	return user
}
