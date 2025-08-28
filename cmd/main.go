package main

import (
	"fmt"
	"os"
	"time"

	"top-spenders/internal/csv"
	"top-spenders/internal/spenders"
	"top-spenders/internal/transactions"
)

type Transaction = transactions.Transaction

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

	txnList := processor.Transactions

	targetMonth := time.January
	targetYear := 2020
	// Convert slice of values to slice of pointers
	var transactionPtrs []*Transaction
	for i := range txnList {
		transactionPtrs = append(transactionPtrs, &txnList[i])
	}
	topSpenders := spenders.AggregateTopSpenders(transactionPtrs, targetMonth, targetYear)
	for i, spender := range topSpenders {
		fmt.Printf("Rank %d: %s %s (%s) - £%.2f total spent\n",
			i+1, spender.FirstName, spender.LastName, spender.Email, spender.TotalSpent)
	}
}
