package table_test

import (
	"csv/table"
	"csv/table/table_row"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTableRowValidator(t *testing.T) {
	row := table_row.TableRow{Cols: []string{"21-08-2001", "First", "Last", "5000"}}

	validator := table.Validator{}

	result := validator.Validate(&row)

	fmt.Println(result)

	assert.Equal(t, 0, len(result))
}

func TestTableRowValidatorInvalid(t *testing.T) {
	row := table_row.TableRow{Cols: []string{"33-08-2001", "", "", "-1"}}

	validator := table.Validator{}

	result := validator.Validate(&row)

	fmt.Println(result)

	assert.Equal(t, 4, len(result))
}
