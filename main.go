package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ivangzn/cvres/encode"
	"github.com/ivangzn/cvres/styles"
	"gopkg.in/yaml.v3"
)

func main() {
	styleName := flag.String("style", "ale", "style name to be used. Use -styles to list them all.")
	listStyles := flag.Bool("styles", false, "show all style names.")
	flag.Parse()

	if *listStyles {
		names := strings.Join(styles.Names(), ", ")
		fmt.Printf("avaiable styles: %s\n", names)
		return
	}

	inPath := flag.Arg(0)
	if inPath == "" {
		exit("missing input file path")
	}

	outPath := flag.Arg(1)
	if outPath == "" {
		outPath = "resume.html"
	}

	resume := LoadResumeData(inPath)
	GenerateResume(resume, outPath, strings.ToLower(*styleName))
}

func LoadResumeData(path string) *encode.Resume {
	inFile, err := os.Open(path)
	if err != nil {
		exit(err)
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
		exit("file extension not supported")
	}
	if err != nil {
		exit(err)
	}

	return resume
}

func GenerateResume(data *encode.Resume, path string, styleName string) {
	outFile, err := os.Create(path)
	if err != nil {
		exit(err)
	}

	defer outFile.Close()
	style, err := styles.NewStyle(styleName, data)
	if err != nil {
		exit(err)
	}

	_, err = style.WriteTo(outFile)
	if err != nil {
		exit(err)
	}
}

func exit(cause any) {
	fmt.Println(cause)
	os.Exit(1)
}
