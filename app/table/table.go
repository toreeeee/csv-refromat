package table

import (
	"csv/table/table_row"
	"csv/utils/array"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Table struct {
	rowValidators []ITableRowValidator

	Headings table_row.TableRow
	Rows     []table_row.TableRow
}

func (t *Table) GetValidRows() []table_row.TableRow {
	return array.Filter(t.Rows, func(r *table_row.TableRow) bool {
		return r.Valid()
	})
}

func (t *Table) GetInvalidRows() []table_row.TableRow {
	return array.Filter(t.Rows, func(r *table_row.TableRow) bool {
		return !r.Valid()
	})
}

func getLongestWordCountInColumnRow(rows []table_row.TableRow, col int) int {
	longest := 0

	for i := 0; i < len(rows); i++ {
		if col >= len(rows[i].Cols) {
			continue
		}

		s := len(rows[i].Cols[col])
		if s > longest {
			longest = s
		}
	}

	return longest
}

func EncodeWithErrors(separator string, heading *table_row.TableRow, rows []table_row.TableRow) string {
	heading.Cols = append(heading.Cols, "errors")
	for i := 0; i < len(rows); i++ {
		rows[i].Cols = append(rows[i].Cols, strings.Join(rows[i].Errors, ", "))
	}

	return Encode(separator, heading, rows)
}

type EncodingOrderResult struct {
	content string
	idx     int
}

func Encode(separator string, heading *table_row.TableRow, rows []table_row.TableRow) string {
	headingCount := len(heading.Cols)
	longestWords := make([]int, headingCount)

	for i := 0; i < headingCount; i++ {
		longestWords[i] = len(heading.Cols[i])
	}
	for i := 0; i < headingCount; i++ {
		length := getLongestWordCountInColumnRow(rows, i)
		if length > longestWords[i] {
			longestWords[i] = length
		}
	}

	amountRows := len(rows)
	batchSize := int(math.Ceil(float64(amountRows) / float64(128.0)))

	encodingChannel := make(chan EncodingOrderResult)

	batchJobs := int(math.Ceil(float64(amountRows) / float64(batchSize)))

	for i := 0; i < batchJobs; i++ {
		batchStart := batchSize * i
		batchEnd := (batchSize + batchStart)

		if batchEnd > amountRows {
			batchEnd = amountRows
		}

		go func(batchStart int, batchEnd int, idx int) {
			encodedValues := make([]string, batchSize)
			c := 0
			for j := batchStart; j < batchEnd; j++ {
				encodedValues[c] = rows[j].Encode(separator, longestWords)
				c++
			}

			encodingChannel <- EncodingOrderResult{content: strings.Join(array.Filter(encodedValues, func(s *string) bool {
				return len(*s) > 0
			}), "\n"), idx: idx}
		}(batchStart, batchEnd, i)
	}

	sortedRows := make([]EncodingOrderResult, 0)
	for i := 0; i < batchJobs; i++ {
		v := <-encodingChannel
		sortedRows = append(sortedRows, v)
	}

	sort.SliceStable(sortedRows, func(i, j int) bool {
		return sortedRows[i].idx < sortedRows[j].idx
	})

	output := make([]string, 0)
	for i := 0; i < batchJobs; i++ {
		output = append(output, sortedRows[i].content)
	}

	return fmt.Sprintf(
		"%s\n%s", heading.Encode(separator, longestWords),
		strings.Join(output, "\n"))
}

func (t *Table) Encode(separator string) string {
	return Encode(separator, &t.Headings, t.Rows)
}
