package command

import (
	"csv-format/table"
	"csv-format/utils/console"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var lintDelimiter *string

func Lint(cmd *cobra.Command, args []string) {
	filePath := args[0]

	if !fileExists(filePath) {
		fmt.Printf("%sFile does not exist%s\n", console.Red, console.Reset)
		os.Exit(1)
	}

	if len(*lintDelimiter) != 1 {
		fmt.Printf("%sDelimiter should be 1 character: '%s' given%s\n", console.Red, *lintDelimiter, console.Reset)
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
			fmt.Printf("%s%s%s\n", console.Red, strings.Join(parsed.Rows[i].Errors, ", "), console.Reset)
		}
	}

	if hasError {
		fmt.Printf("%sFile contains errors%s\n", console.Red, console.Reset)
	} else {
		fmt.Printf("%sFile is valid%s\n", console.Green, console.Reset)
	}
}
