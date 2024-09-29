# cvres

Command-line application that allows you to create professional resumes in a
specific visual format quickly and easily, using HTML.

[See an example](example/example.pdf).

## Building

To build the executable:

```bash
git clone https://github.com/ivangzn/cvres
go build .
```

## Usage

To fill your resume or curriculum vitae, you first need to create a file
containing all your data.

Supported type of files:

- YAML, [see example](example/example.yaml)
- JSON, [see example](example/example.json)

In order to render the HTML properly, you must have installed
[Alegreya](https://fonts.google.com/specimen/Alegreya) font in your system.

To generate the curriculum or resume:

```bash
cvres ./example.json ./output.html
```

This will generate an HTML file, if you want to save it to PDF, you can use
any web browser that has that feature.  
