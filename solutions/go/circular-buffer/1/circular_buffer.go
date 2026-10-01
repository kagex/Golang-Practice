package circularbuffer

import "errors"


// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
    data     []byte
    readIndex  int
    writeIndex int
    count	   int
}

func NewBuffer(size int) *Buffer {
    return &Buffer {
        data : make([]byte, size),
    }
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, errors.New("buffer is empty")
	}

	c := b.data[b.readIndex]
	b.readIndex = (b.readIndex + 1) % len(b.data)
	b.count--
	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.count == len(b.data) {
		return errors.New("buffer is full")
	}

	b.data[b.writeIndex] = c
	b.writeIndex = (b.writeIndex + 1) % len(b.data)
	b.count++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.count == len(b.data) {
		b.readIndex = (b.readIndex + 1) % len(b.data)
		b.count--
	}
	_ = b.WriteByte(c)
}

func (b *Buffer) Reset() {
	b.readIndex, b.writeIndex, b.count = 0, 0, 0
    clear(b.data)
}
