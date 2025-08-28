package main

import (
	"fmt"
	"os"

	"top-spenders/internal/csv"
)

func main() {
	csvName := os.Getenv("CSV_FILE")
	if len(csvName) == 0 {
		csvName = "sample-transactions.csv"
	}

	file, err := os.Open(csvName)
	if err != nil {
		fmt.Printf("Error opening CSV file \"%s\" :%s", csvName, err)
		return
	}
	defer file.Close()

	processor := &csv.TransactionsCSVProcessor{}
	err = csv.ProcessCSV(file, processor)
	if err != nil {
		fmt.Printf("Error processing file(/s) %v", err)
	}
}
