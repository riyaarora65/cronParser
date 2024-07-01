package validators

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinuteFieldValidator(t *testing.T) {
	validator := &MinuteFieldValidator{}

	validCases := []string{"*", "*/5", "0-59", "0,15,30,45", "0", "59"}
	invalidCases := []string{"60", "invalid", "*/invalid", "0-60", "0,-15,30,45"}

	for _, tc := range validCases {
		assert.NoError(t, validator.Validate(tc), "expected no error for valid case: %s", tc)
	}

	for _, tc := range invalidCases {
		assert.Error(t, validator.Validate(tc), "expected error for invalid case: %s", tc)
	}
}

func TestHourFieldValidator(t *testing.T) {
	validator := &HourFieldValidator{}

	validCases := []string{"*", "*/3", "0-23", "0,6,12,18", "0", "23"}
	invalidCases := []string{"24", "invalid", "*/invalid", "0-24", "0,-6,12,18"}

	for _, tc := range validCases {
		assert.NoError(t, validator.Validate(tc), "expected no error for valid case: %s", tc)
	}

	for _, tc := range invalidCases {
		assert.Error(t, validator.Validate(tc), "expected error for invalid case: %s", tc)
	}
}

func TestDayOfMonthFieldValidator(t *testing.T) {
	validator := &DayOfMonthFieldValidator{}

	validCases := []string{"*", "*/2", "1-31", "1,15,30", "1", "31"}
	invalidCases := []string{"0", "32", "invalid", "*/invalid", "1-32", "1,-15,30"}

	for _, tc := range validCases {
		assert.NoError(t, validator.Validate(tc), "expected no error for valid case: %s", tc)
	}

	for _, tc := range invalidCases {
		assert.Error(t, validator.Validate(tc), "expected error for invalid case: %s", tc)
	}
}

func TestMonthFieldValidator(t *testing.T) {
	validator := &MonthFieldValidator{}

	validCases := []string{"*", "*/2", "1-12", "1,6,12", "1", "12"}
	invalidCases := []string{"0", "13", "invalid", "*/invalid", "1-13", "1,-6,12"}

	for _, tc := range validCases {
		assert.NoError(t, validator.Validate(tc), "expected no error for valid case: %s", tc)
	}

	for _, tc := range invalidCases {
		assert.Error(t, validator.Validate(tc), "expected error for invalid case: %s", tc)
	}
}

func TestDayOfWeekFieldValidator(t *testing.T) {
	validator := &DayOfWeekFieldValidator{}

	validCases := []string{"*", "*/2", "1-7", "1,3,5,7", "1", "7"}
	invalidCases := []string{"0", "8", "invalid", "*/invalid", "1-8", "1,-3,5,7"}

	for _, tc := range validCases {
		assert.NoError(t, validator.Validate(tc), "expected no error for valid case: %s", tc)
	}

	for _, tc := range invalidCases {
		assert.Error(t, validator.Validate(tc), "expected error for invalid case: %s", tc)
	}
}