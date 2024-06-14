package main

import (
	"errors"
	"flag"
	"fmt"
	table "hello/table"
	"hello/table/tableRow"
	"io/fs"
	"os"
	"strings"
	"time"
)

func parseCSV(content string) {
	lines := strings.Split(content, "\n")

	var rows []tableRow.TableRow

	for _, line := range lines {
		rows = append(rows, tableRow.New(line, 2))
	}

	fmt.Println(rows)
}

var (
	inputFile  = flag.String("i", "", "-i <file_path>")
	outputFile = flag.String("o", "out.csv", "-o <file_path>")
	format     = flag.Bool("f", false, "add -f to format file and save to input path")
	delimiter  = flag.String("d", ":", "-d <delimiter>")
)

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func reformatFile(t *table.Table) {
	start := time.Now()
	encoded := table.Encode(*delimiter, t.Headings, t.GetInvalidRows())
	duration := time.Since(start)
	fmt.Printf("Reformat took %s\n", duration.String())

	err := os.WriteFile(*inputFile, []byte(encoded), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Took %s\n", time.Since(start).String())
}

func process() {
	flag.Parse()

	if *inputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if !fileExists(*inputFile) {
		flag.Usage()
		fmt.Printf("input '%s' file not found", *inputFile)
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*inputFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file := string(fileContent)

	start := time.Now()
	parsed := table.Parse(file, *delimiter)
	parseDuration := time.Since(start)
	fmt.Printf("Parsing took %s\n", parseDuration.String())

	if *format {
		reformatFile(&parsed)
		return
	}

	encoded := table.EncodeWithErrors(*delimiter, parsed.Headings, parsed.GetInvalidRows())
	duration := time.Since(start)

	fmt.Printf("\nTook %s\n\n", duration.String())

	err = os.WriteFile(*outputFile, []byte(encoded), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Main function
func main() {
	start := time.Now()
	process()
	duration := time.Since(start)

	fmt.Printf("process took %s\n", duration.String())
}
