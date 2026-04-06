# cvres

[![Build Status](https://img.shields.io/github/actions/workflow/status/ivangzn/cvres/go.yml?branch=main)](https://github.com/ivangzn/cvres/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ivangzn/cvres)](https://goreportcard.com/report/github.com/ivangzn/cvres)
[![License](https://img.shields.io/github/license/ivangzn/cvres)](https://github.com/ivangzn/cvres/blob/main/LICENSE)

## About

Command-line application that allows to create professional resumes with
different visual formats quickly and easily.

[See an example](example/example.pdf).

## Building

Clone this repo:

```bash
git clone git@github.com:ivangzn/cvres.git
```

Compile cvres:

```bash
go build .
```

## Usage

To fill your resume, you first need to create a file containing all your
profile data.

Supported formats:

- YAML, [see example](example/example.yaml).
- JSON, [see example](example/example.json).

See a more readable specification of the format [here](example/README.md).

To generate a resume:

```bash
# from a file
cvres resume.yaml > output.html

# with a specific style
cvres -style ale resume.json > output.html

# from stdin (pipe or redirect)
cvres < resume.yaml > output.html
curl https://example.com/resume.yaml | cvres > output.html
cat resume.json | cvres > output.html
```

*This generates a minified HTML file. To save it as PDF, open it in any web
browser and use the print dialog.*

## Styles

The following styles are available:

- Ale

## Flags

- `-style` sets the style to use when generating the resume (default: `ale`).
- `-styles` prints all available styles.
- `-help` prints usage information.

