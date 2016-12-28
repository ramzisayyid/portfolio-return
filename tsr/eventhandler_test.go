package tsr

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestSplitEvent(t *testing.T) {
	want := decimal.NewFromFloat(4.5)

	event := SplitEventHandler{
		decimal.NewFromFloat(2),
		decimal.NewFromFloat(3),
	}

	got := event.Adjust(decimal.NewFromFloat(3))

	if want.Cmp(got) != 0 {
		t.Errorf("For instruction %q, got %q, want %q.",
			"TestSplitEvent", got, want)
	}
}

func TestDividendEvent(t *testing.T) {
	want := decimal.NewFromFloat(3.005)

	event := DividendEventHandler{
		Quote:    decimal.NewFromFloat(30),
		Dividend: decimal.NewFromFloat(0.05),
	}

	got := event.Adjust(decimal.NewFromFloat(3))

	if want.Cmp(got) != 0 {
		t.Errorf("For instruction %q, got %q, want %q.",
			"DividendSplitEvent", got, want)
	}
}
