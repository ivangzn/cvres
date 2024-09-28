# cvres

Command-line application that allows you to create professional resumes in a
specific visual format quickly and easily, using HTML.

HTML allows you to open your curriculum or resume using any web browser, and
save as PDF.

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

To generate the curriculum or resume:

```bash
cvres ./example.json ./output.html
```
