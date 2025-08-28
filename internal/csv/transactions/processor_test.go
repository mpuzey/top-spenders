package transactions_test

import (
	"reflect"
	"testing"
	"time"
	"top-spenders/internal/csv/transactions"
)

func TestProcess_ValidRecord(t *testing.T) {
	date, _ := time.Parse("02/01/2006 15:04", "12/05/2020 08:22")

	record := []string{
		"Niyah",
		"Singleton",
		"niyah.singleton@mailinator.com",
		"CARD SPEND",
		"5462",
		"68.5819483",
		"GBP",
		"GBP",
		"1",
		"12/05/2020 08:22",
	}

	expectedTransaction := transactions.Transaction{
		FirstName:    "Niyah",
		LastName:     "Singleton",
		EmailAddress: "niyah.singleton@mailinator.com",
		Description:  "CARD SPEND",
		MerchantCode: "5462",
		Amount:       68.5819483,
		FromCurrency: "GBP",
		ToCurrency:   "GBP",
		Rate:         1.0,
		Date:         date,
		GBPAmount:    68.5819483, // Same as Amount since currency is already GBP
	}

	processor := &transactions.TransactionsProcessor{}

	err := processor.Process(record)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	transaction := processor.ProcessedRecords[0]
	if !reflect.DeepEqual(transaction, expectedTransaction) {
		t.Errorf("transaction mismatch")
	}
}
