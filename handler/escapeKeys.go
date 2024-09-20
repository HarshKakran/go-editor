package handler

import (
	"bufio"
	"log"
)

func HandleEscKeys(r *bufio.Reader, x, y *int, lines []string) rune {
	nxtChar, _, err := r.ReadRune()
	if err != nil {
		log.Panic(err)
	}

	if nxtChar == '[' {
		arrowKey, _, err := r.ReadRune()
		if err != nil {
			log.Panic(err)
		}

		switch arrowKey {
		case 'A': // up key
			if *x > 0 {
				*x--
			}
		case 'B': // down key
			if *x < len(lines) {
				*x++
			}
		case 'D': // left key
			if *y > 0 {
				*y--
			}
		case 'C': // right key
			if *y < len(lines[*x]) {
				*y++
			}
		}
	} else {
		return nxtChar
	}

	MoveCursor(*x, *y)
	return nxtChar
}
