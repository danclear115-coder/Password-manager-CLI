package dbConfig

import (
	types "pasMan/types"

)

func GetPasswords() ([]types.Password, error) {

	var passwords []types.Password
	result := DB.Find(&passwords)
	return passwords, result.Error	

}
