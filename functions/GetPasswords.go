package functions

import (
	"fmt"
	database "pasMan/dbConfig"
)

func GetPasswords() {

	passwords, err := database.GetPasswords()
	if err != nil {
		fmt.Println("Database error:", err)
		return
	}

	fmt.Println()
	fmt.Println("+------+------------------------------+------------------------------+--------------------+")
	fmt.Println("| ID   | SERVICE NAME                 | PASSWORD                     | SECURITY LEVEL     |")
	fmt.Println("+------+------------------------------+------------------------------+--------------------+")

	for _, password := range passwords {
		fmt.Printf("| %-4d | %-28s | %-28s | %-18s |\n",
			password.Id,
			password.ServiceName,
			password.Password,
			password.SecLevel,
		)

		fmt.Println("+------+------------------------------+------------------------------+--------------------+")
	}

	fmt.Println()
}