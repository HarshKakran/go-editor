package buffer

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/HarshKakran/go-editor/edi"
	"github.com/HarshKakran/go-editor/handler"
	"github.com/HarshKakran/go-editor/terminal"
	"golang.org/x/sys/unix"
)

type Buffer struct {
	UI       *edi.UI
	FilePath string
	Data     []string
}

type EState struct {
	Termios unix.Termios
}

func NewBuffer(ui *edi.UI, filePath string) *Buffer {
	return &Buffer{
		UI:       ui,
		FilePath: filePath,
		Data:     make([]string, 0),
	}
}

func (b *Buffer) Load() {
	fh, err := os.Open(b.FilePath)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		scannedText := scanner.Text()
		fmt.Println(scannedText)
		time.Sleep(200)
		b.Data = append(b.Data, scanner.Text())
	}

	// fmt.Println(b.Data)
}

func (b *Buffer) SetFocus() {
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	defer terminal.Restore(int(os.Stdin.Fd()), oldState)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		terminal.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()

	for {
		str, err := b.UI.ReadChar()
		if err != nil {
			panic(err)
		}

		switch str {
		case "[":
			bfReader := bufio.NewReader(b.UI.R)
			var x, y int
			handler.HandleEscKeys(bfReader, &x, &y, b.Data)
		default:
			fmt.Println(str)
		}
	}
}

func (b *Buffer) GetData() string {
	return strings.Join(b.Data, "\r\n")
}
