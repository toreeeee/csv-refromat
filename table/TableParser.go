package table

import (
	"csv-format/table/tableRow"
	"csv-format/utils/array"
	"csv-format/utils/console"
	"encoding/json"
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

	const amountThreads = 32

	batchSize := int(math.Round(float64(amountLines) / float64(amountThreads)))
	parsingChannel := make(chan ParsedTableRowBatch)
	batchJobs := int(math.Ceil(float64(amountLines) / float64(batchSize)))

	console.Info("amount jobs: %d", batchJobs)

	for i := 0; i < batchJobs; i++ {
		batchStart := batchSize * i
		batchEnd := (batchSize + batchStart)

		if batchStart <= 1 {
			batchStart = 1
		}

		if batchEnd > amountLines {
			batchEnd = amountLines
		}

		console.Info("batchStart: %d, batchEnd: %d", batchStart, batchEnd)

		go func(batchStart int, batchEnd int, batchSize int, idx int) {
			output := make([]tableRow.TableRow, batchSize)

			c := 0
			for j := batchStart; j < batchEnd; j++ {
				//console.Info("%d", j)
				output[c] = tableRow.New(lines[j], j+1, amountHeadings, delimiter)
				c++
			}

			//console.Info("%s", output)

			parsingChannel <- ParsedTableRowBatch{rows: output, idx: idx}

		}(batchStart, batchEnd, batchEnd-batchStart, i)
	}

	//startJobs := time.Now()

	sortedRows := make([]ParsedTableRowBatch, 0)
	for i := 0; i < batchJobs; i++ {
		//console.Info("processing batch %d", i)
		rows := <-parsingChannel
		sortedRows = append(sortedRows, rows)
		//batchStart := batchSize * i

		//console.Info("BatchStart: %d;", batchStart)
		//batchEnd := batchSize + batchStart

		//for j := 0; j < len(rows); j++ {
		//	if j+batchStart >= amountLines-1 {
		//		continue
		//	}
		//
		//	table.Rows[j+batchStart] = rows[j]
		//
		//	fmt.Println(j + batchStart)
		//}

		//table.Rows = append(table.Rows, rows...)
	}
	sort.SliceStable(sortedRows, func(i, j int) bool {
		return sortedRows[i].idx < sortedRows[j].idx
	})

	j, _ := json.Marshal(array.Map(sortedRows, func(t *ParsedTableRowBatch, i int) []tableRow.TableRow {
		return t.rows
	}))

	console.Info("%s", string(j))

	for i := 0; i < batchJobs; i++ {
		table.Rows = append(table.Rows, sortedRows[i].rows...)
	}

	//table.Rows = append(table.Rows, rows...)

	//table.Rows = filter(table.Rows, func(t *tableRow.TableRow) bool {
	//	return len(t.Cols) != 0
	//})

	//console.Info("waiting for chan: %s", time.Since(startJobs))

	//for i := 1; i < amountLines; i++ {
	//	if len(lines[i]) == 0 {
	//		continue
	//	}
	//	row := tableRow.New(lines[i], i+1, amountHeadings, delimiter)
	//	//for _, v := range table.rowValidators {
	//	//	row.Errors = append(row.Errors, v.validate(&row)...)
	//	//}
	//	table.Rows[i-1] = row
	//
	//}

	return table
}
