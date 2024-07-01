package services

import (
	"fmt"
	"strings"
	"mygolearning/deliverooProject/app/services/factory"
	customErr "mygolearning/deliverooProject/app/errors"
)

// CronServiceInterface defines methods for the cron service
type CronServiceInterface interface {
	ParseAndPrint(cronString string)
}

// CronService implements the CronServiceInterface
type CronService struct {
	fields []string
}

// NewCronService initializes a new CronService instance
func NewCronService() *CronService {
	return &CronService{}
}

// ParseAndPrint parses the cron string and prints the expanded fields
func (cs *CronService) ParseAndPrint(cronString string) {
	cs.fields = strings.Fields(cronString)
	if len(cs.fields) < 6 {
		fmt.Println(customErr.ErrInvalidFieldFormat)
		return
	}

	fieldNames := []string{"minute", "hour", "day of month", "month", "day of week"}

	cronFactory := factory.NewCronFactory()
	parsers, err := cronFactory.CreateParser(cs.fields[:5])
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, parser := range parsers {
		expanded, err := parser.ExpandField()
		if err != nil {
			fmt.Printf("%-14s %s\n", fieldNames[i], err.Error())
		} else {
			fmt.Printf("%-14s %s\n", fieldNames[i], strings.Trim(fmt.Sprint(expanded), "[]"))
		}
	}

	command := strings.Join(cs.fields[5:], " ")
	fmt.Printf("%-14s %s\n", "command", command)
}
