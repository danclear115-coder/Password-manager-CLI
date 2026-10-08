package functions

import (
	"fmt"
	"os"
	database "pasMan/dbConfig"
	types "pasMan/types"
	"strings"

	"golang.org/x/term"
)

const (
	colID    = 4
	colName  = 28
	colPass  = 28
	colLevel = 18
)

func outf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	fmt.Print(strings.ReplaceAll(s, "\n", "\r\n"))
}

func outln(a ...any) {
	s := strings.TrimSuffix(fmt.Sprintln(a...), "\n")
	fmt.Print(s + "\r\n")
}

func UpdatePassword() {
	passwords, err := database.GetPasswords()
	if err != nil {
		outln("Database error:", err)
		return
	}

	if len(passwords) == 0 {
		outln("There are no passwords yet.")
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

		outln("PASSWORD CHANGE")
		outln("↑ ↓ — select password")
		outln("Enter — update")
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
			editPassword(&passwords[selected])
			continue
		}

		if n == 1 && key[0] == 27 {
			clearScreen()
			return
		}
	}
}

func fit(s string, w int) string {
	r := []rune(s)
	if len(r) > w {
		r = append(r[:w-1], '…')
	}
	return string(r) + strings.Repeat(" ", w-len(r))
}

func tableRow(id, name, pass, level string) string {
	return fmt.Sprintf("| %s | %s | %s | %s |",
		fit(id, colID),
		fit(name, colName),
		fit(pass, colPass),
		fit(level, colLevel),
	)
}

func tableSeparator() string {
	return fmt.Sprintf("+%s+%s+%s+%s+",
		strings.Repeat("-", colID+2),
		strings.Repeat("-", colName+2),
		strings.Repeat("-", colPass+2),
		strings.Repeat("-", colLevel+2),
	)
}

func printPasswordList(passwords []types.Password, selected int) {
	sep := tableSeparator()

	outln(sep)
	outln(tableRow("ID", "SERVICE NAME", "PASSWORD", "SECURITY LEVEL"))
	outln(sep)

	for i, p := range passwords {
		row := tableRow(
			fmt.Sprint(p.Id),
			p.ServiceName,
			p.Password,
			p.SecLevel,
		)

		if i == selected {
			outln("\033[7m" + row + "\033[0m")
		} else {
			outln(row)
		}

		outln(sep)
	}
}

func editPassword(password *types.Password) {
	serviceName := password.ServiceName
	passwordValue := password.Password
	secLevel := password.SecLevel

	field := 0

	for {

		clearScreen()

		outln("Update password")
		outln()
		outln("Tab — next field")
		outln("Enter - save")
		outln("Esc — cancel")
		outln()

		printField("Service Name", serviceName, field == 0)
		printField("Password", passwordValue, field == 1)
		printField("Security Level", secLevel, field == 2)

		key := make([]byte, 3)
		n, err := os.Stdin.Read(key)
		if err != nil {
			return
		}

		if n == 1 && key[0] == 9 {
			field++
			if field > 2 {
				field = 0
			}
			continue
		}

		if n == 1 && key[0] == 13 {
			err := database.UpdatePassword(
				serviceName,
				passwordValue,
				secLevel,
				password.Id,
			)
			if err != nil {
				clearScreen()
				outln("Update error:", err)
				waitForKey()
				return
			}

			password.ServiceName = serviceName
			password.Password = passwordValue
			password.SecLevel = secLevel

			clearScreen()
			outln("Press any key...")
			waitForKey()
			return
		}

		if n == 1 && key[0] == 27 {
			return
		}

		if n == 1 && (key[0] == 8 || key[0] == 127) {
			switch field {
			case 0:
				serviceName = trimLastRune(serviceName)
			case 1:
				passwordValue = trimLastRune(passwordValue)
			case 2:
				secLevel = trimLastRune(secLevel)
			}
			continue
		}

		if n == 1 && key[0] >= 32 && key[0] <= 126 {
			switch field {
			case 0:
				serviceName += string(key[0])
			case 1:
				passwordValue += string(key[0])
			case 2:
				secLevel += string(key[0])
			}
		}
	}
}

func trimLastRune(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return s
	}
	return string(r[:len(r)-1])
}

func printField(name string, value string, selected bool) {
	line := fmt.Sprintf("%-16s: %s", name, value)
	if selected {
		line = "\033[7m" + line + "\033[0m"
	}
	outln(line)
}

func clearScreen() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func waitForKey() {
	outln()
	outln("Press any key...")
	key := make([]byte, 1)
	os.Stdin.Read(key)
}
