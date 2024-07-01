package parsers

import (
	"strconv"
	"strings"
)

// ListParser implements the CronParserInterface
type ListParser struct {
	values []int
}

func NewListParser(field string) (*ListParser, error) {
	parts := strings.Split(field, ",")
	var values []int

	for _, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}

	return &ListParser{values: values}, nil
}

func (lp *ListParser) ExpandField() ([]int, error) {
	return lp.values, nil
}
