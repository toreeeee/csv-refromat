package table

import (
	"csv-format/table/tableRow"
	"fmt"
	"math"
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

	encodingChannel := make(chan string)

	fmt.Printf("amount rows %d\n", amountRows)
	fmt.Printf("batchSize: %d\n", batchSize)

	batchJobs := int(math.Ceil(float64(amountRows) / float64(batchSize)))

	for i := 0; i < batchJobs; i++ {
		batchStart := batchSize * i
		batchEnd := batchSize + batchStart

		go func(batchStart int, batchEnd int) {
			encodedValues := make([]string, batchSize)
			c := 0
			for j := batchStart; j < batchEnd; j++ {
				if j >= amountRows {
					break
				}
				encodedValues[c] = rows[j].Encode(separator, longestWords)
				c++
			}
			encodingChannel <- strings.Join(encodedValues, "\n")
		}(batchStart, batchEnd)
	}

	start := time.Now()
	output := make([]string, amountRows)
	for i := 0; i < batchJobs; i++ {
		v := <-encodingChannel
		output = append(output, v)
	}
	fmt.Printf("getting data took %s \n", time.Since(start))

	//fmt.Println(output)

	return fmt.Sprintf(
		"%s\n%s", heading.Encode(separator, longestWords),
		strings.Join(output, "\n"))
}

func (t *Table) Encode(separator string) string {
	return Encode(separator, &t.Headings, t.Rows)
}
