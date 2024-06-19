package command

import (
	"csv-format/table"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
)

var outputFile *string
var outputDelimiter *string
var delimiter *string
var writeToFile *bool

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func Format(cmd *cobra.Command, args []string) {
	inputFile := args[0]

	if !fileExists(inputFile) {
		fmt.Println("Error: Input file not found.")
		os.Exit(1)
	}

	if len(*delimiter) != 1 {
		fmt.Printf("Input delimiter should be 1 character: '%s' given\n", *delimiter)
		os.Exit(1)
	}

	if len(*outputDelimiter) != 1 {
		fmt.Printf("Output delimiter should be 1 character: '%s' given\n", *outputDelimiter)
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(inputFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file := string(fileContent)
	parsed := table.Parse(file, *delimiter)

	encoded := table.Encode(*outputDelimiter, &parsed.Headings, parsed.Rows)

	var outPath string
	if len(*outputFile) != 0 {
		outPath = *outputFile
	} else {
		outPath = inputFile
	}

	if *writeToFile {
		err = os.WriteFile(outPath, []byte(encoded), fs.ModePerm)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Wrote output to file")
	} else {
		fmt.Println(encoded)
	}
}
