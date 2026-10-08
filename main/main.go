package main

import (
	"fmt"
	ui "pasMan/ui"
	database "pasMan/dbConfig"
)

func main() {

	database.InitDB()

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
