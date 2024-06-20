package table

import (
	"csv-format/table/tableRow"
	"csv-format/utils/console"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

type Table struct {
	rowValidators []ITableRowValidator

	Headings tableRow.TableRow
	Rows     []tableRow.TableRow
}

func filter[T any](input []T, testFn func(*T) bool) (ret []T) {
	for _, s := range input {
		if testFn(&s) {
			ret = append(ret, s)
		}
	}
	return
}

func (t *Table) GetValidRows() []tableRow.TableRow {
	return filter(t.Rows, func(r *tableRow.TableRow) bool {
		return r.Valid()
	})
}

func (t *Table) GetInvalidRows() []tableRow.TableRow {
	return filter(t.Rows, func(r *tableRow.TableRow) bool {
		return !r.Valid()
	})
}

func getLongestWordCountInColumnRow(rows []tableRow.TableRow, col int) int {
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

func EncodeWithErrors(separator string, heading *tableRow.TableRow, rows []tableRow.TableRow) string {
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

// TODO: fix encoder ordering

func Encode(separator string, heading *tableRow.TableRow, rows []tableRow.TableRow) string {
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
	batchSize := int(math.Ceil(float64(amountRows) / float64(64.0)))

	encodingChannel := make(chan EncodingOrderResult)

	fmt.Printf("amount rows %d\n", amountRows)
	fmt.Printf("batchSize: %d\n", batchSize)

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

			encodingChannel <- EncodingOrderResult{content: strings.Join(filter(encodedValues, func(s *string) bool {
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

	start := time.Now()
	output := make([]string, 0)
	for i := 0; i < batchJobs; i++ {
		output = append(output, sortedRows[i].content)
	}

	console.Info("getting data took %s", time.Since(start))

	//output = filter(output, func(s *string) bool {
	//	return len(*s) != 0
	//})

	//fmt.Println(output)

	return fmt.Sprintf(
		"%s\n%s", heading.Encode(separator, longestWords),
		strings.Join(output, "\n"))
}

func (t *Table) Encode(separator string) string {
	return Encode(separator, &t.Headings, t.Rows)
}
