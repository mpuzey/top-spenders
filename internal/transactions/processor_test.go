package transactions_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/matt/top-spenders/internal/transactions"
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

	if len(processor.Transactions) != 1 {
		t.Errorf("expected 1 transaction, got %d", len(processor.Transactions))
	}

	transaction := processor.Transactions[0]
	if !reflect.DeepEqual(transaction, expectedTransaction) {
		t.Errorf("transaction mismatch")
	}
}

func TestProcess_HeaderRow(t *testing.T) {
	record := []string{
		"First name",
		"Last name",
		"Email",
		"Description",
		"Merchant code",
		"Amount",
		"From Currency",
		"To Currency",
		"Rate",
		"Date",
	}

	processor := &transactions.TransactionsProcessor{}
	err := processor.Process(record)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(processor.Transactions) != 0 {
		t.Errorf("expected no transactions for header row, got %d", len(processor.Transactions))
	}
}

func TestProcess_GGMToGBPConversion(t *testing.T) {
	date, _ := time.Parse("02/01/2006 15:04", "02/01/2020 03:07")

	record := []string{
		"Amanda",
		"Burn",
		"amanda.burn@mailinator.com",
		"CARD SPEND",
		"5013",
		"41.8411403",
		"GBP",
		"GGM",
		"47.0892",
		"02/01/2020 03:07",
	}

	expectedTransaction := transactions.Transaction{
		FirstName:    "Amanda",
		LastName:     "Burn",
		EmailAddress: "amanda.burn@mailinator.com",
		Description:  "CARD SPEND",
		MerchantCode: "5013",
		Amount:       41.8411403,
		FromCurrency: "GBP",
		ToCurrency:   "GGM",
		Rate:         47.0892,
		Date:         date,
		GBPAmount:    41.8411403, // Same as Amount since FromCurrency is GBP
	}

	processor := &transactions.TransactionsProcessor{}
	err := processor.Process(record)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	transaction := processor.Transactions[0]
	if !reflect.DeepEqual(transaction, expectedTransaction) {
		t.Errorf("transaction mismatch")
	}
}

func TestProcess_GGMFromGBPConversion(t *testing.T) {
	date, _ := time.Parse("02/01/2006 15:04", "23/05/2020 09:20")

	record := []string{
		"Andreea",
		"Suarez",
		"andreea.suarez@mailinator.com",
		"CARD SPEND",
		"5411",
		"1.5103011",
		"GGM",
		"GBP",
		"47.7478",
		"23/05/2020 09:20",
	}

	expectedTransaction := transactions.Transaction{
		FirstName:    "Andreea",
		LastName:     "Suarez",
		EmailAddress: "andreea.suarez@mailinator.com",
		Description:  "CARD SPEND",
		MerchantCode: "5411",
		Amount:       1.5103011,
		FromCurrency: "GGM",
		ToCurrency:   "GBP",
		Rate:         47.7478,
		Date:         date,
		GBPAmount:    1.5103011 / 47.7478, // Convert from GGM to GBP
	}

	processor := &transactions.TransactionsProcessor{}
	err := processor.Process(record)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	transaction := processor.Transactions[0]
	if !reflect.DeepEqual(transaction, expectedTransaction) {
		t.Errorf("transaction mismatch")
	}
}

func TestProcess_SELLGOLDTransaction(t *testing.T) {
	date, _ := time.Parse("02/01/2006 15:04", "17/02/2020 19:17")

	record := []string{
		"Ebrahim",
		"Pickett",
		"ebrahim.pickett@mailinator.com",
		"SELL GOLD",
		"",
		"0.9928468",
		"GGM",
		"GBP",
		"47.7494",
		"17/02/2020 19:17",
	}

	expectedTransaction := transactions.Transaction{
		FirstName:    "Ebrahim",
		LastName:     "Pickett",
		EmailAddress: "ebrahim.pickett@mailinator.com",
		Description:  "SELL GOLD",
		MerchantCode: "",
		Amount:       0.9928468,
		FromCurrency: "GGM",
		ToCurrency:   "GBP",
		Rate:         47.7494,
		Date:         date,
		GBPAmount:    0.9928468 / 47.7494, // Convert from GGM to GBP
	}

	processor := &transactions.TransactionsProcessor{}
	err := processor.Process(record)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	transaction := processor.Transactions[0]
	if !reflect.DeepEqual(transaction, expectedTransaction) {
		t.Errorf("transaction mismatch")
	}
}

func TestProcess_BUYGOLDTransaction(t *testing.T) {
	date, _ := time.Parse("02/01/2006 15:04", "29/07/2020 17:32")

	record := []string{
		"Kaelan",
		"Rodriquez",
		"kaelan.rodriquez@mailinator.com",
		"BUY GOLD",
		"",
		"75.5180979",
		"GBP",
		"GGM",
		"47.7555",
		"29/07/2020 17:32",
	}

	expectedTransaction := transactions.Transaction{
		FirstName:    "Kaelan",
		LastName:     "Rodriquez",
		EmailAddress: "kaelan.rodriquez@mailinator.com",
		Description:  "BUY GOLD",
		MerchantCode: "",
		Amount:       75.5180979,
		FromCurrency: "GBP",
		ToCurrency:   "GGM",
		Rate:         47.7555,
		Date:         date,
		GBPAmount:    75.5180979, // Same as Amount since FromCurrency is GBP
	}

	processor := &transactions.TransactionsProcessor{}
	err := processor.Process(record)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	transaction := processor.Transactions[0]
	if !reflect.DeepEqual(transaction, expectedTransaction) {
		t.Errorf("transaction mismatch")
	}
}

func TestProcess_InvalidAmount(t *testing.T) {
	record := []string{
		"Test",
		"User",
		"test@example.com",
		"CARD SPEND",
		"1234",
		"invalid_amount",
		"GBP",
		"GBP",
		"1",
		"12/05/2020 08:22",
	}

	processor := &transactions.TransactionsProcessor{}
	err := processor.Process(record)

	if err == nil {
		t.Errorf("expected validation error for invalid amount")
	}
}

func TestProcess_InvalidRate(t *testing.T) {
	record := []string{
		"Test",
		"User",
		"test@example.com",
		"CARD SPEND",
		"1234",
		"100.0",
		"GBP",
		"GGM",
		"invalid_rate",
		"12/05/2020 08:22",
	}

	processor := &transactions.TransactionsProcessor{}
	err := processor.Process(record)

	if err == nil {
		t.Errorf("expected validation error for invalid rate")
	}
}

func TestProcess_InvalidDate(t *testing.T) {
	record := []string{
		"Test",
		"User",
		"test@example.com",
		"CARD SPEND",
		"1234",
		"100.0",
		"GBP",
		"GBP",
		"1",
		"invalid_date",
	}

	processor := &transactions.TransactionsProcessor{}
	err := processor.Process(record)

	if err == nil {
		t.Errorf("expected validation error for invalid date")
	}
}

func TestProcess_MultipleTransactions(t *testing.T) {
	processor := &transactions.TransactionsProcessor{}

	records := [][]string{
		{
			"Niyah", "Singleton", "niyah.singleton@mailinator.com", "CARD SPEND", "5462",
			"68.5819483", "GBP", "GBP", "1", "12/05/2020 08:22",
		},
		{
			"Amanda", "Burn", "amanda.burn@mailinator.com", "CARD SPEND", "5013",
			"41.8411403", "GBP", "GGM", "47.0892", "02/01/2020 03:07",
		},
		{
			"Andreea", "Suarez", "andreea.suarez@mailinator.com", "CARD SPEND", "5411",
			"1.5103011", "GGM", "GBP", "47.7478", "23/05/2020 09:20",
		},
	}

	for _, record := range records {
		err := processor.Process(record)
		if err != nil {
			t.Errorf("unexpected error processing record: %v", err)
		}
	}

	if len(processor.Transactions) != 3 {
		t.Errorf("expected 3 transactions, got %d", len(processor.Transactions))
	}
}
