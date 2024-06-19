package command

import (
	"csv-format/table"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var lintDelimiter *string

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func Lint(cmd *cobra.Command, args []string) {
	filePath := args[0]

	if !fileExists(filePath) {
		fmt.Printf("%sFile does not exist%s\n", Red, Reset)
		os.Exit(1)
	}

	if len(*lintDelimiter) != 1 {
		fmt.Printf("%sDelimiter should be 1 character: '%s' given%s\n", Red, *lintDelimiter, Reset)
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file := string(fileContent)
	parsed := table.Parse(file, *lintDelimiter)

	hasError := false
	for i := 0; i < len(parsed.Rows); i++ {
		if !parsed.Rows[i].Valid() {
			hasError = true
			fmt.Printf("%s%s%s\n", Red, strings.Join(parsed.Rows[i].Errors, ", "), Reset)
		}
	}

	if hasError {
		fmt.Printf("\u001B[31mFile contains errors%s\n", Reset)
	} else {
		fmt.Printf("%sFile is valid%s\n", Green, Reset)
	}
}
