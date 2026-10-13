package ui

import (
	"fmt"

	"pasMan/ui/theme"
)

func PrintMenu() {
	body := []string{
		"",
		theme.Key("1") + "  Show passwords",
		theme.Key("2") + "  Create password",
		theme.Key("3") + "  Update password",
		theme.Key("4") + "  Delete passwords",
		"",
		theme.Divider(theme.MenuWidth - 4),
		theme.Key("5") + "  " + theme.Dim("Exit"),
	}

	fmt.Println()
	theme.Print(theme.Box("pasMan", body, theme.MenuWidth))
	fmt.Println(theme.Hints("1-5", "choose an option"))
	fmt.Println()
	fmt.Print(theme.Prompt(""))
}
