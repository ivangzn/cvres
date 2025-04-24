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

Supported type of files:

- YAML, [see example](example/example.yaml).
- JSON, [see example](example/example.json).

To generate the resume:

```bash
cvres ./example.json ./output.html
```

This will generate a minified HTML file, if you want to save it as PDF, you can
use any web browser.  

To generate the resume, using a specific style:

```bash
cvres -style ale ./example.json ./output.html 
```

## Styles

The following styles are avaiable:

- Ale

## Flags

- `-style` sets the style to use when generating the resume.
- `-styles` prints all avaiable styles.
