package main

import (
	"fmt"
	"os"

	"github.com/HarshKakran/go-editor/buffer"
	"github.com/HarshKakran/go-editor/edi"
	"github.com/HarshKakran/go-editor/handler"
)

func main() {
	filePath := os.Args[0]

	ui := edi.NewUI()
	buffer := buffer.NewBuffer(ui, filePath)

	buffer.Load()

	handler.ClearScreen()
	fmt.Println("Press ctrl+c to exit or :wq to save and exit.")

	fmt.Println(buffer.GetData())
	buffer.SetFocus()
}
