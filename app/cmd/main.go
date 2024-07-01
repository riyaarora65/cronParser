package main

import (
	"fmt"
	"os"
	"github.com/riyaarora65/cron_parser/app/services"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cron-expander \"<cron_string>\"")
		return
	}

	cronString := os.Args[1]
	cronService := services.NewCronService()
	cronService.ParseAndPrint(cronString)
}
