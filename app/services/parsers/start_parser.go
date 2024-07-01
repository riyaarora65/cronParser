package parsers

import (
	"errors"
)

// StarParser implements the CronParserInterface
type StarParser struct {
	min, max int
}

func NewStarParser(min, max int) *StarParser {
	return &StarParser{min: min, max: max}
}

func (sp *StarParser) ExpandField() ([]int, error) {
	if sp.min > sp.max {
		return nil, errors.New("min value cannot be greater than max value")
	}

	var expanded []int
	for i := sp.min; i <= sp.max; i++ {
		expanded = append(expanded, i)
	}
	return expanded, nil
}
