package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	mockClock := clock.NewMock()
	days := 90
	api := New(mockClock, days)

	//Assert
	assert.Implements(t, (*ServerInterface)(nil), api)

	concreteAPI := api.(*API)
	assert.Equal(t, mockClock, concreteAPI.clock)
	assert.Equal(t, days, concreteAPI.days)
}

func TestTickers(t *testing.T) {
	// Arrange
	username := "a"
	days := 90
	mockClock := clock.NewMock()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tickers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(userKey, username)
	s := ServerInterfaceWrapper{
		Handler: New(mockClock, days),
	}

	tickersJSON := `[
		{
			"price": "366.34",
			"symbol": "GOOG"
		},
		{
			"price": "2965.49",
			"symbol": "NVDA"
		},
		{
			"price": "1682.45",
			"symbol": "JNJ"
		},
		{
			"price": "2179.59",
			"symbol": "AMZN"
		}
	]`

	// Act
	err := s.Tickers(c)

	// Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, tickersJSON, rec.Body.String())
	}
}

func TestStockHistory(t *testing.T) {
	// Arrange
	username := "a"
	days := 4
	stock := "GOOG"
	mockClock := clock.NewMock()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tickers/:stock/history", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(userKey, username)
	c.SetParamNames("stock")
	c.SetParamValues(stock)
	s := ServerInterfaceWrapper{
		Handler: New(mockClock, days),
	}

	stockHistoryJSON := `[
		{
			"price": "366.34",
			"date": "1970-01-01"
		},
		{
			"price": "446.42",
			"date": "1969-12-31"
		},
		{
			"price": "2605.32",
			"date": "1969-12-30"
		},
		{
			"price": "788.67",
			"date": "1969-12-29"
		}
	]`

	// Act
	err := s.StockHistory(c)

	// Assert
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, stockHistoryJSON, rec.Body.String())
	}

}

func TestStockHistory_NotFound(t *testing.T) {
	// Arrange
	username := "a"
	days := 4
	stock := "efnwofnwjflw"
	mockClock := clock.NewMock()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tickers/:stock/history", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(userKey, username)
	c.SetParamNames("stock")
	c.SetParamValues(stock)
	s := ServerInterfaceWrapper{
		Handler: New(mockClock, days),
	}

	expectedError := NotFoundError{
		msg: "Ticker " + stock + " not found",
	}
	// Act
	err := s.StockHistory(c)

	// Assert
	assert.Equal(t, expectedError, err)
}
