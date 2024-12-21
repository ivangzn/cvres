package styles

import (
	"errors"
	"io"

	"github.com/ivangzn/cvres/encode"
	"github.com/ivangzn/cvres/styles/ale"
)

// encoderCtor represents a constructor for an encode.Encoder.
type encoderCtor func(encode.Resume) encode.Encoder

var encoders = map[string]encoderCtor{
	"ale": ale.New,
}

// Names returns the list of available styles.
func Names() []string {
	names := make([]string, 0, len(encoders))
	for name := range encoders {
		names = append(names, name)
	}
	return names
}

// WriteResume writes a resume with a given style.
func WriteResume(style string, resume *encode.Resume, out io.Writer) error {
	encoder, err := getStyle(style, resume)
	if err != nil {
		return err
	}

	_, err = encoder.WriteTo(out)
	return err
}

// getStyle gets a style by its name, returns nil and an error if the style doesn't exist.
func getStyle(name string, resume *encode.Resume) (encode.Encoder, error) {
	enc, ok := encoders[name]
	if !ok {
		return nil, errors.New("style not found")
	}
	return enc(*resume), nil
}
