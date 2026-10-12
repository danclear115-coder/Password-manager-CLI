package dbConfig

import (
	"fmt"
	"pasMan/auth"
)

const maxLoginAttempts = 3

func RequireAuth() bool {
	if !auth.IsConfigured() {
		return setupPassword()
	}
	return login()
}

func setupPassword() bool {
	fmt.Println("ПЕРВЫЙ ЗАПУСК — ЗАДАЙТЕ ПАРОЛЬ")
	fmt.Println("Этот пароль будет запрашиваться при каждом входе в приложение.")

	for {
		var pass string
		fmt.Print("Новый пароль: ")
		fmt.Scanln(&pass)

		if pass == "" {
			fmt.Println("Пароль не может быть пустым.")
			continue
		}

		var confirm string
		fmt.Print("Повторите пароль: ")
		fmt.Scanln(&confirm)

		if pass != confirm {
			fmt.Println("Пароли не совпадают, попробуйте ещё раз.")
			continue
		}

		if err := auth.SetPassword(pass); err != nil {
			fmt.Printf("Не удалось сохранить пароль: %v\n", err)
			return false
		}

		fmt.Println("Пароль сохранён. Запомните его — восстановить будет нельзя.")
		return true
	}
}

func login() bool {
	fmt.Println("АВТОРИЗАЦИЯ")

	for attempt := 1; attempt <= maxLoginAttempts; attempt++ {
		var pass string
		fmt.Print("Пароль: ")
		fmt.Scanln(&pass)

		ok, err := auth.Check(pass)
		if err != nil {
			fmt.Printf("Ошибка проверки пароля: %v\n", err)
			return false
		}
		if ok {
			return true
		}

		left := maxLoginAttempts - attempt
		if left > 0 {
			fmt.Printf("Неверный пароль. Попыток до удаления базы данных: %d.\n", left)
		}
	}

	fmt.Println("Слишком много неверных попыток. База была удалена.")
	DeleteDatabase()
	return false
}
