package table

import (
	"hello/table/tableRow"
	"strconv"
	"strings"
	"time"
)

type ITableRowValidator interface {
	validate(row *tableRow.TableRow) []string
}

type Validator struct {
}

func (v *Validator) validate(row *tableRow.TableRow) []string {
	var result []string

	if birthday, err := row.GetColumn(0); err != nil {
		_, err := time.Parse("%d-%m-%Y", birthday)
		if err != nil {
			result = append(result, err.Error())
		}
	} else {
		result = append(result, "Birthday is missing from row")
	}

	if firstName, err := row.GetColumn(1); err != nil {
		if strings.Contains(firstName, " ") {
			result = append(result, "Incorrect first name")
		}
	} else {
		result = append(result, "First name is missing from row")
	}

	if lastName, err := row.GetColumn(1); err != nil {
		if strings.Contains(lastName, " ") {
			result = append(result, "Incorrect last name")
		}
	} else {
		result = append(result, "Last name is missing from row")
	}

	if salary, err := row.GetColumn(1); err != nil {
		_, err := strconv.ParseFloat(salary, 32)
		if err != nil {
			result = append(result, err.Error())
		}
	} else {
		result = append(result, "Salary is missing from row")
	}

	return result
}