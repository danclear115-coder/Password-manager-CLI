package dbConfig

import (
	customErrors "pasMan/customErrors"
	types "pasMan/types"
)

func IdCheck(id uint) error {

	var currentPassword types.Password

	err := DB.Where("id = ?", id).First(&currentPassword).Error

	if err != nil {
		return customErrors.NoIdErr
	}

	return err

}
