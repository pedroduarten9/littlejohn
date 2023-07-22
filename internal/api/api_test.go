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
	api := New(mockClock)

	//Assert
	assert.Implements(t, (*ServerInterface)(nil), api)

	concreteAPI := api.(*API)
	assert.Equal(t, concreteAPI.clock, mockClock)
}

func TestTickers(t *testing.T) {
	// Arrange
	username := "a"
	mockClock := clock.NewMock()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tickers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(userKey, username)
	s := ServerInterfaceWrapper{
		Handler: New(mockClock),
	}

	tickersJSON := `[
		{
			"price": "1976.28",
			"symbol": "GOOG"
		},
		{
			"price": "918.64",
			"symbol": "NVDA"
		},
		{
			"price": "2761.92",
			"symbol": "JNJ"
		},
		{
			"price": "2204.77",
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
