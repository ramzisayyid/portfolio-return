package tsr

import (
	"github.com/shopspring/decimal"
	"time"
)

// Finance is an interface to abstract quote retrieval
type Finance interface {
	// GetQuote retrieves the closing price of a stock on a given date
	GetQuote(symbol string, date time.Time) (decimal.Decimal, error)

	// GetLatestQuote retrieves the (near) real-time trading price of a stock
	GetLatestQuote(symbol string) (decimal.Decimal, error)

	// GetAverageQuote returns the average closing price of a stock between 2 given dates.
	// The calculation includes the start date but does not include the end date.
	GetAverageQuote(symbol string, start time.Time, end time.Time) (decimal.Decimal, error)
}
