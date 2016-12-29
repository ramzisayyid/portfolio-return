package tsr

import (
	"errors"
	"github.com/FlashBoys/go-finance"
	"github.com/shopspring/decimal"
	"time"
)

// YahooService implements the QuoteService interface using the go-finance package.
// Given how simple this service is, and the fact that it is mainly a wrapper around an
// external interface, I did not include any unit tests.
type YahooService struct{}

func (y *YahooService) GetQuote(symbol string, date time.Time) (decimal.Decimal, error) {
	bars, err := finance.GetQuoteHistory(symbol, time.Now(), time.Now(), finance.IntervalDaily)

	if err != nil {
		return decimal.Zero, err
	}

	if len(bars) == 0 {
		return decimal.Zero, ArgumentError{"date"}
	}

	for _, b := range bars {
		if b.Date.Truncate(24 * time.Hour).Equal(date.Truncate(24 * time.Hour)) {
			return b.Close, nil
		}
	}

	return decimal.Zero, errors.New("date not found in result set")
}

func (y *YahooService) GetLatestQuote(symbol string) (decimal.Decimal, error) {
	q, err := finance.GetQuote(symbol)

	if err != nil {
		return decimal.Zero, err
	}

	return q.LastTradePrice, nil
}

func (y *YahooService) GetAverageQuote(symbol string, start time.Time, end time.Time) (decimal.Decimal, error) {
	bars, err := finance.GetQuoteHistory(symbol, time.Now(), time.Now(), finance.IntervalDaily)

	if err != nil {
		return decimal.Zero, err
	}

	var avg decimal.Decimal
	var days int

	for _, p := range bars {
		if p.Date.After(start.AddDate(0, 0, -1)) && p.Date.Before(end) {
			avg = avg.Add(p.Close)
			days += 1
		}
	}

	return avg.Div(decimal.NewFromFloat(float64(days))), nil
}
