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
	Resume      resume.Resume
	EmailSvg    template.HTML
	LocationSvg template.HTML
	LinkedinSvg template.HTML
}

// Render renders a resume using the Ale style.
func Render(w io.Writer, r *resume.Resume) (int64, error) {
	data := templateData{
		Resume:      *r,
		EmailSvg:    template.HTML(static.EmailSvg),
		LocationSvg: template.HTML(static.LocationSvg),
		LinkedinSvg: template.HTML(static.LinkedinSvg),
	}

	tmpl, err := template.New("resume").Funcs(template.FuncMap{
		"ToUpper": strings.ToUpper,
	}).Parse(html)
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
