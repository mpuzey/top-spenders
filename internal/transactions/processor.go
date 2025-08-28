package transactions

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	ErrValidation = fmt.Errorf("validation error")
)

// TransactionsProcessor stores a slice of parsed transactions and a mutex for safety across goroutines during parsing
type TransactionsProcessor struct {
	Transactions []Transaction
	mu           sync.Mutex
}

// Process parses a single transaction record and adds it to the Transactions slice on the TransactionsProcessor
func (m *TransactionsProcessor) Process(record []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	err := m.parse(record)
	if err != nil {
		return err
	}

	return nil
}

func (m *TransactionsProcessor) parse(record []string) error {
	if record[0] == "First name" {
		return nil
	}

	amount, err := strconv.ParseFloat(record[5], 64)
	if err != nil {
		return ErrValidation
	}

	rate, err := strconv.ParseFloat(record[8], 64)
	if err != nil {
		return ErrValidation
	}

	// Parse date from record
	date, err := time.Parse("02/01/2006 15:04", record[9])
	if err != nil {
		return ErrValidation
	}

	transaction := Transaction{
		FirstName:    record[0],
		LastName:     record[1],
		EmailAddress: record[2],
		Description:  record[3],
		MerchantCode: record[4],
		Amount:       amount,
		FromCurrency: record[6],
		ToCurrency:   record[7],
		Rate:         rate,
		Date:         date,
	}

	// Calculate GBP amount
	transaction.GBPAmount = transaction.NormalizeToGBP()

	m.Transactions = append(m.Transactions, transaction)
	return nil
}
