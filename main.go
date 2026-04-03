package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ivangzn/cvres/resume"
	"github.com/ivangzn/cvres/styles"
)

func main() {
	styleName := flag.String("style", "ale", "style name to be used. Use -styles to list them all.")
	listStyles := flag.Bool("styles", false, "show all style names.")
	flag.Parse()

	// Validate CLI flags.
	if *listStyles {
		names := strings.Join(styles.Names(), ", ")
		fmt.Printf("available styles: %s\n", names)
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

	// Generate resume.
	in, err := os.Open(inPath)
	if err != nil {
		exit(err)
	}
	defer in.Close()

	out, err := os.Create(outPath)
	if err != nil {
		exit(err)
	}
	defer out.Close()

	style, err := styles.NewStyle(*styleName)
	if err != nil {
		exit(err)
	}

	decoder, err := resume.NewDecoder(in, filepath.Ext(inPath))
	if err != nil {
		exit(err)
	}

	var data resume.Data
	err = decoder.Decode(&data)
	if err != nil {
		exit(err)
	}

	res := resume.NewResume(style, data)
	if err != nil {
		exit(err)
	}

	_, err = res.WriteTo(out)
	if err != nil {
		exit(err)
	}
}

func exit(cause any) {
	fmt.Fprintln(os.Stderr, cause)
	os.Exit(1)
}
