package ui

import (
	"fmt"
	"io"
	"os"

	"github.com/HarshKakran/go-editor/terminal"
)

type UI struct {
	T  *terminal.Terminal
	R  *os.File
	W  *os.File
	CX int
	CY int
}

func NewUI(t *terminal.Terminal) *UI {
	return &UI{
		R:  os.Stdin,
		W:  os.Stdout,
		T:  t,
		CX: 0,
		CY: 0,
	}
}

func (u *UI) ReadKeyPress(buffer []byte) (int, error) {
	// os.Stdin.Read can be used while reading in RAW mode.
	n, err := u.R.Read(buffer)
	if err != nil || n == 0 {
		if err != io.EOF {
			return n, err
		}
	}
	return n, err
}

func (u *UI) Exit(msg string, err error) {
	u.ClearScreenForExit()

	fmt.Printf("%v: %v", msg, err)
	os.Exit(1)
}

func (u *UI) ClearScreenForExit() {
	u.W.Write([]byte("\x1b[2J"))
	u.W.Write([]byte("\x1b[H"))
}
