package table_test

import (
	"csv-format/table"
	"csv-format/table/table_row"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	parsed := table.Parse("first:second:third", ":")

	expected := table.Table{Rows: make([]table_row.TableRow, 0), Headings: table_row.TableRow{Cols: []string{"first", "second", "third"}}}

	assert.Equal(t, expected, parsed)

	assert.Equal(t, 0, len(parsed.Rows))

	first, err := parsed.Headings.GetColumn(0)
	assert.Nil(t, err)
	assert.Equal(t, "first", first)
	second, err := parsed.Headings.GetColumn(1)
	assert.Nil(t, err)
	assert.Equal(t, "second", second)
	third, err := parsed.Headings.GetColumn(2)
	assert.Nil(t, err)
	assert.Equal(t, "third", third)
}

func BenchmarkParser(b *testing.B) {
	input := "first:second:third\nhello:world:hello"
	delimiter := ":"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = table.Parse(input, delimiter)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "

func randInt(min int, max int) int {
	return min + rand.IntN(max-min)
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.IntN(len(letterBytes))]
	}
	return string(b)
}

var result table.Table

func benchmarkParse(amountRows int, b *testing.B) {
	header := fmt.Sprintf(
		"%s:%s:%s",
		randStringBytes(randInt(10, 16)),
		randStringBytes(randInt(10, 16)),
		randStringBytes(randInt(10, 16)),
	)
	content := make([]string, amountRows)

	for i := 0; i < amountRows; i++ {
		content = append(content, fmt.Sprintf(
			"%s:%s:%s",
			randStringBytes(randInt(10, 16)),
			randStringBytes(randInt(10, 16)),
			randStringBytes(randInt(10, 16)),
		))
	}

	str := fmt.Sprintf("%s\n%s", header, strings.Join(content, "\n"))

	b.ResetTimer()
	var r table.Table
	for i := 0; i < b.N; i++ {
		r = table.Parse(str, ":")
	}
	b.StopTimer()
	result = r
}

func BenchmarkParseRandomData1(b *testing.B) {
	benchmarkParse(1, b)
}
func BenchmarkParseRandomData5(b *testing.B) {
	benchmarkParse(5, b)
}
func BenchmarkParseRandomData10(b *testing.B) {
	benchmarkParse(10, b)
}
func BenchmarkParseRandomData100(b *testing.B) {
	benchmarkParse(100, b)
}
func BenchmarkParseRandomData1000(b *testing.B) {
	benchmarkParse(1000, b)
}
func BenchmarkParseRandomData10000(b *testing.B) {
	benchmarkParse(10000, b)
}
func BenchmarkParseRandomData100000(b *testing.B) {
	benchmarkParse(100000, b)
}
func BenchmarkParseRandomData1000000(b *testing.B) {
	benchmarkParse(1000000, b)
}
