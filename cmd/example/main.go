package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/vladimirvikulin/Lab2-Software-Architecture"
)

var (
	exprPtr    = flag.String("e", "", "Expression to evaluate")
	filePtr    = flag.String("f", "", "File containing expression to evaluate")
	outFilePtr = flag.String("o", "", "File to output result")
)

func main() {
	flag.Parse()

	var reader io.Reader
	if *filePtr != "" {
		file, err := os.Open(*filePtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else if *exprPtr != "" {
		reader = strings.NewReader(*exprPtr)
	} else {
		fmt.Fprintln(os.Stderr, "Either -e or -f flag must be provided")
		os.Exit(1)
	}

	var writer io.Writer
	if *outFilePtr != "" {
		file, err := os.Create(*outFilePtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating file:", err)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	computeHandler := &lab2.ComputeHandler{
		Input:  reader,
		Output: writer,
	}

	err := computeHandler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error computing expression:", err)
		os.Exit(1)
	}
}
