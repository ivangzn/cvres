package styles

import (
	"errors"

	"github.com/ivangzn/cvres/encode"
	"github.com/ivangzn/cvres/styles/ale"
)

// encoderCtor represents a function that generates an encode.Encoder
type encoderCtor func(encode.Resume) encode.Encoder

var encoders = map[string]encoderCtor{
	"ale": ale.New,
}

func NewStyle(name string, resume *encode.Resume) (encode.Encoder, error) {
	enc, ok := encoders[name]
	if !ok {
		return nil, errors.New("style not found")
	}
	return enc(*resume), nil
}
