package encode

import (
	"testing"
)

func TestByteCounter(t *testing.T) {

	tests := []struct {
		input    []byte
		expected int64
	}{
		{[]byte("hello world&;"), 13},
		{[]byte(""), 0},
	}

	for _, tt := range tests {
		bc := ByteCounter{}

		_, err := bc.Write(tt.input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		got := bc.Count()
		if got != tt.expected {
			t.Errorf("expected %d bytes written, got %d", tt.expected, got)
		}
	}
}
