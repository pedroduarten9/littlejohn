package domain

import (
	"hash/fnv"
	"math/rand"
	"time"
)

const maxTickersPerUser = 10

func GenerateStocks(username string, date time.Time) []Stock {
	formattedDate := date.Format(time.DateOnly)
	tickers := generateTickers(username)

	stocks := make([]Stock, len(tickers))
	for i, ticker := range tickers {
		stocks[i] = Stock{
			Symbol: string(ticker),
			Price:  generatePrice(formattedDate, ticker),
		}
	}
	return stocks
}

func generateTickers(username string) []Ticker {
	seed := generateSeedFromString(username)
	rand.Seed(seed)
	numberOfTickers := rand.Intn(maxTickersPerUser) + 1

	tickers := make([]Ticker, numberOfTickers)
	selectedTickers := map[int]bool{}
	for i := 0; i < numberOfTickers; i++ {
		tickerIdx := generateTickerIdx(selectedTickers)
		tickers[i] = Tickers[tickerIdx]
		selectedTickers[tickerIdx] = true
	}

	return tickers
}

func generateTickerIdx(selectedTickers map[int]bool) int {
	tickerIdx := rand.Intn(len(Tickers))
	for selectedTickers[tickerIdx] {
		tickerIdx = rand.Intn(len(Tickers))
	}

	return tickerIdx
}

func generatePrice(date string, ticker Ticker) Currency {
	seed := generateSeedFromString(date + string(ticker))
	rand.Seed(seed)

	return Currency{
		units: rand.Int63n(2900) + 100,
		cents: rand.Int31n(100),
	}
}

func generateSeedFromString(input string) int64 {
	hash := fnv.New64a()
	hash.Write([]byte(input))
	seed := int64(hash.Sum64())
	return seed
}
