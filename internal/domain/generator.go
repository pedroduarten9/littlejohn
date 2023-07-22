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
			Ticker: string(ticker),
			Price:  generateCurrency(formattedDate, ticker),
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

func GenerateStockPrices(date time.Time, ticker Ticker, days int) []StockPrice {
	stockPrices := make([]StockPrice, days)
	for i := 0; i < days; i++ {
		stockDate := date.Add(time.Duration(-i) * time.Hour * 24)
		stockPrices[i] = generateStockPrice(stockDate, ticker)
	}

	return stockPrices
}

func generateStockPrice(date time.Time, ticker Ticker) StockPrice {
	formattedDate := date.Format(time.DateOnly)
	seed := generateSeedFromString(formattedDate + string(ticker))
	rand.Seed(seed)

	return StockPrice{
		Date:  date,
		Price: generateCurrency(formattedDate, ticker),
	}
}

func generateCurrency(date string, ticker Ticker) Currency {
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
