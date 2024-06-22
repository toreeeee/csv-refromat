package table_test

import (
	"csv/table"
	"csv/table/table_row"
	"testing"
)

func benchmarkEncode(amountRows int, b *testing.B) {
	t := table.Table{
		Headings: table_row.TableRow{Cols: []string{"first", "second", "third"}},
		Rows:     []table_row.TableRow{},
	}
	delimiter := ":"

	for i := 0; i < amountRows; i++ {
		t.Rows = append(t.Rows, table_row.TableRow{Cols: []string{"first", "second", "third"}})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = t.Encode(delimiter)
	}
}

func BenchmarkEncode1Row(b *testing.B) {
	benchmarkEncode(1, b)
}

func BenchmarkEncode5Row(b *testing.B) {
	benchmarkEncode(5, b)
}

func BenchmarkEncode10Row(b *testing.B) {
	benchmarkEncode(10, b)
}

func BenchmarkEncode100Row(b *testing.B) {
	benchmarkEncode(100, b)
}

func BenchmarkEncode1000Row(b *testing.B) {
	benchmarkEncode(1000, b)
}

func BenchmarkEncode10000Row(b *testing.B) {
	benchmarkEncode(10000, b)
}

func BenchmarkEncode100000Row(b *testing.B) {
	benchmarkEncode(100000, b)
}
func BenchmarkEncode1000000Row(b *testing.B) {
	benchmarkEncode(1000000, b)
}
