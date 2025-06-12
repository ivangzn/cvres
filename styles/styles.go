package styles

import (
	"errors"

	"github.com/ivangzn/cvres/resume"
	"github.com/ivangzn/cvres/styles/ale"
	"github.com/ivangzn/cvres/styles/plain"
)

// styles maps style names to their corresponding rendering functions.
var styles = map[string]resume.StyleFunc{
	"ale":   ale.Render,
	"plain": plain.Render,
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
func NewStyle(name string) (resume.StyleFunc, error) {
	s, ok := styles[name]
	if !ok {
		return nil, errors.New("style not found")
	}
	return s, nil
}
