package spenders

import (
	"sort"
	"time"

	"github.com/matt/top-spenders/internal/transactions"
)

// Transaction is an alias which resolves import issues with the transactions type
type Transaction = transactions.Transaction

// AggregateTopSpenders gets the top five motnhly spenders in the transaction list
func AggregateTopSpenders(transactions []*Transaction, targetMonth time.Month, targetYear int) []*Spender {
	cardSpends := filterCardSpends(transactions, targetMonth, targetYear)

	for _, tx := range cardSpends {
		tx.GBPAmount = tx.NormalizeToGBP()
	}

	spenders := aggregateByUser(cardSpends)

	// Sort by total spend (descending)
	sort.Slice(spenders, func(i, j int) bool {
		return spenders[i].TotalSpent > spenders[j].TotalSpent
	})

	topCount := 5
	if len(spenders) < topCount {
		topCount = len(spenders)
	}

	topSpenders := spenders[:topCount]

	return topSpenders
}

func filterCardSpends(transactions []*Transaction, targetMonth time.Month, targetYear int) []*Transaction {
	var cardSpends []*Transaction

	for _, tx := range transactions {
		if tx.Description != "CARD SPEND" {
			continue
		}

		if tx.Date.Month() == targetMonth && tx.Date.Year() == targetYear {
			cardSpends = append(cardSpends, tx)
		}
	}

	return cardSpends
}

func aggregateByUser(cardSpends []*Transaction) []*Spender {
	userSpends := make(map[string]*Spender)

	for _, tx := range cardSpends {
		email := tx.EmailAddress

		if spender, exists := userSpends[email]; exists {
			// Update existing spender
			spender.TotalSpent += tx.GBPAmount
			spender.TransactionCount++

			// Update date ranges
			if tx.Date.Before(spender.FirstTransaction) {
				spender.FirstTransaction = tx.Date
			}
			if tx.Date.After(spender.LastTransaction) {
				spender.LastTransaction = tx.Date
			}
		} else {
			// Create new spender
			userSpends[email] = &Spender{
				FirstName:        tx.FirstName,
				LastName:         tx.LastName,
				Email:            email,
				TotalSpent:       tx.GBPAmount,
				TransactionCount: 1,
				FirstTransaction: tx.Date,
				LastTransaction:  tx.Date,
			}
		}
	}

	// Convert map to slice and calculate derived fields
	var spenders []*Spender
	for _, spender := range userSpends {
		spender.AverageSpend = spender.TotalSpent / float64(spender.TransactionCount)

		// Calculate unique spending days - u
		spender.SpendingDays = calculateSpendingDays(cardSpends, spender.Email)

		spenders = append(spenders, spender)
	}

	return spenders
}

// Helper function to calculate unique spending days for a user, used email address as the unique identifier
func calculateSpendingDays(transactions []*Transaction, email string) int {
	uniqueDays := make(map[string]bool)

	for _, tx := range transactions {
		if tx.EmailAddress == email {
			// Use date string (YYYY-MM-DD) as key to count unique days
			dayKey := tx.Date.Format("2006-01-02")
			uniqueDays[dayKey] = true
		}
	}

	return len(uniqueDays)
}
