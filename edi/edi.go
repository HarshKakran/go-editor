package edi

import (
	"fmt"
	"os"
	"time"

	"github.com/HarshKakran/go-editor/buffer"
	"github.com/HarshKakran/go-editor/ui"
)

type Editor struct {
	B *buffer.Buffer
	U *ui.UI
}

func NewEditor(b *buffer.Buffer, u *ui.UI) *Editor {
	return &Editor{
		B: b,
		U: u,
	}
}

func (e *Editor) ProcessKeyPress() error {
	n, err := e.U.ReadKeyPress(e.B.RBuf)
	if err != nil {
		fmt.Print("Error reading input: ", err)
		return err
	}

	switch e.B.RBuf[0] {
	case '\x1b':
		// ESC KEY
		if n >= 3 && e.B.RBuf[1] == '[' { //Arrow keys
			switch e.B.RBuf[2] {
			case 'A':
				//Up
				e.U.CY--
			case 'B':
				// Down
				e.U.CY++
			case 'C':
				// Right
				e.U.CX++
			case 'D':
				//Left
				e.U.CX--
			}
		}
	case 127:
		if len(e.B.LBuf) > 0 {
			e.B.LBuf = e.B.LBuf[:len(e.B.LBuf)-1]
			e.PrintLBuf()
		}

	case CtrlKey('c'):
		fmt.Printf("^C")
		time.Sleep(1 * time.Second)
		e.U.T.ExitRawMode()

		e.U.ClearScreenForExit()

		os.Exit(0)
	case CtrlKey('s'):
		fmt.Print("Saving the file...")
		time.Sleep(1 * time.Second)
		e.U.T.ExitRawMode()

		e.U.ClearScreenForExit()

		os.Exit(0)
	case '\r':
		// Handling ENTER key press

		if len(e.B.LBuf) <= 0 {
			e.B.FBuf = append(e.B.FBuf, []byte(" \r\n"))
		} else {
			e.B.LBuf = append(e.B.LBuf, []byte("\r\n")...)
			e.B.FBuf = append(e.B.FBuf, e.B.LBuf)
			e.B.LBuf = make([]byte, 0)
		}

		e.PrintFBuf()
	default:
		// fmt.Printf("Key Pressed: %v (ASCII Code: %d)\r\n", string(e.B.RBuf[0]), e.B.RBuf[0])
		// fmt.Printf("%v", string(e.B.RBuf[0]))
		e.B.LBuf = append(e.B.LBuf, e.B.RBuf[0])
		e.PrintLBuf()

	}

	return nil
}

func CtrlKey(k byte) byte {
	return k & 0x1f
}

func (e *Editor) RefreshScreen() {
	// hide cursor while drawing
	e.B.FBuf = append(e.B.FBuf, []byte("\x1b[?25l"))

	// // escape sequence to clear the screen
	// e.B.FBuf = append(e.B.FBuf, []byte("\x1b[2J"))
	// Move the cursor to the top-left corner of the screen (1;1)
	e.B.FBuf = append(e.B.FBuf, []byte("\x1b[H"))

	e.DrawEmptyRows()

	cursorPos := fmt.Sprintf("\x1b[%d;%dH", e.U.CY+1, e.U.CX+1)
	e.B.FBuf = append(e.B.FBuf, []byte(cursorPos))
	// show cursor after drawing
	e.B.FBuf = append(e.B.FBuf, []byte("\x1b[?25h"))
	e.PrintFBuf()
}

func (e *Editor) DrawEmptyRows() {
	for y := 0; y < e.U.T.Height; y++ {
		if y == e.U.T.Height/3 {
			welcomeMsg := "Welcome to Edi-Go. An editor written in Golang"
			if len(welcomeMsg) > e.U.T.Width {
				welcomeMsg = welcomeMsg[:e.U.T.Width]
			}

			var padding string
			paddingLen := (e.U.T.Width - len(welcomeMsg)) / 2
			if paddingLen > 0 {
				padding += "~"
			}
			for i := 0; i < paddingLen; i++ {
				padding += " "
			}

			e.B.FBuf = append(e.B.FBuf, []byte(padding+welcomeMsg))
		} else {
			e.B.FBuf = append(e.B.FBuf, []byte("~"))
		}
		// clear each line instead of clearing the screen in one go
		e.B.FBuf = append(e.B.FBuf, []byte("\x1b[K"))

		if y < e.U.T.Height-1 {
			e.B.FBuf = append(e.B.FBuf, []byte("\r\n"))
		}
	}
}

func (e *Editor) PrintFBuf() {
	for _, row := range e.B.FBuf {
		e.U.W.Write(row)
	}
}
func (e *Editor) PrintRBuf() {
	e.U.W.Write(e.B.RBuf)
}

func (e *Editor) PrintLBuf() {
	e.U.W.Write([]byte("\x1b[1K\r"))
	e.U.W.Write(e.B.LBuf)
}
