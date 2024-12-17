package encode

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ivangzn/cvres/static"
)

func (c *Curriculum) Html(w io.Writer) error {
	var errs []error

	html := "<head>%s</head><style>%s</style><body>"
	_, err := fmt.Fprintf(w, html, static.Head, static.Css)
	errs = append(errs, err)

	err = c.Person.Html(w)
	errs = append(errs, err)

	err = c.Contact.Html(w)
	errs = append(errs, err)

	for _, s := range c.Sections {
		err = s.Html(w)
		errs = append(errs, err)
	}

	_, err = fmt.Fprint(w, "</body>")
	errs = append(errs, err)

	return errors.Join(errs...)
}

func (p *Person) Html(w io.Writer) error {
	var errs []error

	_, err := fmt.Fprint(w, `<header id="who">`)
	errs = append(errs, err)

	if p.Name != "" {
		_, err = fmt.Fprintf(w, `<h1 id="who-name">%s</h1>`, p.Name)
		errs = append(errs, err)
	}
	if p.Role != "" {
		_, err = fmt.Fprintf(w, `<h2 id="who-what">%s</h2>`, p.Role)
		errs = append(errs, err)
	}

	_, err = fmt.Fprint(w, "</header>")
	errs = append(errs, err)

	return errors.Join(errs...)
}

func (c *Contact) Html(w io.Writer) error {
	var errs []error

	_, err := fmt.Fprint(w, `<section id="contact"><ul>`)
	errs = append(errs, err)

	itemFmt := "<li>%s<span>%s</span>"
	if c.Email != "" {
		_, err = fmt.Fprintf(w, itemFmt, static.EmailSvg, c.Email)
		errs = append(errs, err)
	}
	if c.Location != "" {
		_, err = fmt.Fprintf(w, itemFmt, static.LocationSvg, c.Location)
		errs = append(errs, err)
	}
	if c.LinkedIn != "" {
		_, err = fmt.Fprintf(w, itemFmt, static.LinkedinSvg, c.LinkedIn)
		errs = append(errs, err)
	}

	_, err = fmt.Fprint(w, "</section>")
	errs = append(errs, err)

	return errors.Join(errs...)
}

func (s *Section) Html(w io.Writer) error {
	var errs []error

	_, err := fmt.Fprint(w, `<section class="info">`)
	errs = append(errs, err)

	if s.Title != "" {
		title := strings.ToUpper(s.Title)
		_, err = fmt.Fprintf(w, `<h3 class="info-title">%s</h3><hr>`, title)
		errs = append(errs, err)
	}
	for _, a := range s.Articles {
		err = a.Html(w)
		errs = append(errs, err)
	}

	_, err = fmt.Fprint(w, "</section>")
	errs = append(errs, err)

	return errors.Join(errs...)
}
func (a *Article) Html(w io.Writer) error {
	var errs []error

	_, err := fmt.Fprint(w, "<article>")
	errs = append(errs, err)

	if a.What != "" {
		_, err = fmt.Fprintf(w, `<h4 class="info-what">%s</h4>`, a.What)
		errs = append(errs, err)
	}
	if a.Where != "" {
		_, err = fmt.Fprintf(w, `<h5 class="info-where">%s</h5>`, a.Where)
		errs = append(errs, err)
	}
	if a.When != "" {
		_, err = fmt.Fprintf(w, `<span class="info-when">%s</span>`, a.When)
		errs = append(errs, err)
	}
	if a.Desc != "" {
		_, err = fmt.Fprintf(w, `<p>%s</p>`, a.Desc)
		errs = append(errs, err)
	}

	// List of items
	if len(a.List) > 0 {
		_, err = fmt.Fprint(w, "<ul>")
		errs = append(errs, err)

		for _, i := range a.List {
			_, err = fmt.Fprintf(w, "<li>%s</li>", i)
			errs = append(errs, err)
		}

		_, err = fmt.Fprint(w, "</ul>")
		errs = append(errs, err)
	}

	// Full list of items
	if len(a.FullList) > 0 {
		_, err = fmt.Fprint(w, `<ul class="full-list">`)
		errs = append(errs, err)

		for _, i := range a.FullList {
			_, err = fmt.Fprintf(w, "<li>%s</li>", i)
			errs = append(errs, err)
		}

		_, err = fmt.Fprint(w, "</ul>")
		errs = append(errs, err)
	}

	_, err = fmt.Fprint(w, "</article>")
	errs = append(errs, err)

	return errors.Join(errs...)
}
