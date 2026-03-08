//go:build windows

package ui

import (
	"os"

	"golang.org/x/sys/windows"
)

func enableWindowsVT() {
	stdout := windows.Handle(os.Stdout.Fd())
	var mode uint32
	windows.GetConsoleMode(stdout, &mode)
	windows.SetConsoleMode(stdout, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	stderr := windows.Handle(os.Stderr.Fd())
	var mode2 uint32
	windows.GetConsoleMode(stderr, &mode2)
	windows.SetConsoleMode(stderr, mode2|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
