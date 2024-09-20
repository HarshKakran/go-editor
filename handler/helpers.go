package handler

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearToEndOfScreen() {
	fmt.Print("\033[J")
}

func ClearLine() {
	fmt.Print("\033[2K\r")
}

func MoveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", x-1, y+1)
}

func CtrlKey(k uint32) uint32 {
	return k & 0x1f
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func SaveToFile(filePath string, textBuffer []byte) error {
	return os.WriteFile(filePath, textBuffer, 0666)
}
