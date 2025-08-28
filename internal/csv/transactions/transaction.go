package transactions

import "time"

// Transation represents a buy or sell Gold transaction
type Transaction struct {
	FirstName    string
	LastName     string
	EmailAddress string
	Description  string // SELL GOLD or CARD SPEND
	MerchantCode string
	Amount       float64
	FromCurrency string
	ToCurrency   string
	Rate         float64
	Date         time.Time
	// Derived field for analysis
	GBPAmount float64 // normalized amount in GBP
}

func (t *Transaction) NormalizeToGBP() float64 {
	if t.FromCurrency == "GBP" {
		return t.Amount
	}
	return t.Amount / t.Rate // Convert from other currency to GBP
}
