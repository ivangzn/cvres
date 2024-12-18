package encode

// ByteCounter is a io.Writer that counts how many bytes have been written.
type ByteCounter struct {
	written int64
}

// Write adds the number of bytes written to ByteCounter.Count
func (bc *ByteCounter) Write(p []byte) (int, error) {
	n := len(p)
	bc.written += int64(n)
	return n, nil
}

func (bc *ByteCounter) Count() int64 {
	return bc.written
}
