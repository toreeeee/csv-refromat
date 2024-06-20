package table

import (
	"csv-format/table/table_row"
	"strconv"
	"strings"
	"time"
)

type ITableRowValidator interface {
	Validate(row *table_row.TableRow) []string
}

type Validator struct {
}

func (v *Validator) Validate(row *table_row.TableRow) []string {
	var result []string

	if birthday, err := row.GetColumn(0); err == nil {
		_, err := time.Parse("02-01-2006", birthday)
		if err != nil {
			result = append(result, err.Error())
		}
	} else {
		result = append(result, "Birthday is missing from row")
	}

	if firstName, err := row.GetColumn(1); err == nil {
		if strings.Contains(firstName, " ") || len(firstName) == 0 {
			result = append(result, "Incorrect first name")
		}
	} else {
		result = append(result, "First name is missing from row")
	}

	if lastName, err := row.GetColumn(2); err == nil {
		if strings.Contains(lastName, " ") || len(lastName) == 0 {
			result = append(result, "Incorrect last name")
		}
	} else {
		result = append(result, "Last name is missing from row")
	}

	if salary, err := row.GetColumn(3); err == nil {
		value, err := strconv.ParseFloat(salary, 32)
		if err != nil {
			result = append(result, err.Error())
		}

		if value < 0 {
			result = append(result, "Salary cannot be negative")
		}
	} else {
		result = append(result, "Salary is missing from row")
	}

	return result
}
