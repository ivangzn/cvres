package encode

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"strings"
)

//go:embed styles.css
var css string

var head string = `
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Alegreya:ital,wght@0,400..900;1,400..900&display=swap" rel="stylesheet">`

const emailSvg string = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#000000" style="width: 1em; height: 1em; fill: rgb(0, 0, 0);"><path d="M19.68 20c1.414 0 2.56-1.194 2.56-2.667V10.5l-9.765 4.072a1.005 1.005 0 01-.475.095c-.232 0-.392-.034-.472-.099L1.76 10.5v6.833C1.76 18.806 2.906 20 4.32 20h15.36z"></path><path d="M12 11.9l10.24-4.267v-.966C22.24 5.194 21.094 4 19.68 4H4.32C2.906 4 1.76 5.194 1.76 6.667v.966L12 11.9z"></path></svg>`
const locationSvg string = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#000000" style="width: 1em; height: 1em; fill: rgb(0, 0, 0);"><path d="M12.3 24a39.034 39.034 0 01-4.5-4.707c-2.057-2.547-4.5-6.341-4.5-9.957-.002-3.775 2.191-7.18 5.556-8.625 3.364-1.445 7.237-.646 9.81 2.025 1.693 1.748 2.642 4.124 2.634 6.6 0 3.616-2.443 7.41-4.5 9.957A39.041 39.041 0 0112.3 24zm0-18.663c-1.378 0-2.651.762-3.34 2a4.127 4.127 0 000 3.999c.689 1.237 1.962 2 3.34 2 2.13 0 3.857-1.791 3.857-4 0-2.208-1.727-3.999-3.857-3.999z"></path></svg>`
const linkedinSvg string = `<svg role="img" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" fill="#000000" style="width: 1em; height: 1em; fill: rgb(0, 0, 0);"><path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"></path></svg>`

func (c *Curriculum) Html(w io.Writer) error {
	var errs []error

	html := "<head>%s</head><style>%s</style><body>"
	_, err := fmt.Fprintf(w, html, head, css)
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
		_, err = fmt.Fprintf(w, itemFmt, emailSvg, c.Email)
		errs = append(errs, err)
	}
	if c.Location != "" {
		_, err = fmt.Fprintf(w, itemFmt, locationSvg, c.Location)
		errs = append(errs, err)
	}
	if c.LinkedIn != "" {
		_, err = fmt.Fprintf(w, itemFmt, linkedinSvg, c.LinkedIn)
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
