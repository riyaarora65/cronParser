package parsers

import (
	"errors"
	customErr "mygolearning/deliverooProject/app/errors"
	"strconv"
	"strings"
)

// RangeParser implements the CronParserInterface
type RangeParser struct {
	start, end int
}

func NewRangeParser(field string) (*RangeParser, error) {
	parts := strings.SplitN(field, "-", 2) 
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, errors.New(customErr.ErrInvalidRangeFormat)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, errors.New(customErr.ErrInvalidFieldValue)
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.New(customErr.ErrInvalidFieldValue)
	}

	if start > end {
		return nil, errors.New(customErr.ErrInvalidStartAndEndRange)
	}

	return &RangeParser{start: start, end: end}, nil
}

func (rp *RangeParser) ExpandField() ([]int, error) {
	var expanded []int
	for i := rp.start; i <= rp.end; i++ {
		expanded = append(expanded, i)
	}
	return expanded, nil
}
