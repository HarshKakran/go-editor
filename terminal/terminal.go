package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Terminal struct {
	Width    int
	Height   int
	oldState *term.State
}

func NewTerminal() (*Terminal, error) {
	w, h, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return &Terminal{}, err
	}

	return &Terminal{
		Width:  w,
		Height: h,
	}, nil
}

func (t *Terminal) EnableRawMode() error {
	// Exit if stdin is disconnected.
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		fmt.Println("Stdin is not connected to the terminal. Exiting...")
		os.Exit(1)
	}

	var err error

	// Enter raw mode
	t.oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	return err
}

func (t *Terminal) ExitRawMode() {
	term.Restore(int(os.Stdin.Fd()), t.oldState)
}
