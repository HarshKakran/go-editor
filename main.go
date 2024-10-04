package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/HarshKakran/go-editor/buffer"
	"github.com/HarshKakran/go-editor/edi"
	"github.com/HarshKakran/go-editor/terminal"
	"github.com/HarshKakran/go-editor/ui"
)

func main() {
	terminal, err := terminal.NewTerminal()
	if err != nil {
		// get cursor position if error received from NewTerminal()
		n, err := os.Stdout.Write([]byte("\x1b[999C\x1b[999B"))
		if err != nil || n != 12 {
			panic(err)
		}

		terminal.Height, terminal.Width = GetCursorPosition()
	}

	terminal.EnableRawMode()
	defer terminal.ExitRawMode()

	ui := ui.NewUI(terminal)
	buffer := buffer.NewBuffer()
	editor := edi.NewEditor(buffer, ui)
	editor.RefreshScreen()

	for {
		if err := editor.ProcessKeyPress(); err != nil {
			ui.Exit("While processing key press", err)
		}
	}
}

func GetCursorPosition() (int, int) {
	buf := make([]byte, 32)

	n, err := os.Stdout.Write([]byte("\x1b[6n"))
	if err != nil || n != 4 {
		panic(err)
	}

	_, err = os.Stdin.Read(buf)
	if err != nil {
		panic(err)
	}

	if buf[0] != '\x1b' || buf[1] != '[' {
		panic(fmt.Errorf("invalid escape sequence"))
	}

	fields := strings.Split(string(buf[2:]), ";")

	for _, field := range fields {
		fmt.Println("p", field)
	}

	rows, err := strconv.Atoi(fields[0])
	if err != nil {
		panic(err)
	}

	cols, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}

	return rows, cols
}
