package factory

import (
	cronParsers "mygolearning/deliverooProject/app/services/parsers"
	"mygolearning/deliverooProject/app/services/validators"
	"strings"
)

// CronParserInterface defines the interface for cron parsers
type CronParserInterface interface {
	ExpandField() ([]int, error)
}

// CronFactoryInterface defines methods for the cron factory
type CronFactoryInterface interface {
	CreateParser(fields []string) ([]CronParserInterface, error)
}

// CronFactory implements the CronFactoryInterface
type CronFactory struct{}

// NewCronFactory initializes a new CronFactory instance
func NewCronFactory() *CronFactory {
	return &CronFactory{}
}

// getValidator returns the appropriate validator for the given field
func (cf *CronFactory) getValidator(index int) validators.FieldValidator {
	switch index {
	case 0:
		return &validators.MinuteFieldValidator{}
	case 1:
		return &validators.HourFieldValidator{}
	case 2:
		return &validators.DayOfMonthFieldValidator{}
	case 3:
		return &validators.MonthFieldValidator{}
	case 4:
		return &validators.DayOfWeekFieldValidator{}
	}
	return nil
}

// CreateParser creates a new CronParser based on the field string
func (cf *CronFactory) CreateParser(fields []string) ([]CronParserInterface, error) {
	var parsers []CronParserInterface

	ranges := []struct {
		min, max int
	}{
		{0, 59}, // minute
		{0, 23}, // hour
		{1, 31}, // day of month
		{1, 12}, // month
		{1, 7},  // day of week (0=Sunday to 6=Saturday)
	}

	for i, field := range fields[:5] {
		validator := cf.getValidator(i)
		if err := validator.Validate(field); err != nil {
			return nil, err
		}
		switch {
		case field == "*":
			parsers = append(parsers, cronParsers.NewStarParser(ranges[i].min, ranges[i].max))
		case strings.Contains(field, "/"):
			parser, err := cronParsers.NewStepParser(field, ranges[i].min, ranges[i].max)
			if err != nil {
				return nil, err
			}
			parsers = append(parsers, parser)
		case strings.Contains(field, "-"):
			parser, err := cronParsers.NewRangeParser(field)
			if err != nil {
				return nil, err
			}
			parsers = append(parsers, parser)
		case strings.Contains(field, ","):
			parser, err := cronParsers.NewListParser(field)
			if err != nil {
				return nil, err
			}
			parsers = append(parsers, parser)
		default:
			parser, err := cronParsers.NewListParser(field) // For single values
			if err != nil {
				return nil, err
			}
			parsers = append(parsers, parser)
		}
	}
	return parsers, nil
}
