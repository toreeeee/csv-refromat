package table

import (
	"csv/table/table_row"
	"csv/utils/console"
	"csv/utils/extra_math"
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
	rows []table_row.TableRow
	idx  int
}

func Parse(text string, delimiter string) Table {
	table := Table{}

	lines := strings.Split(text, "\n")
	amountLines := len(lines)

	table.Headings = table_row.TableRow{Cols: parseHeadings(strings.Split(lines[0], delimiter))}
	table.Rows = make([]table_row.TableRow, 0)
	amountHeadings := table.getAmountHeadings()

	if amountHeadings < 2 {
		console.Warn("Table only has %d columns", amountHeadings)
	}

	amountThreads := extra_math.Min(64, amountLines-1)

	batchSize := int(math.Round(float64(amountLines) / float64(amountThreads)))
	parsingChannel := make(chan ParsedTableRowBatch, batchSize)
	batchJobs := int(math.Ceil(float64(amountLines) / float64(batchSize)))

	for i := 0; i < batchJobs; i++ {
		batchStart := batchSize * i
		batchEnd := batchSize + batchStart

		if batchStart <= 1 {
			batchStart = 1
		}

		if batchEnd > amountLines {
			batchEnd = amountLines
		}

		go func(batchStart int, batchEnd int, batchSize int, idx int) {
			output := make([]table_row.TableRow, batchSize)

			c := 0
			for j := batchStart; j < batchEnd; j++ {
				output[c] = table_row.New(lines[j], j+1, amountHeadings, delimiter)
				c++
			}

			parsingChannel <- ParsedTableRowBatch{rows: output, idx: idx}

		}(batchStart, batchEnd, batchEnd-batchStart, i)
	}

	sortedRows := make([]ParsedTableRowBatch, 0)
	for i := 0; i < batchJobs; i++ {
		sortedRows = append(sortedRows, <-parsingChannel)
	}
	sort.SliceStable(sortedRows, func(i, j int) bool {
		return sortedRows[i].idx < sortedRows[j].idx
	})

	for i := 0; i < batchJobs; i++ {
		table.Rows = append(table.Rows, sortedRows[i].rows...)
	}

	return table
}
