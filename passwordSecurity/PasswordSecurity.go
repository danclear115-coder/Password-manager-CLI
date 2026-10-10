package security

import (
	"strings"
	"unicode"
)

var commonPasswords = map[string]bool{
	"password": true, "123456": true, "12345678": true, "qwerty": true,
	"abc123":   true, "letmein":  true, "admin":    true, "welcome":  true,
	"iloveyou": true, "monkey":  true, "dragon":   true, "111111":   true,
	"123456789": true, "qwerty123": true, "passw0rd": true,
}

func CheckPasswordStrength(password string) string {
	if password == "" {
		return "Very Weak"
	}

	runes := []rune(password)
	length := len(runes)

	var hasLower, hasUpper, hasDigit, hasSymbol bool
	for _, r := range runes {
		switch {
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsDigit(r):
			hasDigit = true
		default:
			hasSymbol = true
		}
	}

	lowered := strings.ToLower(password)
	stripped := strings.TrimRight(lowered, "0123456789!@#$%^&*._-")
	if commonPasswords[lowered] || commonPasswords[stripped] {
		return "Very Weak"
	}

	score := 0
	switch {
	case length >= 16:
		score += 3
	case length >= 12:
		score += 2
	case length >= 8:
		score += 1
	}

	classes := 0
	for _, ok := range []bool{hasLower, hasUpper, hasDigit, hasSymbol} {
		if ok {
			classes++
		}
	}
	score += classes - 1

	if hasRepeatedChars(runes) {
		score--
	}

	if hasSequence([]rune(lowered)) {
		score--
	}

	var level string
	switch {
	case score <= 1:
		level = "Very Weak"
	case score == 2:
		level = "Weak"
	case score <= 4:
		level = "Medium"
	case score == 5:
		level = "Strong"
	default:
		level = "Very Strong"
	}

	return level
}

func hasRepeatedChars(r []rune) bool {
	for i := 0; i+2 < len(r); i++ {
		if r[i] == r[i+1] && r[i+1] == r[i+2] {
			return true
		}
	}
	return false
}

func hasSequence(r []rune) bool {
	isAlnum := func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsDigit(c)
	}
	for i := 0; i+2 < len(r); i++ {
		a, b, c := r[i], r[i+1], r[i+2]
		if !isAlnum(a) || !isAlnum(b) || !isAlnum(c) {
			continue
		}
		if (b == a+1 && c == b+1) || (b == a-1 && c == b-1) {
			return true
		}
	}
	return false
}
