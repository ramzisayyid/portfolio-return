package tsr

import (
	"github.com/shopspring/decimal"
	"time"
)

type Finance interface {
	GetQuote(symbol string, date time.Time) (decimal.Decimal, error)
	GetLatestQuote(symbol string) (decimal.Decimal, error)
	GetAverageQuote(symbol string, start time.Time, end time.Time) (decimal.Decimal, error)
}
