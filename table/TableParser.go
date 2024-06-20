package table

import (
	"csv-format/table/tableRow"
	"math"
	"sort"
	"strings"
)

func trimSpaces(s *string) string {
	return strings.Trim(*s, " ")
}

func (t *Table) getAmountHeadings() int {
	return len(t.Headings.Cols)
}

func parseHeadings(items []string) []string {
	for i := 0; i < len(items); i++ {
		items[i] = trimSpaces(&items[i])
	}

	return items
}

type ParsedTableRowBatch struct {
	rows []tableRow.TableRow
	idx  int
}

func Parse(text string, delimiter string) Table {
	table := Table{rowValidators: []ITableRowValidator{&Validator{}}}

	lines := strings.Split(text, "\n")
	amountLines := len(lines)

	table.Headings = tableRow.TableRow{Cols: parseHeadings(strings.Split(lines[0], delimiter))}
	table.Rows = make([]tableRow.TableRow, 0)
	amountHeadings := table.getAmountHeadings()

	const amountThreads = 64

	batchSize := int(math.Round(float64(amountLines) / float64(amountThreads)))
	parsingChannel := make(chan ParsedTableRowBatch)
	batchJobs := int(math.Ceil(float64(amountLines) / float64(batchSize)))

	for i := 0; i < batchJobs; i++ {
		batchStart := batchSize * i
		batchEnd := (batchSize + batchStart)

		if batchStart <= 1 {
			batchStart = 1
		}

		if batchEnd > amountLines {
			batchEnd = amountLines
		}

		go func(batchStart int, batchEnd int, batchSize int, idx int) {
			output := make([]tableRow.TableRow, batchSize)

			c := 0
			for j := batchStart; j < batchEnd; j++ {
				output[c] = tableRow.New(lines[j], j+1, amountHeadings, delimiter)
				c++
			}

			parsingChannel <- ParsedTableRowBatch{rows: output, idx: idx}

		}(batchStart, batchEnd, batchEnd-batchStart, i)
	}

	sortedRows := make([]ParsedTableRowBatch, 0)
	for i := 0; i < batchJobs; i++ {
		rows := <-parsingChannel
		sortedRows = append(sortedRows, rows)
	}
	sort.SliceStable(sortedRows, func(i, j int) bool {
		return sortedRows[i].idx < sortedRows[j].idx
	})

	for i := 0; i < batchJobs; i++ {
		table.Rows = append(table.Rows, sortedRows[i].rows...)
	}

	return table
}
