// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
)

const (
	BasicAuthScopes = "basicAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Stock defines model for Stock.
type Stock struct {
	Price  string `json:"price"`
	Symbol string `json:"symbol"`
}

// StockPrice defines model for StockPrice.
type StockPrice struct {
	Date  openapi_types.Date `json:"date"`
	Price string             `json:"price"`
}

// StockPrices defines model for StockPrices.
type StockPrices = interface{}

// Stocks defines model for Stocks.
type Stocks = interface{}

// StockPath defines model for StockPath.
type StockPath = string

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /tickers)
	Tickers(ctx echo.Context) error

	// (GET /tickers/{stock}/history)
	StockHistory(ctx echo.Context, stock StockPath) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Tickers converts echo context to params.
func (w *ServerInterfaceWrapper) Tickers(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Tickers(ctx)
	return err
}

// StockHistory converts echo context to params.
func (w *ServerInterfaceWrapper) StockHistory(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "stock" -------------
	var stock StockPath

	err = runtime.BindStyledParameterWithLocation("simple", false, "stock", runtime.ParamLocationPath, ctx.Param("stock"), &stock)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter stock: %s", err))
	}

	ctx.Set(BasicAuthScopes, []string{})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StockHistory(ctx, stock)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/tickers", wrapper.Tickers)
	router.GET(baseURL+"/tickers/:stock/history", wrapper.StockHistory)

}