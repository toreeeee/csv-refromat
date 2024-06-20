package main

import (
	"csv-format/command"
	table "csv-format/table"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

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

// Main function
func main() {
	command.Init()
}
