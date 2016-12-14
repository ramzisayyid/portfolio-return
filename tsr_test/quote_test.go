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

func NewTestCase(name string, start float64, end float64, events []tsr.EventHandler, want float64, err error) testCase {
	return testCase{
		name:   name,
		start:  decimal.NewFromFloat(start),
		end:    decimal.NewFromFloat(end),
		events: events,
		want:   decimal.NewFromFloat(want),
		err:    err,
	}
}

func TestCalculateTSR(t *testing.T) {
	var tests = []testCase{
		NewTestCase("Divide by zero", 0, 1, nil, 0, tsr.ArgumentError{}),
		NewTestCase("Zero return", 2, 0, nil, 0, nil),
		NewTestCase("No events", 2, 3, nil, 1.5, nil),
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
