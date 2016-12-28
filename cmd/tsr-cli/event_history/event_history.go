package main

import (
	"fmt"
	"os"
	"time"

	"github.com/FlashBoys/go-finance"
)

func main() {
	// Set time range from Jan 2010 up to the current date.
	// This example will return a slice of both dividends and splits.
	start, _ := time.Parse(time.RFC3339, "2013-01-01T16:00:00+00:00")
	end := time.Now()

	// Request event history for AAPL.
	events, err := finance.GetDividendSplitHistory(os.Args[1], start, end)
	if err == nil {
		for _, e := range events {
			fmt.Printf("%v on %v\n", e.EventType, e.Date)
		}
	}
}
