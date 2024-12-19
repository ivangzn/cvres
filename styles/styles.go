package styles

import (
	"errors"

	"github.com/ivangzn/cvres/encode"
	"github.com/ivangzn/cvres/styles/ale"
)

// encoderCtor represents a constructor for an encode.Encoder.
type encoderCtor func(encode.Resume) encode.Encoder

var encoders = map[string]encoderCtor{
	"ale": ale.New,
}

// NewStyle returns a resume encoder that uses a given style.
func NewStyle(name string, resume *encode.Resume) (encode.Encoder, error) {
	enc, ok := encoders[name]
	if !ok {
		return nil, errors.New("style not found")
	}
	return enc(*resume), nil
}

// Names returns the list of available styles.
func Names() []string {
	names := make([]string, 0, len(encoders))
	for name := range encoders {
		names = append(names, name)
	}
	return names
}
