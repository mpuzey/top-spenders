package spenders

import "time"

// Spender defines a user spending money
type Spender struct {
	FirstName        string
	LastName         string
	Email            string
	TotalSpent       float64 // in GBP
	TransactionCount int
	FirstTransaction time.Time
	LastTransaction  time.Time
	// Marketing-relevant fields
	AverageSpend float64
	SpendingDays int // unique days they spent money
}
