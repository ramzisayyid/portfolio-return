package tsr

import (
	"github.com/shopspring/decimal"
)

type EventHandler interface {
	Adjust(decimal.Decimal) decimal.Decimal
}

type SplitEventHandler struct {
	Nominator   decimal.Decimal
	Denominator decimal.Decimal
}

type DividendEventHandler struct {
	Quote    decimal.Decimal
	Dividend decimal.Decimal
}

func (e SplitEventHandler) Adjust(amount decimal.Decimal) decimal.Decimal {
	return amount.Mul(e.Nominator).Div(e.Denominator)
}

func (e DividendEventHandler) Adjust(amount decimal.Decimal) decimal.Decimal {
	return amount.Mul(e.Dividend).Div(e.Quote).Add(amount)
}
