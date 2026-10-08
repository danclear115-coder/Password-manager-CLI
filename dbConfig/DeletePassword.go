package dbConfig

import (
	types "pasMan/types"
)

func DeletePassword(id uint) error {
	result := DB.Delete(&types.Password{}, id)
	return result.Error
}
