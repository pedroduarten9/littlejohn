package main

import (
	"littlejohn/internal/api"

	"github.com/benbjohnson/clock"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

const days = 90

func main() {
	e := echo.New()

	logger, _ := zap.NewDevelopment()
	e.Use(api.LoggerMiddleware(logger))
	e.Use(api.RequestIDMiddleware())
	e.Use(api.AuthenticationMiddleware())
	e.Use(middleware.Recover())

	api.RegisterHandlers(e, api.New(clock.New(), days))
	e.Logger.Fatal(e.Start(":8000"))
}
