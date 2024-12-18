# cvres

[![Build Status](https://img.shields.io/github/actions/workflow/status/ivangzn/cvres/go.yml?branch=main)](https://github.com/ivangzn/cvres/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ivangzn/cvres)](https://goreportcard.com/report/github.com/ivangzn/cvres)
[![License](https://img.shields.io/github/license/ivangzn/cvres)](https://github.com/ivangzn/cvres/blob/main/LICENSE)

## About

Command-line application that allows you to create professional resumes in a
specific visual format quickly and easily, using HTML.

[See an example](example/example.pdf).

## Building

Clone this repo:

```bash
git clone https://github.com/ivangzn/cvres
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
