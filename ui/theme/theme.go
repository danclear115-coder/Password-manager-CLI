// Package theme contains all terminal styling for pasMan: colours, frames,
// the password table, key hints and prompts.
//
// It depends only on the standard library, so any package can import it
// without creating an import cycle. Every function returns plain strings
// (no "\n" inside), which makes the output safe to print both in normal
// mode (fmt.Println) and in raw mode (where "\r\n" is required).
package theme

import (
	"fmt"
	"os"
	"strings"
)

// Layout (visible widths, without the 2-space left margin).
const (
	MenuWidth = 46
	FormWidth = 68

	colID    = 4
	colName  = 28
	colPass  = 28
	colLevel = 18

	// TableWidth is the visible width of the password table.
	TableWidth = 5 + (colID + 2) + (colName + 2) + (colPass + 2) + (colLevel + 2)
)

// Palette (xterm-256).
const (
	reset = "\033[0m"

	cAccent = "38;5;78"  // green
	cBorder = "38;5;240" // dark grey
	cDim    = "38;5;245" // grey
	cText   = "38;5;252" // near white
	cOk     = "38;5;78"
	cWarn   = "38;5;179" // amber
	cErr    = "38;5;203" // soft red

	bgSelected = "48;5;236"
	bgKey      = "48;5;238"
)

// colorOn is false when stdout is not a terminal, when NO_COLOR is set or
// when TERM=dumb, so redirected output stays clean.
var colorOn = detectColor()

func detectColor() bool {
	if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		return false
	}
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func paint(code, s string) string {
	if !colorOn || s == "" {
		return s
	}
	return "\033[" + code + "m" + s + reset
}

// Accent returns s in bold accent colour.
func Accent(s string) string { return paint("1;"+cAccent, s) }

// Bold returns s in bright bold text.
func Bold(s string) string { return paint("1;"+cText, s) }

// Dim returns s in muted grey.
func Dim(s string) string { return paint(cDim, s) }

// Caution returns s in amber.
func Caution(s string) string { return paint(cWarn, s) }

// ---------------------------------------------------------------- text utils

// visibleWidth is the on-screen width of s, ignoring ANSI escape sequences.
func visibleWidth(s string) int {
	n := 0
	inEsc := false
	for _, r := range s {
		switch {
		case inEsc:
			if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
				inEsc = false
			}
		case r == '\033':
			inEsc = true
		default:
			n++
		}
	}
	return n
}

func padRight(s string, w int) string {
	if d := w - visibleWidth(s); d > 0 {
		return s + strings.Repeat(" ", d)
	}
	return s
}

// Clip shortens s to at most w runes, ending with "…" when it was cut.
func Clip(s string, w int) string {
	if w <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) > w {
		r = append(r[:w-1], '…')
	}
	return string(r)
}

// Fit clips s to w runes and pads it with spaces to exactly w.
func Fit(s string, w int) string {
	c := Clip(s, w)
	return c + strings.Repeat(" ", w-len([]rune(c)))
}

// Print prints lines one per row (normal terminal mode).
func Print(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
}

// ------------------------------------------------------------------- pieces

// Logo is the start-up banner.
func Logo() []string {
	art := [3]string{
		"┌─┐┌─┐┌─┐╔╦╗┌─┐┌┐┌",
		"├─┘├─┤└─┐║║║├─┤│││",
		"┴  ┴ ┴└─┘╩ ╩┴ ┴┘└┘",
	}
	lines := []string{""}
	for i, a := range art {
		l := "  " + Accent(a)
		if i == 1 {
			l += "   " + Dim("terminal password manager")
		}
		lines = append(lines, l)
	}
	return append(lines, "")
}

// Title is a breadcrumb ("pasMan ❯ Section ❯ Sub") followed by a thin rule.
func Title(width int, parts ...string) []string {
	crumb := "  " + Accent("pasMan")
	for i, p := range parts {
		crumb += Dim(" ❯ ")
		if i == len(parts)-1 {
			crumb += Bold(p)
		} else {
			crumb += p
		}
	}
	return []string{crumb, "  " + paint(cBorder, strings.Repeat("─", width))}
}

// Divider is a thin horizontal line of w cells.
func Divider(w int) string { return paint(cBorder, strings.Repeat("─", w)) }

// frame draws a rounded box. Content width is width-4 (border + 1 space each side).
func frame(title string, body []string, width int, border string) []string {
	inner := width - 2
	b := func(s string) string { return paint(border, s) }

	var top string
	if title != "" {
		t := " " + title + " "
		fill := inner - 1 - visibleWidth(t)
		if fill < 0 {
			fill = 0
		}
		top = b("╭─") + paint("1;"+cAccent, t) + b(strings.Repeat("─", fill)+"╮")
	} else {
		top = b("╭" + strings.Repeat("─", inner) + "╮")
	}

	out := []string{"  " + top}
	for _, l := range body {
		out = append(out, "  "+b("│")+" "+padRight(l, inner-2)+" "+b("│"))
	}
	return append(out, "  "+b("╰"+strings.Repeat("─", inner)+"╯"))
}

// Box draws a rounded box with an optional title in the top border.
func Box(title string, body []string, width int) []string {
	return frame(title, body, width, cBorder)
}

// Key renders a "keycap", e.g. [ 1 ] or [ Enter ].
func Key(k string) string {
	if !colorOn {
		return "[" + k + "]"
	}
	return "\033[" + bgKey + ";1;" + cText + "m " + k + " " + reset
}

// Hints renders a footer of key hints: Hints("Enter", "save", "Esc", "back").
func Hints(pairs ...string) string {
	var parts []string
	for i := 0; i+1 < len(pairs); i += 2 {
		parts = append(parts, Key(pairs[i])+" "+Dim(pairs[i+1]))
	}
	return "  " + strings.Join(parts, "   ")
}

// ----------------------------------------------------------------- messages

func Success(msg string) string { return "  " + paint("1;"+cOk, "✔") + " " + msg }
func Failure(msg string) string { return "  " + paint("1;"+cErr, "✖") + " " + paint(cErr, msg) }
func Warning(msg string) string { return "  " + paint("1;"+cWarn, "▲") + " " + paint(cWarn, msg) }
func Notice(msg string) string  { return "  " + paint("1;"+cAccent, "◆") + " " + msg }

// Prompt is the input prompt, printed without a trailing newline.
func Prompt(label string) string {
	if label == "" {
		return "  " + Accent("❯") + " "
	}
	return "  " + Accent("❯") + " " + label + ": "
}

// Field renders a labelled input box for forms. The active field is
// highlighted and shows a cursor; long values show their tail.
func Field(label, value string, active bool, width int) []string {
	capacity := width - 5 // frame + padding = 4, cursor = 1
	r := []rune(value)
	if len(r) > capacity {
		r = append([]rune{'…'}, r[len(r)-capacity+1:]...)
	}
	text := string(r)

	border := cBorder
	marker := "  "
	name := Dim(label)
	if active {
		border = cAccent
		marker = paint("1;"+cAccent, "❯ ")
		name = Bold(label)
		text += paint(cAccent, "█")
	}

	out := []string{"  " + marker + name}
	return append(out, frame("", []string{text}, width, border)...)
}

// -------------------------------------------------------------------- table

// Row is one table row; all values are plain strings.
type Row struct {
	ID, Service, Password, Level string
}

type levelStyle struct {
	filled int
	color  string
}

var levels = map[string]levelStyle{
	"very weak":   {1, "38;5;203"},
	"weak":        {2, "38;5;208"},
	"medium":      {3, "38;5;179"},
	"strong":      {4, "38;5;149"},
	"very strong": {5, "38;5;78"},
}

func levelCell(level string, w int) string {
	st, ok := levels[strings.ToLower(strings.TrimSpace(level))]
	if !ok {
		return Dim(Fit(level, w))
	}
	meter := strings.Repeat("▰", st.filled) + strings.Repeat("▱", 5-st.filled)
	return paint(st.color, meter+" "+Fit(level, w-6))
}

// withBg keeps a background colour alive across inner resets.
func withBg(s string) string {
	if !colorOn {
		return s
	}
	bg := "\033[" + bgSelected + "m"
	return bg + strings.ReplaceAll(s, reset, reset+bg) + reset
}

// Table renders the password table. selected is the highlighted row index
// (use -1 for none). The selected row also gets a "❯" marker in the gutter,
// so it is visible even without colours.
func Table(rows []Row, selected int) []string {
	cols := [4]int{colID, colName, colPass, colLevel}

	border := func(l, m, r string) string {
		segs := make([]string, len(cols))
		for i, w := range cols {
			segs[i] = strings.Repeat("─", w+2)
		}
		return "  " + paint(cBorder, l+strings.Join(segs, m)+r)
	}
	bar := paint(cBorder, "│")
	line := func(cells [4]string) string {
		var sb strings.Builder
		sb.WriteString(bar)
		for _, c := range cells {
			sb.WriteString(" " + c + " " + bar)
		}
		return sb.String()
	}

	header := [4]string{
		Accent(Fit("ID", colID)),
		Accent(Fit("SERVICE NAME", colName)),
		Accent(Fit("PASSWORD", colPass)),
		Accent(Fit("SECURITY LEVEL", colLevel)),
	}

	out := []string{
		border("╭", "┬", "╮"),
		"  " + line(header),
	}

	if len(rows) == 0 {
		msg := "no passwords yet"
		inner := TableWidth - 2
		left := (inner - len(msg)) / 2
		out = append(out,
			border("├", "┴", "┤"),
			"  "+bar+strings.Repeat(" ", left)+Dim(msg)+strings.Repeat(" ", inner-left-len(msg))+bar,
			border("╰", "─", "╯"),
		)
		return out
	}

	out = append(out, border("├", "┼", "┤"))
	for i, r := range rows {
		row := line([4]string{
			Dim(Fit(r.ID, colID)),
			Bold(Fit(r.Service, colName)),
			paint(cText, Fit(r.Password, colPass)),
			levelCell(r.Level, colLevel),
		})
		gutter := "  "
		if i == selected {
			gutter = paint("1;"+cAccent, "❯ ")
			row = withBg(row)
		}
		out = append(out, gutter+row)
	}
	out = append(out, border("╰", "┴", "╯"))

	count := fmt.Sprintf("%d entries", len(rows))
	if len(rows) == 1 {
		count = "1 entry"
	}
	return append(out, "  "+Dim(count))
}
