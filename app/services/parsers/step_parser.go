package parsers

import (
	"errors"
	"strconv"
	customErr "mygolearning/deliverooProject/app/errors"
	"strings"
)

// StepParser implements the CronParserInterface
type StepParser struct {
	min, max int
	step     int
}

func NewStepParser(field string, min, max int) (*StepParser, error) {
	parts := strings.FieldsFunc(field, func(r rune) bool { return r == '/' })
	if len(parts) != 2 {
		return nil, errors.New(customErr.ErrInvalidStepFormat)
	}

	step, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.New(customErr.ErrInvalidFieldValue)
	}

	if step <= 0 || (min != max && step > max-min) {
		return nil, errors.New(customErr.ErrInvalidStepValue)
	}

	return &StepParser{min: min, max: max, step: step}, nil
}

func (sp *StepParser) ExpandField() ([]int, error) {
	var expanded []int
	for i := sp.min; i <= sp.max; i += sp.step {
		expanded = append(expanded, i)
	}
	return expanded, nil
}
