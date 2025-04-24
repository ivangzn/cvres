package styles

import (
	"errors"

	"github.com/ivangzn/cvres/resume"
	"github.com/ivangzn/cvres/styles/ale"
)

// styleCtor represents a constructor for a resume style.
type styleCtor func() resume.Style

// styles holds references for the name and constructor of each style supported.
var styles = map[string]styleCtor{
	"ale": ale.New,
}

// Names returns the list of available styles.
func Names() []string {
	names := make([]string, 0, len(styles))
	for name := range styles {
		names = append(names, name)
	}
	return names
}

// NewStyle finds a style by its name. Returns nil and an error if the style doesn't exist.
func NewStyle(name string) (resume.Style, error) {
	newStyle, ok := styles[name]
	if !ok {
		return nil, errors.New("style not found")
	}
	return newStyle(), nil
}
