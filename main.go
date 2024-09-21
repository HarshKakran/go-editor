package main

import (
	"fmt"

	"github.com/HarshKakran/go-editor/buffer"
	"github.com/HarshKakran/go-editor/terminal"
)

func main() {
	terminal.EnableRawMode()
	defer terminal.ExitRawMode()

	buffer := buffer.NewBuffer()

	for {
		if err := buffer.ProcessKeyPress(); err != nil {
			break
		}
	}

	fmt.Println("Exiting Raw Mode")
}
