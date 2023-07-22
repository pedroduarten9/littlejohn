//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config generation-config.yaml ../../littlejohn.yaml

package api

import (
	"fmt"
	"littlejohn/internal/domain"
	"net/http"
	"strings"

	"github.com/benbjohnson/clock"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type API struct {
	clock  clock.Clock
	logger *zap.Logger
	days   int
}

func New(clock clock.Clock, logger *zap.Logger, days int) ServerInterface {
	return &API{
		clock:  clock,
		logger: logger,
		days:   days,
	}
}

func (a API) Tickers(ctx echo.Context) error {
	stocks := domain.GenerateStocks(ctx.Get(userKey).(string), a.clock.Now())

	return ctx.JSON(http.StatusOK, convertStocks(stocks))
}

func (a API) StockHistory(ctx echo.Context, stock StockPath) error {
	ticker := domain.Ticker(strings.ToUpper(string(stock)))

	if !domain.ExistsTicker(ticker) {
		err := NotFoundError{msg: fmt.Sprintf("Ticker %s not found", stock)}
		a.logger.Error(err.Error())
		return err
	}

	stockPrices := domain.GenerateStockPrices(a.clock.Now(), ticker, a.days)
	return ctx.JSON(http.StatusOK, convertStockPrices(stockPrices))
}

func convertStocks(domainStocks []domain.Stock) []Stock {
	stocks := make([]Stock, len(domainStocks))
	for i, domainStock := range domainStocks {
		stocks[i] = Stock{
			Price:  domainStock.Price.String(),
			Ticker: domainStock.Ticker,
		}
	}

	return stocks
}

func convertStockPrices(domainStockPrices []domain.StockPrice) []StockPrice {
	stockPrices := make([]StockPrice, len(domainStockPrices))
	for i, domainStockPrice := range domainStockPrices {
		stockPrices[i] = StockPrice{
			Date:  openapi_types.Date{domainStockPrice.Date},
			Price: domainStockPrice.Price.String(),
		}
	}

	return stockPrices
}
