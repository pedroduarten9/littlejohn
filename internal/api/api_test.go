package api

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	// Act
	mockClock := clock.NewMock()
	days := 90
	logger, _ := zap.NewDevelopment()
	api := New(mockClock, logger, days)

	//Assert
	assert.Implements(t, (*ServerInterface)(nil), api)

	concreteAPI := api.(*API)
	assert.Equal(t, mockClock, concreteAPI.clock)
	assert.Equal(t, logger, concreteAPI.logger)
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
	logger, _ := zap.NewDevelopment()
	s := ServerInterfaceWrapper{
		Handler: New(mockClock, logger, days),
	}

	tickersJSON := `[
		{
			"price": "366.34",
			"ticker": "GOOG"
		},
		{
			"price": "2965.49",
			"ticker": "NVDA"
		},
		{
			"price": "1682.45",
			"ticker": "JNJ"
		},
		{
			"price": "2179.59",
			"ticker": "AMZN"
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
	logger, _ := zap.NewDevelopment()
	s := ServerInterfaceWrapper{
		Handler: New(mockClock, logger, days),
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

func TestStockHistory_ThirdPage(t *testing.T) {
	// Arrange
	username := "a"
	days := 4
	stock := "GOOG"
	page := "3"
	mockClock := clock.NewMock()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tickers/:stock/history", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(userKey, username)
	c.SetParamNames("stock")
	c.SetParamValues(stock)
	c.QueryParams()["page"] = []string{page}
	logger, _ := zap.NewDevelopment()
	s := ServerInterfaceWrapper{
		Handler: New(mockClock, logger, days),
	}

	stockHistoryJSON := `[
		{
			"price": "1228.33",
			"date": "1969-12-24"
		},
		{
			"price": "1391.30",
			"date": "1969-12-23"
		},
		{
			"price": "2624.49",
			"date": "1969-12-22"
		},
		{
			"price": "1540.48",
			"date": "1969-12-21"
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

func TestStockHistory_0Page(t *testing.T) {
	// Arrange
	username := "a"
	days := 4
	stock := "GOOG"
	page := "0"
	mockClock := clock.NewMock()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tickers/:stock/history", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(userKey, username)
	c.SetParamNames("stock")
	c.SetParamValues(stock)
	c.QueryParams()["page"] = []string{page}
	logger, _ := zap.NewDevelopment()
	s := ServerInterfaceWrapper{
		Handler: New(mockClock, logger, days),
	}

	expectedError := BadRequestError{
		msg: "Page " + page + " is lower than min page " + strconv.Itoa(minPage),
	}

	// Act
	err := s.StockHistory(c)

	// Assert
	assert.Equal(t, expectedError, err)

}

func TestStockHistory_MaxPage(t *testing.T) {
	// Arrange
	username := "a"
	days := 4
	stock := "GOOG"
	page := "54"
	mockClock := clock.NewMock()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tickers/:stock/history", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(userKey, username)
	c.SetParamNames("stock")
	c.SetParamValues(stock)
	c.QueryParams()["page"] = []string{page}
	logger, _ := zap.NewDevelopment()
	s := ServerInterfaceWrapper{
		Handler: New(mockClock, logger, days),
	}

	expectedError := BadRequestError{
		msg: "Page " + page + " is bigger than max page " + strconv.Itoa(maxPage),
	}

	// Act
	err := s.StockHistory(c)

	// Assert
	assert.Equal(t, expectedError, err)

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
	logger, _ := zap.NewDevelopment()
	s := ServerInterfaceWrapper{
		Handler: New(mockClock, logger, days),
	}

	expectedError := NotFoundError{
		msg: "Ticker " + stock + " not found",
	}
	// Act
	err := s.StockHistory(c)

	// Assert
	assert.Equal(t, expectedError, err)
}
