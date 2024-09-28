package encode

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/yosssi/gohtml"
)

//go:embed styles.css
var css string

const emailSvg string = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#000000" style="width: 1em; height: 1em; fill: rgb(0, 0, 0);"><path d="M19.68 20c1.414 0 2.56-1.194 2.56-2.667V10.5l-9.765 4.072a1.005 1.005 0 01-.475.095c-.232 0-.392-.034-.472-.099L1.76 10.5v6.833C1.76 18.806 2.906 20 4.32 20h15.36z"></path><path d="M12 11.9l10.24-4.267v-.966C22.24 5.194 21.094 4 19.68 4H4.32C2.906 4 1.76 5.194 1.76 6.667v.966L12 11.9z"></path></svg>`
const locationSvg string = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#000000" style="width: 1em; height: 1em; fill: rgb(0, 0, 0);"><path d="M12.3 24a39.034 39.034 0 01-4.5-4.707c-2.057-2.547-4.5-6.341-4.5-9.957-.002-3.775 2.191-7.18 5.556-8.625 3.364-1.445 7.237-.646 9.81 2.025 1.693 1.748 2.642 4.124 2.634 6.6 0 3.616-2.443 7.41-4.5 9.957A39.041 39.041 0 0112.3 24zm0-18.663c-1.378 0-2.651.762-3.34 2a4.127 4.127 0 000 3.999c.689 1.237 1.962 2 3.34 2 2.13 0 3.857-1.791 3.857-4 0-2.208-1.727-3.999-3.857-3.999z"></path></svg>`
const linkedinSvg string = `<svg role="img" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" fill="#000000" style="width: 1em; height: 1em; fill: rgb(0, 0, 0);"><path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"></path></svg>`

func (c *Curriculum) Html() string {
	html := "<style>" + css + "</style>"
	html += "<body>"
	html += c.Person.Html()
	html += c.Contact.Html()
	for _, s := range c.Sections {
		html += s.Html() + ""
	}
	html += "</body>"
	return gohtml.Format(html)
}

func (p *Person) Html() string {
	html := `<header id="who">`
	if p.Name != "" {
		html += `<h1 id="who-name">` + p.Name + "</h1>"
	}
	if p.Role != "" {
		html += `<h2 id="who-what">` + p.Role + "</h2>"
	}
	html += "</header>"
	return html
}

func (c *Contact) Html() string {
	html := `<section id="contact"><ul>`
	itemFmt := "<li>%s<span>%s</span>"
	if c.Email != "" {
		html += fmt.Sprintf(itemFmt, emailSvg, c.Email)
	}
	if c.Location != "" {
		html += fmt.Sprintf(itemFmt, locationSvg, c.Location)
	}
	if c.LinkedIn != "" {
		html += fmt.Sprintf(itemFmt, linkedinSvg, c.LinkedIn)
	}
	html += "</section>"
	return gohtml.Format(html)
}

func (s *Section) Html() string {
	html := `<section class="info">`
	if s.Title != "" {
		title := strings.ToUpper(s.Title)
		html += `<h3 class="info-title">` + title + "</h3>"
		html += "<hr>"
	}
	for _, a := range s.Articles {
		html += a.Html() + ""
	}
	html += "</section>"
	return gohtml.Format(html)
}

func (a *Article) Html() string {
	html := "<article>"
	if a.What != "" {
		html += `<h4 class="info-what">` + a.What + "</h4>"
	}
	if a.Where != "" {
		html += `<h5 class="info-where">` + a.Where + "</h5>"
	}
	if a.When != "" {
		html += `<span class="info-when">` + a.When + "</span>"
	}
	if a.Desc != "" {
		html += `<p>` + a.Desc + "</p>"
	}
	if len(a.List) > 0 {
		html += "<ul>"
		for _, i := range a.List {
			html += fmt.Sprintf("<li>%s</li>", i)
		}
		html += "</ul>"
	}
	if len(a.FullList) > 0 {
		html += `<ul class="full-list">`
		for _, i := range a.FullList {
			html += fmt.Sprintf("<li>%s</li>", i)
		}
		html += "</ul>"
	}
	html += "</article>"
	return gohtml.Format(html)
}
