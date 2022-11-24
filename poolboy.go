package poolboy

import (
	"io"
	"sync"
)

type AppendByteBuffer struct {
	Buffer []byte
}

type Pool struct {
	size uint64
	pool sync.Pool
}

func (a *AppendByteBuffer) NewBuffer() []byte {
	buf := a.Buffer

	// Getting the start and end of the capacity of the buffer
	beginning := int64(len(buf))
	end := int64(cap(buf))

	i := beginning

	// Initializing the buffer if needed
	if end == 0 {
		i = 64
		buf = make([]byte, i)
	} else {
		buf = buf[:end]
	}

	return buf
}

// Implementing io.ReaderFrom interface
func (b *AppendByteBuffer) AppendBuffer(from io.Reader) {
	buf := b.NewBuffer()

}
