package buffer

type Buffer struct {
	RBuf []byte
	LBuf []byte
	FBuf [][]byte
}

func NewBuffer() *Buffer {
	return &Buffer{
		RBuf: make([]byte, 32),
		LBuf: make([]byte, 0),
		FBuf: make([][]byte, 0),
	}
}
