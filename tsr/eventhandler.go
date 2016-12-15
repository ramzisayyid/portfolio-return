package tsr

import (
	"github.com/shopspring/decimal"
)

// EventHandler is an interface used to abstract adjustments to stock value
// after a dividend or stock
type EventHandler interface {
	Adjust(decimal.Decimal) decimal.Decimal
}

// SplitEventHandler adjusts the number of stock units after a split.
type SplitEventHandler struct {
	Take decimal.Decimal
	Give decimal.Decimal
}

// DividendEventHandler adjusts the number of units after a dividend distribution.
// It assumes that all dividends were re-invested in the stock at the price provided.
type DividendEventHandler struct {
	Quote    decimal.Decimal
	Dividend decimal.Decimal
}

func (e SplitEventHandler) Adjust(amount decimal.Decimal) decimal.Decimal {
	return amount.Mul(e.Give).Div(e.Take)
}

func (e DividendEventHandler) Adjust(amount decimal.Decimal) decimal.Decimal {
	return amount.Mul(e.Dividend).Div(e.Quote).Add(amount)
}
