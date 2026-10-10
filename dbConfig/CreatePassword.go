package dbConfig

import (
	types "pasMan/types"
	security "pasMan/passwordSecurity"
)

func CreatePassword(serviceName string, password string) error {

	level := security.CheckPasswordStrength(password)

	newPassword := types.Password{
		ServiceName: serviceName,
		Password:    password,
		SecLevel: level,
	}

	result := DB.Create(&newPassword)
	return result.Error

}
