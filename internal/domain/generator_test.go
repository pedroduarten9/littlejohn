package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateStocks_SameDay_SameUser(t *testing.T) {
	// Arrange
	username := "username"
	date := time.Date(2023, 3, 3, 3, 3, 12, 0, time.UTC)

	// Act
	stocksDay1 := GenerateStocks(username, date)
	stocksDay2 := GenerateStocks(username, date)

	// Assert
	assert.ElementsMatch(t, stocksDay1, stocksDay2)
}

func TestGenerateStocks_SameDay_DifferentUser(t *testing.T) {
	// Arrange
	username := "username"
	username2 := "username-2"
	date := time.Date(2023, 3, 3, 3, 3, 12, 0, time.UTC)

	// Act
	stocksDay1 := GenerateStocks(username, date)
	stocksDay2 := GenerateStocks(username2, date)

	// Assert
	assert.False(t, reflect.DeepEqual(stocksDay1, stocksDay2))
}

func TestGenerateStocks_DifferentDay_SameUser(t *testing.T) {
	// Arrange
	username := "username"
	date := time.Date(2023, 3, 3, 3, 3, 12, 0, time.UTC)
	date2 := time.Date(2023, 3, 2, 3, 3, 12, 0, time.UTC)

	// Act
	stocksDay1 := GenerateStocks(username, date)
	stocksDay2 := GenerateStocks(username, date2)

	// Assert
	assert.False(t, reflect.DeepEqual(stocksDay1, stocksDay2))

}

func TestGenerateStocks_DifferentDay_DifferentUser(t *testing.T) {
	// Arrange
	username := "username"
	username2 := "username-2"
	date := time.Date(2023, 3, 3, 3, 3, 12, 0, time.UTC)
	date2 := time.Date(2023, 3, 2, 3, 3, 12, 0, time.UTC)

	// Act
	stocksDay1 := GenerateStocks(username, date)
	stocksDay2 := GenerateStocks(username2, date2)

	// Assert
	assert.False(t, reflect.DeepEqual(stocksDay1, stocksDay2))

}

func TestGenerateStockPrices(t *testing.T) {
	// Arrange
	ticker := Ticker("CRM")
	date := time.Date(2023, 3, 3, 3, 3, 12, 0, time.UTC)
	days := 2

	expectedStockPrices := []StockPrice{
		{
			Date: "2023-03-03",
			Price: Currency{
				units: int64(1636),
				cents: int32(9),
			},
		},
		{
			Date: "2023-03-02",
			Price: Currency{
				units: int64(1276),
				cents: int32(16),
			},
		},
	}
	// Act
	stockPrices := GenerateStockPrices(date, ticker, days)

	// Assert
	assert.ElementsMatch(t, expectedStockPrices, stockPrices)
}
