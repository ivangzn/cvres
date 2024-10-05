package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/ivangzn/cvres/encode"
	"github.com/yosssi/gohtml"
	"gopkg.in/yaml.v3"
)

func main() {
	// Handle input
	formatted := flag.Bool("formatted", false, "formats the output file, if possible.")
	flag.Parse()

	inPath := flag.Arg(0)
	if inPath == "" {
		panic("missing input file path.")
	}

	outPath := flag.Arg(1)
	if outPath == "" {
		outPath = "cv.html"
	}

	// Read input inFile
	inFile, err := os.Open(inPath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	cv := &encode.Curriculum{}
	inType := filepath.Ext(inPath)
	switch inType {
	case ".yaml", ".yml":
		err = yaml.NewDecoder(inFile).Decode(cv)
	case ".json":
		err = json.NewDecoder(inFile).Decode(cv)
	default:
		panic("unsupported file extension.")
	}
	if err != nil {
		panic(err)
	}

	// Generate Cv
	outFile, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = cv.Html(outFile)
	if err != nil {
		panic(err)
	}

	if *formatted {
		err = format(outPath)
		if err != nil {
			panic(err)
		}
	}
}

func format(filePath string) error {
	dir := filepath.Dir(filePath)
	base := filepath.Base(filePath)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)
	path := filepath.Join(dir, name+"-formatted"+ext)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	html, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = file.Write(gohtml.FormatBytes(html))
	return err
}
