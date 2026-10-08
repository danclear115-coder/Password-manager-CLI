package functions

import (
	"fmt"
	database "pasMan/dbConfig"
)

func AddPassword() {

	serviceName := ""
	fmt.Println("Enter service name")
	fmt.Scanln(&serviceName)

	password := ""
	fmt.Println("Enter password")
	fmt.Scanln(&password)

	err := database.CreatePassword(serviceName, password)

	if err != nil {
		fmt.Println("Database error")
	}

}
