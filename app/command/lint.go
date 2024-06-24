package command

import (
	"csv/table"
	"csv/utils/console"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var lintDelimiter *string

func Lint(cmd *cobra.Command, args []string) {
	filePath := args[0]

	if !fileExists(filePath) {
		console.Error("File does not exist")
		os.Exit(1)
	}

	if len(*lintDelimiter) != 1 {
		console.Error("Delimiter should be 1 character: '%s' given", *lintDelimiter)
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
			console.Error("%s", strings.Join(parsed.Rows[i].Errors, ", "))
		}
	}

	if hasError {
		console.Error("File contains errors")
	} else {
		console.Error("File is valid")
	}
}
