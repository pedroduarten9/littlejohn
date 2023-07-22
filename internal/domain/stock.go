package domain

import (
	"fmt"
	"time"
)

type Currency struct {
	units int64
	cents int32
}

func (c Currency) String() string {
	return fmt.Sprintf("%d.%d", c.units, c.cents)
}

type Stock struct {
	Ticker string
	Price  Currency
}

type StockPrice struct {
	Date  time.Time
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

func ExistsTicker(ticker Ticker) bool {
	for _, t := range Tickers {
		if t == ticker {
			return true
		}
	}
	return false
}
