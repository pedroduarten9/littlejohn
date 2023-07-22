//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config generation-config.yaml ../../littlejohn.yaml

package api

import "github.com/labstack/echo/v4"

type API struct {
}

func NewAPI() ServerInterface {
	return &API{}
}

func (a API) Tickers(ctx echo.Context) error {
	return nil
}

func (a API) StockHistory(ctx echo.Context, stock StockPath) error {
	return nil
}
