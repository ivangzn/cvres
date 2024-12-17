package static

import _ "embed"

//go:embed styles.css
var Css string

//go:embed head.html
var Head string

//go:embed email.svg
var EmailSvg string

//go:embed location.svg
var LocationSvg string

//go:embed linkedin.svg
var LinkedinSvg string
