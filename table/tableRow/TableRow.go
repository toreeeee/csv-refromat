package tableRow

import (
	"csv-format/utils/array"
	"fmt"
	"math"
	"strings"
)

type TableRow struct {
	Cols   []string
	Errors []string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func (row TableRow) GetColumn(idx int) (string, error) {
	if len(row.Cols) > idx {
		return row.Cols[idx], nil
	}

	return "", &errorString{"Index out of range"}
}

func getSpaces(l int) string {
	str := ""
	for i := 0; i < l; i++ {
		str = str + " "
	}

	return str
}

func (row *TableRow) Encode(separator string, colLengths []int) string {
	return strings.Join(array.Map(row.Cols, func(t *string, i int) string {
		colLength := colLengths[i] + 6
		valueLength := len(*t)
		spacesFront := math.Ceil((float64(colLength-valueLength) / float64(2.0)))
		spacesEnd := float64(colLength) - (spacesFront + float64(valueLength))

		return fmt.Sprintf("%s%s%s", getSpaces(int(spacesFront)), *t, getSpaces(int(spacesEnd)))
	}), separator)
}

func New(line string, expectedColumns int, delimiter string) TableRow {
	split := strings.Split(line, delimiter)

	for i := 0; i < len(split); i++ {
		split[i] = strings.Trim(split[i], " ")
	}

	row := TableRow{Cols: split}

	if len(row.Cols) != expectedColumns {
		row.Errors = append(row.Errors, fmt.Sprintf("Expected %d column(s), got %d", expectedColumns, len(row.Cols)))
	}

	if len(split) < expectedColumns {
		for len(row.Cols) != expectedColumns {
			row.Cols = append(row.Cols, "   ")
		}
	}

	return row
}

func (row *TableRow) Valid() bool {
	return len(row.Errors) == 0
}
