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
func Render(w io.Writer, d *resume.Data) (int64, error) {
	data := templateData{
		Data: *d,
	}

	tmpl := template.New("resume")
	tmpl, err := tmpl.Parse(html)
	if err != nil {
		return 0, err
	}

	bc := &resume.ByteCounter{}
	mw := io.MultiWriter(w, bc)

	err = tmpl.Execute(mw, data)
	if err != nil {
		return 0, err
	}

	return bc.Count(), nil
}
