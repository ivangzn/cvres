// Ale is a resume style based on Google's Alegreya font.

package ale

import (
	_ "embed"
	"html/template"
	"io"
	"strings"

	"github.com/ivangzn/cvres/resume"
	"github.com/ivangzn/cvres/static"
)

//go:embed ale.html
var html string

// templateData contains all the data needed by the template.
type templateData struct {
	Data        resume.Data
	EmailSvg    template.HTML
	LocationSvg template.HTML
	LinkedinSvg template.HTML
}

// Render renders a resume using the Ale style.
func Render(w io.Writer, d *resume.Data) (int64, error) {
	data := templateData{
		Data:        *d,
		EmailSvg:    template.HTML(static.EmailSvg),
		LocationSvg: template.HTML(static.LocationSvg),
		LinkedinSvg: template.HTML(static.LinkedinSvg),
	}

	tmpl := template.New("resume")
	tmpl = tmpl.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
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
