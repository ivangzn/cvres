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
		fmt.Println("missing input file path.")
		os.Exit(1)
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
		fmt.Println(err)
		os.Exit(1)
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
		fmt.Println("unsupported file extension.")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resume
}

func GenerateResume(data *encode.Resume, path string, styleName string) {
	outFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer outFile.Close()
	style, err := styles.NewStyle(styleName, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = style.WriteTo(outFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
