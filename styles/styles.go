package styles

import (
	"errors"

	"github.com/ivangzn/cvres/resume"
	"github.com/ivangzn/cvres/styles/ale"
)

// encoderCtor represents a constructor for an encode.Encoder.
type encoderCtor func() resume.Style

// encoders holds references for the name and constructor of each style supported.
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

// NewStyle finds a style by its name. Returns nil and an error if the style doesn't exist.
func NewStyle(name string) (resume.Style, error) {
	newEncoder, ok := encoders[name]
	if !ok {
		return nil, errors.New("style not found")
	}
	return newEncoder(), nil
}
