//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config generation-config.yaml ../../littlejohn.yaml

package api

import (
	"littlejohn/internal/domain"
	"net/http"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/labstack/echo/v4"
)

type API struct {
	clock clock.Clock
}

func New(clock clock.Clock) ServerInterface {
	return &API{
		clock: clock,
	}
}

func (a API) Tickers(ctx echo.Context) error {
	stocks := domain.GenerateStocks(ctx.Get(userKey).(string), time.Now())

	return ctx.JSON(http.StatusOK, convertStocks(stocks))
}

func (a API) StockHistory(ctx echo.Context, stock StockPath) error {
	return nil
}

func convertStocks(domainStocks []domain.Stock) []Stock {
	stocks := make([]Stock, len(domainStocks))
	for i, domainStock := range domainStocks {
		stocks[i] = Stock{
			Price:  domainStock.Price.String(),
			Symbol: domainStock.Symbol,
		}
	}

	return stocks
}
