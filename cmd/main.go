package main

import (
	"fmt"
	"os"

	"top-spenders/internal/csv"
	"top-spenders/internal/csv/transactions"
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

	processor := &transactions.TransactionsProcessor{}
	err = csv.ReadCSV(file, processor)
	if err != nil {
		fmt.Printf("Error processing file(/s) %v", err)
	}
}
