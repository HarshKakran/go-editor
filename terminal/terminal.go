package terminal

import "golang.org/x/sys/unix"

type State struct {
	Termios unix.Termios
}

func MakeRaw(fd int) (*State, error) {
	return makeraw(fd)
}

func Restore(fd int, state *State) error {
	return restore(fd, state)
}

func GetState(fd int) (*State, error) {
	return getState(fd)
}
