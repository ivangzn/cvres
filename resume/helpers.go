package resume

import (
	"fmt"
	"io"

	"github.com/tdewolff/minify/v2/minify"
)


// HTMLMinifier helps writing minified HTML content.
//
// The Write method writes the HTML content to be minified.
// WriteTo method writes the minified content.
type HTMLMinifier struct {
	content []byte
}

// NewHTMLMinifier creates a minifier that minifies HTML content.
func NewHTMLMinifier() *HTMLMinifier {
	return &HTMLMinifier{}
}

// Write writes HTML content to be minified.
func (m *HTMLMinifier) Write(data []byte) (int, error) {
	m.content = append(m.content, data...)
	return len(data), nil
}

// WriteTo takes whatever was written by the Write method,
// and writes the minified version to w.
func (m *HTMLMinifier) WriteTo(w io.Writer) (int64, error) {
	minified, err := minify.Default.Bytes("text/html", m.content)
	if err != nil {
		return 0, fmt.Errorf("html couldn't be minified: %w", err)
	}

	n, err := w.Write(minified)
	if err != nil {
		return 0, err
	}
	return int64(n), err
}
