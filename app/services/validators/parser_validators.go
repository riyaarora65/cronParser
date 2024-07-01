package validators

import (
	"errors"
	customErr "mygolearning/deliverooProject/app/errors"
	"mygolearning/deliverooProject/app/utils"
	"strconv"
	"strings"
)

// FieldValidator interface
type FieldValidator interface {
	Validate(field string) error
}

// MinuteFieldValidator validates minute fields
type MinuteFieldValidator struct{}

func (v *MinuteFieldValidator) Validate(field string) error {
	isValid,err := isValidField(field, 0, 59)
	if err != nil {
		return err
	}
	if !(*isValid){
		return errors.New(customErr.ErrInvalidRange)
	}
	return nil
}

// HourFieldValidator validates hour fields
type HourFieldValidator struct{}

func (v *HourFieldValidator) Validate(field string) error {
	isValid,err := isValidField(field, 0, 23)
	if err != nil {
		return err
	}
	if !(*isValid) {
		return errors.New(customErr.ErrInvalidRange)
	}
	return nil
}

// DayOfMonthFieldValidator validates day of month fields
type DayOfMonthFieldValidator struct{}

func (v *DayOfMonthFieldValidator) Validate(field string) error {
	isValid,err := isValidField(field, 1, 31)
	if err != nil {
		return err
	}
	if !(*isValid) {
		return errors.New(customErr.ErrInvalidRange)
	}
	return nil
}

// MonthFieldValidator validates month fields
type MonthFieldValidator struct{}

func (v *MonthFieldValidator) Validate(field string) error {
	isValid,err := isValidField(field, 1, 12)
	if err != nil {
		return err
	}
	if !(*isValid) {
		return errors.New(customErr.ErrInvalidRange)
	}
	return nil
}

// DayOfWeekFieldValidator validates day of week fields
type DayOfWeekFieldValidator struct{}

func (v *DayOfWeekFieldValidator) Validate(field string) error {
	isValid,err :=  isValidField(field, 1, 7)
	if err != nil {
		return err
	}
	if !(*isValid) {
		return errors.New(customErr.ErrInvalidRange)
	}
	return nil
}

func isValidField(field string, min, max int) (*bool, error) {
	switch {
	case field == "*":
		return utils.BoolPtr(true), nil
	case strings.Contains(field, "/"):
		return isValidStepField(field, min, max)
	case strings.Contains(field, "-"):
		return isValidRangeField(field, min, max)
	case strings.Contains(field, ","):
		return isValidListField(field, min, max)
	default:
		return isValidNumber(field, min, max)
	}
}

func isValidStepField(field string, min, max int) (*bool,error) {
	parts := strings.Split(field, "/")
	if len(parts) != 2 {
		return utils.BoolPtr(false), nil
	}
	isValid, err := isValidField(parts[0], min, max)
	if err != nil {
		return nil, err
	}
	isValidNum, err := isValidNumber(parts[1], min, max)
	if err != nil {
		return nil, err
	}
	return utils.BoolPtr(*isValid && *isValidNum), nil
}

func isValidRangeField(field string, min, max int)(*bool,error){
	parts := strings.Split(field, "-")
	if len(parts) != 2 {
		return utils.BoolPtr(false), nil
	}
	isValidStartRange, err := isValidNumber(parts[0], min, max)
	if err != nil {
		return nil, err
	}
	isValidEndRange, err := isValidNumber(parts[1], min, max)
	if err != nil {
		return nil, err
	}
	return utils.BoolPtr(*isValidStartRange && *isValidEndRange), nil
}

func isValidListField(field string, min, max int)(*bool,error){
	parts := strings.Split(field, ",")
		for _, part := range parts {
			isValidNum,err := isValidNumber(part, min, max)
			if err != nil {
				return nil, err
			}
			if !*isValidNum {
				return utils.BoolPtr(false), nil
			}
		}
		return utils.BoolPtr(true), nil
}

// isValidNumber checks if the string is a valid number within the range
func isValidNumber(str string, min, max int) (*bool, error) {
	value, err := strconv.Atoi(str)
	if err != nil {
		return nil, errors.New(customErr.ErrInvalidFieldValue)
	}
	return utils.BoolPtr(value >= min && value <= max),nil
}
