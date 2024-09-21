package buffer

import (
	"fmt"
	"io"
	"os"

	"github.com/HarshKakran/go-editor/terminal"
)

type Buffer struct {
	data []byte
}

func NewBuffer() *Buffer {
	return &Buffer{
		data: make([]byte, 4),
	}
}

func (b *Buffer) ProcessKeyPress() error {
	// os.Stdin.Read can be used while reading in RAW mode.
	n, err := os.Stdin.Read(b.data)
	if err != nil || n == 0 {
		if err != io.EOF {
			fmt.Print("Error reading input: ", err)
			return err
		}
	}

	switch b.data[0] {
	case '\x1b':
		// ESC KEY
		if n >= 3 && b.data[1] == '[' { //Arrow keys
			switch b.data[2] {
			case 'A':
				fmt.Print("Up arrow pressed\r\n")
			case 'B':
				fmt.Print("Down arrow pressed\r\n")
			case 'C':
				fmt.Print("Right arrow pressed\r\n")
			case 'D':
				fmt.Print("Left arrow pressed\r\n")
			}
		}
	case CtrlKey('c'):
		fmt.Printf("Encountered ctrl+c. Exiting...")
		terminal.ExitRawMode()
		os.Exit(0)
	default:
		fmt.Printf("Key Pressed: %v (ASCII Code: %d)\r\n", string(b.data[0]), b.data[0])

	}

	return nil
}

func CtrlKey(k byte) byte {
	return k & 0x1f
}
