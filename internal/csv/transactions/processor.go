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

type TransactionsProcessor struct {
	ProcessedRecords []Transaction
	mu               sync.Mutex
}

// Process
func (m *TransactionsProcessor) Process(record []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	transaction, err := m.parse(record)
	if err != nil {
		return err
	}

	m.ProcessedRecords = append(m.ProcessedRecords, transaction)
	return nil
}

func (*TransactionsProcessor) parse(record []string) (Transaction, error) {
	if record[0] == "First Name" {
		return Transaction{}, nil
	}

	amount, err := strconv.ParseFloat(record[5], 64)
	if err != nil {
		return Transaction{}, ErrValidation
	}

	rate, err := strconv.ParseFloat(record[8], 64)
	if err != nil {
		return Transaction{}, ErrValidation
	}

	// Parse using reference time layout
	dateStr := "12/05/2020 08:22"
	date, err := time.Parse("02/01/2006 15:04", dateStr)
	if err != nil {
		return Transaction{}, ErrValidation
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
	return transaction, nil
}
