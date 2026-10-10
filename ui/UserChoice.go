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
		functions.UpdatePassword()
	case "4":
		functions.DeletePassword()
	case "5":
		fmt.Println("Exit from programm")
	default:
		fmt.Println("Unrecognized input")
	}

}
