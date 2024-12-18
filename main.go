package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/ivangzn/cvres/encode"
	"github.com/ivangzn/cvres/styles"
	"gopkg.in/yaml.v3"
)

func main() {
	styleName := flag.String("style", "ale", "style name to be used.")
	flag.Parse()

	inPath := flag.Arg(0)
	if inPath == "" {
		panic("missing input file path.")
	}

	outPath := flag.Arg(1)
	if outPath == "" {
		outPath = "resume.html"
	}

	resume := LoadResumeData(inPath)
	GenerateResume(resume, outPath, *styleName)

}

func LoadResumeData(path string) *encode.Resume {
	inFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	resume := &encode.Resume{}
	inType := filepath.Ext(path)
	switch inType {
	case ".yaml", ".yml":
		err = yaml.NewDecoder(inFile).Decode(resume)
	case ".json":
		err = json.NewDecoder(inFile).Decode(resume)
	default:
		panic("unsupported file extension.")
	}
	if err != nil {
		panic(err)
	}

	return resume
}

func GenerateResume(data *encode.Resume, path string, styleName string) {
	outFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()
	style, err := styles.NewStyle(styleName, data)
	if err != nil {
		panic(err)
	}

	_, err = style.WriteTo(outFile)
	if err != nil {
		panic(err)
	}
}
