package plain

import (
	_ "embed"
	"html/template"
	"io"

	"github.com/ivangzn/cvres/resume"
)

//go:embed plain.html
var html string

// templateData contains all the data needed by the template.
type templateData struct {
	Data resume.Data
}

// Render renders a resume using the Plain style.
func Render(w io.Writer, d *resume.Data) error {
	data := templateData{
		Data: *d,
	}

	tmpl := template.New("resume")
	tmpl, err := tmpl.Parse(html)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
