package domain

import "fmt"

type Currency struct {
	units int64
	cents int32
}

func (c Currency) String() string {
	return fmt.Sprintf("%d.%d", c.units, c.cents)
}

type Stock struct {
	Symbol string
	Price  Currency
}

type StockPrice struct {
	Date  string
	Price Currency
}

type Ticker string

var Tickers = []Ticker{
	"AAPL",
	"MSFT",
	"GOOG",
	"AMZN",
	"FB",
	"TSLA",
	"NVDA",
	"JPM",
	"BABA",
	"JNJ",
	"WMT",
	"PG",
	"PYPL",
	"DIS",
	"ADBE",
	"PFE",
	"V",
	"MA",
	"CRM",
	"NFLX",
}
