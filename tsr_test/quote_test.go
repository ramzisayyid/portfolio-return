package tsr_test

import (
	"github.com/ramzisayyid/portfolio-return/tsr"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

type testCase struct {
	name   string
	start  decimal.Decimal
	end    decimal.Decimal
	events []tsr.EventHandler
	want   decimal.Decimal
	err    error
}

func TestCalculateTSR(t *testing.T) {
	var tests = []testCase{
		{"Divide by zero", decimal.NewFromFloat(0.0), decimal.NewFromFloat(1.0), nil, decimal.NewFromFloat(0), tsr.ArgumentError{}},
		{"Zero return", decimal.NewFromFloat(2.0), decimal.NewFromFloat(0.0), nil, decimal.NewFromFloat(0.0), nil},
		{"No events", decimal.NewFromFloat(2.0), decimal.NewFromFloat(3.0), nil, decimal.NewFromFloat(1.5), nil},
	}

	for _, testcase := range tests {
		got, err := tsr.CalculateReturn(testcase.start, testcase.end, testcase.events)

		if goterr, wanterr := reflect.TypeOf(err), reflect.TypeOf(testcase.err); goterr != wanterr {
			t.Errorf("For instruction %q, unexpected error: %q. Wanted %q",
				testcase.name, goterr, wanterr)
		}
		// check result type
		if want := testcase.want; got.Cmp(want) != 0 {
			t.Errorf("For instruction %q, got %q, want %q.",
				testcase.name, got, want)
		}
	}
}

func TestSplitEvent(t *testing.T) {
	want := decimal.NewFromFloat(4.5)

	event := tsr.SplitEventHandler{
		Nominator:   decimal.NewFromFloat(3),
		Denominator: decimal.NewFromFloat(2),
	}

	got := event.Adjust(decimal.NewFromFloat(3))

	if want.Cmp(got) != 0 {
		t.Errorf("For instruction %q, got %q, want %q.",
			"TestSplitEvent", got, want)
	}
}

func TestDividendEvent(t *testing.T) {
	want := decimal.NewFromFloat(3.005)

	event := tsr.DividendEventHandler{
		Quote:    decimal.NewFromFloat(30),
		Dividend: decimal.NewFromFloat(0.05),
	}

	got := event.Adjust(decimal.NewFromFloat(3))

	if want.Cmp(got) != 0 {
		t.Errorf("For instruction %q, got %q, want %q.",
			"DividendSplitEvent", got, want)
	}
}
