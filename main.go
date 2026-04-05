package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ivangzn/cvres/resume"
	"github.com/ivangzn/cvres/styles"
)

func main() {
	styleName := flag.String("style", "ale", "style name to be used. Use -styles to list them all.")
	listStyles := flag.Bool("styles", false, "show all style names.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: cvres [flags] [file]\n\n")
		fmt.Fprintf(os.Stderr, "Reads resume data (JSON or YAML) from a file or stdin,\n")
		fmt.Fprintf(os.Stderr, "and writes HTML to stdout.\n\n")
		fmt.Fprintf(os.Stderr, "Examples:\n")
		fmt.Fprintf(os.Stderr, "  cvres resume.yaml > output.html\n")
		fmt.Fprintf(os.Stderr, "  cvres -style ale resume.json > output.html\n")
		fmt.Fprintf(os.Stderr, "  cat resume.yaml | cvres > output.html\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	var in io.Reader
	if flag.Arg(0) == "" {
		in = os.Stdin
	} else {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			exit(err)
		}
		in = f
	}

	// Validate CLI flags.
	if *listStyles {
		names := strings.Join(styles.Names(), ", ")
		fmt.Printf("available styles: %s\n", names)
		return
	}

	// Generate resume.
	style, err := styles.NewStyle(*styleName)
	if err != nil {
		exit(err)
	}

	decoder, err := resume.NewDecoder(in)
	if err != nil {
		exit(err)
	}

	var data resume.Data
	err = decoder.Decode(&data)
	if err != nil {
		exit(err)
	}

	res := resume.NewResume(style, data)

	_, err = res.WriteTo(os.Stdout)
	if err != nil {
		exit(err)
	}
}

func exit(cause any) {
	fmt.Fprintln(os.Stderr, cause)
	os.Exit(1)
}
