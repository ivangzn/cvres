package resume

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Resume struct {
	style Style

	Person   Person    `json:"person" yaml:"person"`
	Contact  Contact   `json:"contact" yaml:"contact"`
	Sections []Section `json:"sections" yaml:"sections"`
}

type Person struct {
	Name string `json:"name" yaml:"name"`
	Role string `json:"role" yaml:"role"`
}

type Contact struct {
	Email    string `json:"email" yaml:"email"`
	Location string `json:"location" yaml:"location"`
	LinkedIn string `json:"linkedin" yaml:"linkedin"`
}

type Section struct {
	Title    string    `json:"title" yaml:"title"`
	Articles []Article `json:"articles" yaml:"articles"`
}

type Article struct {
	What     string   `json:"what" yaml:"what"`
	Where    string   `json:"where" yaml:"where"`
	When     string   `json:"when" yaml:"when"`
	Desc     string   `json:"desc" yaml:"desc"`
	List     []string `json:"list" yaml:"list"`
	FullList []string `json:"full-list" yaml:"full-list"`
}

// NewResume creates a Resume that can be written to a io.Writer using WriteTo.
func NewResume(style Style, decoder Decoder) (*Resume, error) {
	res := &Resume{style: style}
	err := decoder.Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// WriteTo writes a resume to a writer.
func (r *Resume) WriteTo(w io.Writer) (int64, error) {
	n, err := r.style.WriteTo(w, *r)
	return n, err
}

// Style applies a specific way of rendering a Resume,
// and writes the result to a io.Writer by using WriteTo.
type Style interface {
	WriteTo(io.Writer, Resume) (int64, error)
}

// Decoder is the interface that wraps the Decode method.
//
// Decode decodes the profile data of a resume, and stores the result in a Resume.
type Decoder interface {
	Decode(any) error
}

// NewDecoder checks the file extension, and tries to guess the correct Decoder for that file.
//
// Returns nil and an error if the file extension isn't supported.
func NewDecoder(file *os.File) (Decoder, error) {
	var decoder Decoder
	inType := filepath.Ext(file.Name())
	switch inType {
	case ".yaml", ".yml":
		decoder = yaml.NewDecoder(file)
	case ".json":
		decoder = json.NewDecoder(file)
	default:
		return nil, errors.New("file extension not supported")
	}
	return decoder, nil
}
