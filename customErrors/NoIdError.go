package customErrors

import (
	"errors"
)

var NoIdErr = errors.New("Task with this ID not found.")
