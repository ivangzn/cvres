package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/ivangzn/cvres/encode"
	"gopkg.in/yaml.v3"
)

func main() {
	// Handle input
	flag.Parse()

	inPath := flag.Arg(0)
	if inPath == "" {
		panic("missing json file path.")
	}

	outPath := flag.Arg(1)
	if outPath == "" {
		outPath = "cv.html"
	}

	// Read input file
	file, err := os.Open(inPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cv := &encode.Curriculum{}
	inType := filepath.Ext(inPath)
	switch inType {
	case ".yaml", ".yml":
		err = yaml.NewDecoder(file).Decode(cv)
	case ".json":
		err = json.NewDecoder(file).Decode(cv)
	default:
		panic("unsupported file extension.")
	}
	if err != nil {
		panic(err)
	}

	// Generate Cv
	file, err = os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = cv.Html(file)
	if err != nil {
		panic(err)
	}
}
