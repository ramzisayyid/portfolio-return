package tsr

import (
	"github.com/shopspring/decimal"
)

// CalculateReturn calculates the return on a specific stock given the purchase date, the sell date
// and a list of events that happened while holding the stock.
func CalculateReturn(start decimal.Decimal, end decimal.Decimal, events []EventHandler) (decimal.Decimal, error) {
	if start.Cmp(decimal.NewFromFloat(0)) == 0 {
		return decimal.NewFromFloat(0), ArgumentError{"start"}
	}

	units := decimal.NewFromFloat(1).Div(start)

	for _, e := range events {
		units = e.Adjust(units)
	}

	return units.Mul(end), nil
}
