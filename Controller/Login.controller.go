package controller

import "cartease-login-microservice/entity"

func Login() *entity.User {
	user := &entity.User{
        USR_Username:      "username",
        USR_Password:      "password",
        USR_Neighborhood:  "neighborhood",
        USR_CEP:           "cep",
        USR_City:          "city",
        USR_UF:            "uf",
        USR_Complement:    "complement",
    }

	return user
}
