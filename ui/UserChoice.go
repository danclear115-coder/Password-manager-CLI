package ui

import (
	"fmt"
	functions "pasMan/functions"
	"pasMan/ui/theme"
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
		fmt.Println()
		fmt.Println(theme.Notice("Goodbye!"))
		fmt.Println()
	default:
		fmt.Println()
		fmt.Println(theme.Warning("Unrecognized input — pick a number from 1 to 5"))
	}

}
