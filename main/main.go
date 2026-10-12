package main

import (
	"fmt"
	database "pasMan/dbConfig"
	auth "pasMan/dbConfig"
	ui "pasMan/ui"
)

func main() {

	database.InitDB()

	if !auth.RequireAuth() {
		return
	}

	for {

		ui.PrintMenu()

		choice := ""
		fmt.Scanln(&choice)

		ui.UserChoice(choice)
		
		if choice == "5" {
			break
		}

	}

}
