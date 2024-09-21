package terminal

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

var oldState *term.State

func EnableRawMode() {
	// Exit if stdin is disconnected.
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		fmt.Println("Stdin is not connected to the terminal. Exiting...")
		os.Exit(1)
	}

	var err error

	// Enter raw mode
	oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
}

func ExitRawMode() {
	term.Restore(int(os.Stdin.Fd()), oldState)
}
