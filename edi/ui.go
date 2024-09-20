package edi

import (
	"io"
	"os"
)

type UI struct {
	R io.Reader
	W io.Writer
}

func NewUI() *UI {
	return &UI{
		R: os.Stdout,
		W: os.Stdin,
	}
}

func (u *UI) ReadChar() (string, error) {
	var rBuf []byte

	for {
		var buf [1]byte
		_, err := u.R.Read(buf[:])
		if err != nil && err != io.EOF {
			return "", err
		}

		rBuf = append(rBuf, buf[0])

		if len(rBuf) > 0 {
			break
		}
	}

	return string(rBuf), nil
}
