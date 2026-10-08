package dbConfig

import (
	types "pasMan/types"
)

func UpdatePassword(serviceName string, password string, secLevel string, id uint) error {

	idCheckErr := IdCheck(id)

	if idCheckErr != nil {
		return idCheckErr
	}

	updateErr := DB.Model(&types.Password{}).Where("id = ?", id).Updates(map[string]interface{}{
		"service_name": serviceName,
		"password":  password,
		"sec_level": secLevel,
	})

	return updateErr.Error

}
