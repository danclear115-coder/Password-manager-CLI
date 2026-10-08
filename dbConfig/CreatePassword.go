package dbConfig

import (
	types "pasMan/types"
)

func CreatePassword(serviceName string, password string) error {

	newPassword := types.Password{
		ServiceName: serviceName,
		Password:    password,
		SecLevel: "not defined",
	}

	result := DB.Create(&newPassword)
	return result.Error

}
