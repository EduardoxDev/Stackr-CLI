package ui

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	Reset   string
	Bold    string
	Dim     string
	Green   string
	Red     string
	Yellow  string
	Blue    string
	Cyan    string
	Magenta string
	Gray    string
	White   string
)

func init() {
	if runtime.GOOS == "windows" {
		enableWindowsVT()
	}
	Reset   = "\x1b[0m"
	Bold    = "\x1b[1m"
	Dim     = "\x1b[2m"
	Green   = "\x1b[32m"
	Red     = "\x1b[31m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[38;5;99m"
	Cyan    = "\x1b[38;5;141m"
	Magenta = "\x1b[38;5;135m"
	Gray    = "\x1b[90m"
	White   = "\x1b[97m"
}

type Column struct {
	Key   string
	Label string
}

type Row map[string]string

type Spinner struct {
	msg  string
	done chan struct{}
}

func Ok(msg string)   { fmt.Printf("  %s✔%s  %s\n", Green+Bold, Reset, msg) }
func Fail(msg string) { fmt.Fprintf(os.Stderr, "  %s✖%s  %s\n", Red+Bold, Reset, msg) }
func Info(msg string) { fmt.Printf("  %s·%s  %s\n", Cyan+Bold, Reset, msg) }
func Warn(msg string) { fmt.Printf("  %s⚠%s  %s\n", Yellow+Bold, Reset, msg) }
func Hint(msg string) { fmt.Printf("     %s%s%s\n", Gray, msg, Reset) }

func Label(key, val string) {
	fmt.Printf("    %s%-20s%s%s\n", Gray, key, Reset, val)
}

func LabelSecret(key, val string) {
	masked := val
	if len(val) > 6 {
		masked = val[:4] + strings.Repeat("·", len(val)-6) + val[len(val)-2:]
	}
	fmt.Printf("    %s%-20s%s%s%s%s\n", Gray, key, Reset, Dim, masked, Reset)
}

func Header(title string) {
	fmt.Printf("\n  %s%s%s\n  %s%s%s\n\n",
		Bold+White, title, Reset,
		Gray, strings.Repeat("─", len(title)+2), Reset,
	)
}

func SectionTitle(title string) {
	fmt.Printf("\n  %s%s%s\n\n", Bold+Cyan, title, Reset)
}

func StatusBadge(status string) string {
	switch strings.ToUpper(status) {
	case "RUNNING":
		return Green + Bold + "● online   " + Reset
	case "STOPPED":
		return Red + "● stopped  " + Reset
	case "BUILDING":
		return Yellow + "● building " + Reset
	case "ERROR":
		return Red + Bold + "● error    " + Reset
	case "RESTARTING":
		return Cyan + "● restart  " + Reset
	default:
		if status == "" {
			return Gray + "● unknown  " + Reset
		}
		return Gray + status + Reset
	}
}

func EngineBadge(engine string) string {
	switch strings.ToLower(engine) {
	case "postgresql", "postgres":
		return Cyan + "postgres" + Reset
	case "mysql":
		return Blue + "mysql   " + Reset
	case "mongodb", "mongo":
		return Green + "mongo   " + Reset
	case "redis":
		return Red + "redis   " + Reset
	default:
		return Gray + engine + Reset
	}
}

func PrintTable(rows []Row, cols []Column) {
	if len(rows) == 0 {
		return
	}
	widths := make([]int, len(cols))
	for i, col := range cols {
		widths[i] = len(col.Label)
		for _, row := range rows {
			if l := len(StripANSI(row[col.Key])); l > widths[i] {
				widths[i] = l
			}
		}
	}

	fmt.Print("    ")
	for i, col := range cols {
		fmt.Printf("%s%-*s%s  ", Gray, widths[i], strings.ToUpper(col.Label), Reset)
	}
	fmt.Println()

	fmt.Print("    ")
	for _, w := range widths {
		fmt.Printf("%s%s%s  ", Gray, strings.Repeat("─", w), Reset)
	}
	fmt.Println()

	for _, row := range rows {
		fmt.Print("    ")
		for i, col := range cols {
			val := row[col.Key]
			pad := widths[i] - len(StripANSI(val))
			if pad < 0 {
				pad = 0
			}
			fmt.Printf("%s%s  ", val, strings.Repeat(" ", pad))
		}
		fmt.Println()
	}
	fmt.Println()
}

func NewSpinner(msg string) *Spinner {
	s := &Spinner{msg: msg, done: make(chan struct{})}
	go func() {
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		fmt.Print("\x1b[?25l")
		for {
			select {
			case <-s.done:
				return
			default:
				fmt.Printf("\r  %s%s%s  %s", Cyan, frames[i%len(frames)], Reset, msg)
				i++
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()
	return s
}

func (s *Spinner) Stop(msg string) {
	close(s.done)
	time.Sleep(120 * time.Millisecond)
	fmt.Printf("\r\x1b[K\x1b[?25h")
	if msg != "" {
		Ok(msg)
	}
}

func (s *Spinner) Fail(msg string) {
	close(s.done)
	time.Sleep(120 * time.Millisecond)
	fmt.Printf("\r\x1b[K\x1b[?25h")
	if msg != "" {
		Fail(msg)
	}
}

func Banner(version string) {
	// Gradiente roxo escuro → roxo claro → magenta
	grad := []string{
		"\x1b[38;5;54m",  // roxo escuro
		"\x1b[38;5;55m",
		"\x1b[38;5;91m",
		"\x1b[38;5;99m",
		"\x1b[38;5;105m",
		"\x1b[38;5;135m",
		"\x1b[38;5;141m",
		"\x1b[38;5;147m", // lavanda claro
	}

	border   := grad[7] + "  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░" + Reset
	borderMid := grad[7] + "  ░" + Reset
	pad      := "                                                       "

	fmt.Println()
	fmt.Println(border)
	fmt.Println(borderMid + pad + grad[7] + "░" + Reset)

	art := []string{
		`  ░  ██████╗ ████████╗ █████╗  ██████╗██╗  ██╗██████╗   ░`,
		`  ░  ██╔════╝╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝██╔══██╗  ░`,
		`  ░  ╚█████╗    ██║   ███████║██║     █████╔╝ ██████╔╝   ░`,
		`  ░   ╚═══██╗   ██║   ██╔══██║██║     ██╔═██╗ ██╔══██╗   ░`,
		`  ░  ██████╔╝   ██║   ██║  ██║╚██████╗██║  ██╗██║  ██║   ░`,
		`  ░  ╚═════╝    ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝  ░`,
	}

	for i, line := range art {
		ci := i + 1
		if ci >= len(grad) {
			ci = len(grad) - 1
		}
		fmt.Printf("%s%s%s\n", grad[ci], line, Reset)
	}

	fmt.Println(borderMid + pad + grad[7] + "░" + Reset)
	fmt.Println(border)
	fmt.Println()

	fmt.Printf("    %sv%s%s    %s·%s  plataforma de cloud hosting  %s·%s  stackr.lat%s\n\n",
		Bold+White, version, Reset,
		Gray, Reset,
		Gray, Reset+Dim,
		Reset,
	)
}

func PadRight(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return s + strings.Repeat(" ", n-len(s))
}

func StripANSI(s string) string {
	var out []rune
	inEsc := false
	for _, r := range s {
		if r == '\x1b' {
			inEsc = true
			continue
		}
		if inEsc {
			if r == 'm' {
				inEsc = false
			}
			continue
		}
		out = append(out, r)
	}
	return string(out)
}
