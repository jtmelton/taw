// taw is a tool that finds all the file extensions and their count in a directory.

package main

import (
	"flag"
	"os"

	"github.com/jtmelton/taw/domain"
	"github.com/jtmelton/taw/processing"
	"github.com/jtmelton/taw/reporting"
)

var (
	inputDirectory *string
	outputFile     *string
)

func main() {
	inputDirectory = flag.String("inputDirectory", "", "Directory to analyze (Required)")
	outputFile = flag.String("outputFile", "", "Output File (Required)")

	flag.Parse()

	if *inputDirectory == "" || *outputFile == "" {
		flag.PrintDefaults()

		os.Exit(1)
	}

	_options := domain.Options{
		InputDirectory: *inputDirectory,
		OutputFile:     *outputFile,
	}

	extensions := processing.Walk(*inputDirectory, _options)

	reporting.WriteReport(extensions, _options)

	/*
		TODO:
		- do a CI setup
			https://github.com/jandelgado/golang-ci-template-github-actions/blob/master/.github/workflows/test.yml

		- add a test for "create user ... identified by $&*Q#*@#(*" in a sql file
	*/

}
