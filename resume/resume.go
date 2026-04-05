package resume

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// StyleFunc renders a Resume with a given style,
// and writes the result to a io.Writer by using WriteTo.
type StyleFunc func(io.Writer, *Data) error

// Resume is the logical representation of an individual's professional profile data.
type Resume struct {
	style StyleFunc
	Data  Data
}

// Data represents an individual's professional profile data.
type Data struct {
	Person   Person    `json:"person" yaml:"person"`
	Contact  Contact   `json:"contact" yaml:"contact"`
	Sections []Section `json:"sections" yaml:"sections"`
}

// Person represents an individual's personal data.
type Person struct {
	Name string `json:"name" yaml:"name"`
	Role string `json:"role" yaml:"role"`
}

// Contact represents an individual's contact info.
type Contact struct {
	Email    string `json:"email" yaml:"email"`
	Location string `json:"location" yaml:"location"`
	LinkedIn string `json:"linkedin" yaml:"linkedin"`
}

// Section represents a generic container of an individual's relevant professional data.
type Section struct {
	Title    string    `json:"title" yaml:"title"`
	Articles []Article `json:"articles" yaml:"articles"`
}

// Article represents entries within a section of an individual's relevant professional data.
type Article struct {
	What     string   `json:"what" yaml:"what"`
	Where    string   `json:"where" yaml:"where"`
	When     string   `json:"when" yaml:"when"`
	Desc     string   `json:"desc" yaml:"desc"`
	List     []string `json:"list" yaml:"list"`
	FullList []string `json:"full-list" yaml:"full-list"`
}

// NewResume creates a Resume that can be written to a io.Writer using WriteTo.
func NewResume(style StyleFunc, data Data) *Resume {
	return &Resume{
		style: style,
		Data:  data,
	}
}

// WriteTo writes a resume to a writer.
func (r *Resume) WriteTo(w io.Writer) (int64, error) {
	minifier := NewHTMLMinifier()
	err := r.style(minifier, &r.Data)
	if err != nil {
		return 0, fmt.Errorf("style can't write to minifier: %w", err)
	}

	return minifier.WriteTo(w)
}

// Decoder is the interface that wraps the Decode method.
//
// Decode decodes the profile data of a resume, and stores the result in a Resume.
type Decoder interface {
	Decode(any) error
}

// NewDecoder peeks at the input to detect JSON or YAML,
// and returns the matching Decoder.
//
// Returns an error if the format can't be detected.
func NewDecoder(r io.Reader) (Decoder, error) {
	buf := bufio.NewReader(r)

	peeked, err := buf.Peek(buf.Size())
	if err != nil && err != io.EOF && err != bufio.ErrBufferFull {
		return nil, fmt.Errorf("can't read input: %w", err)
	}

	switch guessFormat(peeked) {
	case fileJSON:
		return json.NewDecoder(buf), nil
	case fileYAML:
		return yaml.NewDecoder(buf), nil
	}
	return nil, fmt.Errorf("unknown format")
}
