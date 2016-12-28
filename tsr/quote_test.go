package tsr

import (
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

type testCase struct {
	name   string
	start  decimal.Decimal
	end    decimal.Decimal
	events []EventHandler
	want   decimal.Decimal
	err    error
}

func NewTestCase(name string, start float64, end float64, events []EventHandler, want float64, err error) testCase {
	return testCase{
		name:   name,
		start:  decimal.NewFromFloat(start),
		end:    decimal.NewFromFloat(end),
		events: events,
		want:   decimal.NewFromFloat(want),
		err:    err,
	}
}

type MockDoublerEventHandler struct {}

func (e MockDoublerEventHandler) Adjust(amount decimal.Decimal) decimal.Decimal {
	return amount.Mul(decimal.NewFromFloat(2.0))
}

func TestCalculateTSR(t *testing.T) {
	var tests = []testCase{
		NewTestCase("Divide by zero", 0, 1, nil, 0, ArgumentError{}),
		NewTestCase("Zero return", 2, 0, nil, 0, nil),
		NewTestCase("No events", 2, 3, nil, 1.5, nil),
		NewTestCase("With events", 2, 2, []EventHandler{
			MockDoublerEventHandler{}, MockDoublerEventHandler{},
		}, 4, nil),
	}

	for _, testcase := range tests {
		got, err := CalculateReturn(testcase.start, testcase.end, testcase.events)

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

