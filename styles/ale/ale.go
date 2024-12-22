package ale

import (
	_ "embed"
	"io"
	"strings"

	"github.com/ivangzn/cvres/resume"
	"github.com/ivangzn/cvres/static"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

//go:embed ale.css
var cssContent string

// Ale is a resume style based on Google's Alegreya font.
type Ale struct {
	resume *resume.Resume
}

// New creates the Ale resume style.
func New() resume.Style {
	return &Ale{}
}

// WriteTo writes the resume in HTML format.
func (a *Ale) WriteTo(w io.Writer, r resume.Resume) (n int64, err error) {
	a.resume = &r

	bc := &resume.ByteCounter{}
	mw := io.MultiWriter(w, bc)
	err = a.html().Render(mw)
	if err != nil {
		return 0, err
	}
	return bc.Count(), nil
}

func (a *Ale) html() Node {
	return HTML(
		a.head(),
		a.body(),
		a.css(),
	)
}

func (a *Ale) head() Node {
	return Head(
		Link(Href("https://fonts.googleapis.com")),
		Link(Rel("preconnect"), Href("https://fonts.gstatic.com"), CrossOrigin("")),
		Link(Rel("stylesheet"), Href("https://fonts.googleapis.com/css2?family=Alegreya:ital,wght@0,400..900;1,400..900&display=swap")),
	)
}

func (a *Ale) css() Node {
	return Rawf("<style>%s</style>", cssContent)
}

func (a *Ale) body() Node {
	return Body(
		a.person(),
		a.contact(),
		a.sections(),
	)
}

func (a *Ale) person() Node {
	person := a.resume.Person
	return Header(
		ID("who"),
		H1(ID("who-name"), Text(person.Name)),
		H2(ID("who-what"), Text(person.Role)),
	)
}

func (a *Ale) contact() Node {
	contact := a.resume.Contact
	return Section(
		ID("contact"),
		Ul(
			Li(Raw(static.EmailSvg), Text(contact.Email)),
			Li(Raw(static.LocationSvg), Text(contact.Location)),
			Li(Raw(static.LinkedinSvg), Text(contact.LinkedIn)),
		),
	)
}

func (a *Ale) sections() Node {
	listItem := func(item string) Node {
		return Li(Text(item))
	}

	article := func(a resume.Article) Node {
		hasList := len(a.List) > 0
		hasFullList := len(a.FullList) > 0
		return Article(
			H4(Class("info-what"), Text(a.What)),
			H5(Class("info-where"), Text(a.Where)),
			Span(Class("info-when"), Text(a.When)),
			P(Text(a.Desc)),

			If(hasList,
				Ul(
					Map(a.List, listItem),
				),
			),

			If(hasFullList,
				Ul(
					Class("full-list"),
					Map(a.FullList, listItem),
				),
			),
		)
	}

	section := func(s resume.Section) Node {

		title := strings.ToUpper(s.Title)
		return Section(
			Class("info"),
			H3(Class("info-title"), Text(title)),
			Hr(),
			Map(s.Articles, article),
		)
	}

	return Map(a.resume.Sections, section)
}
