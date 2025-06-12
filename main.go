package main

import (
	"flag"
	"fmt"
	"os"
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

	// Generate resume.
	data, err := os.Open(inPath)
	if err != nil {
		exit(err)
	}
	defer data.Close()

	output, err := os.Create(outPath)
	if err != nil {
		exit(err)
	}
	defer output.Close()

	decoder, err := resume.NewDecoder(data)
	if err != nil {
		exit(err)
	}

	style, err := styles.NewStyle(*styleName)
	if err != nil {
		exit(err)
	}

	resume, err := resume.NewResumeFromDecoder(style, decoder)
	if err != nil {
		exit(err)
	}

	_, err = resume.WriteTo(output)
	if err != nil {
		exit(err)
	}
}

func exit(cause any) {
	fmt.Println(cause)
	os.Exit(1)
}
