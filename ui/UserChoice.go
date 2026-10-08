package ui

import (
	"fmt"
	functions "pasMan/functions"
)

func UserChoice(choice string) {

	switch choice {
	case "1":
		functions.GetPasswords()
	case "2":
		functions.AddPassword()
	case "3":
		functions.UpdatePasswordUI()
	case "4":

	case "5":
		fmt.Println("Exit from programm")
	default:
		fmt.Println("Unrecognized input")
	}

}
