package encode

import (
	"fmt"

	"github.com/tdewolff/minify/v2/minify"
)

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

// MinifyCSS minifies any css content.
// In case of errors, returns the original css.
func MinifyCSS(css string) string {
	minified, err := minify.CSS(css)
	if err != nil {
		fmt.Println("css couldn't be minified, falling back to original css.")
		minified = css
	}
	return minified
}
