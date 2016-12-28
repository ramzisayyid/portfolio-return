package main

import (
	"fmt"
	"github.com/FlashBoys/go-finance"
	"os"
)

func main() {
	// 15-min delayed full quotes for Apple, Twitter, and Facebook.
	symbols := os.Args[1:]
	quotes, err := finance.GetQuotes(symbols)
	if err == nil {
		for _, q := range quotes {
			fmt.Printf("%s\t%s\n", q.Symbol, q.LastTradePrice.StringFixed(2))
		}
	}
}
