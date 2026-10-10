package functions

import (
	"os"

	database "pasMan/dbConfig"
	types "pasMan/types"

	"golang.org/x/term"
)

func DeletePassword() {
	passwords, err := database.GetPasswords()
	if err != nil {
		outln("Database error:", err)
		return
	}

	if len(passwords) == 0 {
		outln("No passwords yet.")
		return
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		outln("Terminal error:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	selected := 0

	for {
		clearScreen()

		outln("DELETE PASSWORD")
		outln("↑ ↓ — select password")
		outln("Enter — delete")
		outln("Esc — exit")
		outln()

		printPasswordList(passwords, selected)

		key := make([]byte, 3)
		n, err := os.Stdin.Read(key)
		if err != nil {
			return
		}

		if n == 3 && key[0] == 27 && key[1] == '[' && key[2] == 'A' {
			if selected > 0 {
				selected--
			}
			continue
		}

		if n == 3 && key[0] == 27 && key[1] == '[' && key[2] == 'B' {
			if selected < len(passwords)-1 {
				selected++
			}
			continue
		}

		if n == 1 && key[0] == 13 {
			p := passwords[selected]

			if !confirmDelete(p) {
				continue
			}

			if err := database.DeletePassword(p.Id); err != nil {
				clearScreen()
				outln("Delete error:", err)
				waitForKey()
				continue
			}

			passwords = append(passwords[:selected], passwords[selected+1:]...)

			clearScreen()
			outln("Password deleted successfully!")
			
			if len(passwords) == 0 {
				clearScreen()
				outln("No passwords left.")
				return
			}

			if selected >= len(passwords) {
				selected = len(passwords) - 1
			}
			continue
		}

		if n == 1 && key[0] == 27 {
			clearScreen()
			return
		}
	}
}

func confirmDelete(p types.Password) bool {
	clearScreen()

	outln("CONFIRM DELETION")
	outln()
	outf("Delete password for \"%s\" (ID %d)?\n", p.ServiceName, p.Id)
	outln()
	outln("y — yes, delete")
	outln("any other key — cancel")

	key := make([]byte, 3)
	n, err := os.Stdin.Read(key)
	if err != nil {
		return false
	}

	return n == 1 && (key[0] == 'y' || key[0] == 'Y')
}
