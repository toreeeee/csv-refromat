package main

import (
	"csv-format/command"
	table "csv-format/table"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"
)

//var (
//	inputFile    = flag.String("i", "", "-i <file_path>")
//	outputFile   = flag.String("o", "invalid-values.csv", "-o <file_path>")
//	format       = flag.Bool("f", false, "add -f to format file and save to input path")
//	delimiter    = flag.String("d", ":", "-d <delimiter>")
//	outDelimiter = flag.String("od", "", "Output delimiter")
//)

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func reformatFile(t *table.Table, outPath string) {
	encoded := table.Encode(".", &t.Headings, t.Rows)

	err := os.WriteFile(outPath, []byte(encoded), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func process() {
	//command.Init()

	//if *inputFile == "" {
	//	flag.Usage()
	//	os.Exit(1)
	//}
	//
	//if !fileExists(*inputFile) {
	//	flag.Usage()
	//	fmt.Printf("input '%s' file not found", *inputFile)
	//	os.Exit(1)
	//}
	//
	//if len(*outDelimiter) == 0 {
	//	*outDelimiter = *delimiter
	//}
	//
	//fileContent, err := os.ReadFile(*inputFile)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//file := string(fileContent)
	//
	//parsed := table.Parse(file, *delimiter)
	//
	//if *format {
	//	reformatFile(&parsed, *inputFile)
	//	return
	//}
	//
	//encoded := table.EncodeWithErrors(*delimiter, &parsed.Headings, parsed.GetInvalidRows())
	//
	//err = os.WriteFile(*outputFile, []byte(encoded), fs.ModePerm)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
}

// Main function
func main() {
	command.Init()

	start := time.Now()
	process()
	duration := time.Since(start)

	fmt.Printf("process took %s\n", duration.String())
}
