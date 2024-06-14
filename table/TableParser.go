package table

import (
	"hello/table/tableRow"
	"strings"
)

func trimSpaces(s string) string {
	return strings.Trim(s, " ")
}

func (t *Table) getAmountHeadings() int {
	return len(t.Headings.Cols)
}

func parseHeadings(items []string) []string {
	for i := 0; i < len(items); i++ {
		items[i] = trimSpaces(items[i])
	}

	return items
}

func Parse(text string, separator string) Table {
	table := Table{rowValidators: []ITableRowValidator{&Validator{}}}

	lines := strings.Split(text, "\n")
	amountLines := len(lines)

	table.Headings = tableRow.TableRow{Cols: parseHeadings(strings.Split(lines[0], separator))}
	table.Rows = make([]tableRow.TableRow, amountLines-1)
	amountHeadings := table.getAmountHeadings()

	for i := 1; i < amountLines; i++ {
		row := tableRow.New(lines[i], amountHeadings)
		for _, v := range table.rowValidators {
			row.Errors = append(row.Errors, v.validate(&row)...)
		}
		table.Rows[i-1] = row

	}

	return table
}
