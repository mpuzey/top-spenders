package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"top-spenders/internal/csv"
	"top-spenders/internal/spenders"
	"top-spenders/internal/transactions"
)

type Transaction = transactions.Transaction

type Config struct {
	Month    time.Month
	Year     int
	FileName string
}

func main() {

	config := parseFlags()

	file, err := os.Open(config.FileName)
	if err != nil {
		fmt.Printf("Error opening CSV file \"%s\" :%s", config.FileName, err)
		return
	}
	defer file.Close()

	processor := &transactions.TransactionsProcessor{}
	err = csv.ReadCSV(file, processor)
	if err != nil {
		fmt.Printf("Error processing file(/s) %v", err)
	}

	txnList := processor.Transactions

	targetMonth := config.Month
	targetYear := config.Year
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

func parseFlags() *Config {
	monthPtr := flag.Int("month", 1, "an int representing the month e.g. 1 (January), 11 (Novemeber)")
	yearPtr := flag.Int("year", 2020, "an int")
	fileNamePtr := flag.String("file", "sample-transactions.csv", "a string")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("month:", *monthPtr)
	fmt.Println("year:", *yearPtr)
	fmt.Println("filename:", *fileNamePtr)
	return &Config{
		Month:    time.Month(*monthPtr),
		Year:     *yearPtr,
		FileName: *fileNamePtr,
	}
}
