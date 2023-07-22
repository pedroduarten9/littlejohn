package main

import (
	"littlejohn/internal/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	api.RegisterHandlers(e, api.NewAPI())
	e.Logger.Fatal(e.Start(":8000"))
}
