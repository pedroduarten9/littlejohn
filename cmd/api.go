package main

import (
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"littlejohn/internal/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	logger, _ := zap.NewDevelopment()
	e.Use(api.LoggerMiddleware(logger))
	e.Use(api.RequestIDMiddleware())
	e.Use(api.AuthenticationMiddleware())
	e.Use(middleware.Recover())

	api.RegisterHandlers(e, api.NewAPI())
	e.Logger.Fatal(e.Start(":8000"))
}
