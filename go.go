package main

import (
	"fmt"
	table "hello/table"
	"hello/table/tableRow"
	"strings"
)

func parseCSV(content string) {
	lines := strings.Split(content, "\n")

	var rows []tableRow.TableRow

	for _, line := range lines {
		rows = append(rows, tableRow.New(line, 2))
	}

	fmt.Println(rows)
}

// Main function
func main() {
	parsed := table.Parse("  first  :  second  :  third  \nhello world hello:5:6\n7:8", ":")

	fmt.Println(parsed)

	fmt.Println("valid: ", parsed.GetValidRows())
	fmt.Println("invalid: ", parsed.GetInvalidRows())

	//parsed.Headings.Cols = append(parsed.Headings.Cols, "errors")
	fmt.Println(table.EncodeWithErrors(":", parsed.Headings, parsed.GetInvalidRows()))
}
