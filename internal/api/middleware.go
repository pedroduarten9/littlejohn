package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"

	"go.uber.org/zap"
)

const userKey = "user"

func RequestIDMiddleware() echo.MiddlewareFunc {
	return middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return random.String(32)
		},
	})
}

func LoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("Request ID", v.RequestID),
				zap.Time("Start time", v.StartTime),
				zap.Duration("Latency", v.Latency),
				zap.String("URI", v.URI),
				zap.Int("Status", v.Status),
			)

			return nil
		},
	})
}

func AuthenticationMiddleware() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username != "" && password == "" {
			c.Set(userKey, username)
			return true, nil
		}

		return false, AuthError{msg: "Unauthorized"}
	})
}

func HttpErrorHandler(err error, ctx echo.Context) {
	switch e := err.(type) {
	case BadRequestError:
		ctx.JSON(http.StatusBadRequest, Error{Message: err.Error()})
	case NotFoundError:
		ctx.JSON(http.StatusNotFound, Error{Message: err.Error()})
	case AuthError:
		ctx.Response().Header().Set(echo.HeaderWWWAuthenticate, "basic realm=Restricted")
		ctx.JSON(http.StatusUnauthorized, Error{Message: err.Error()})
	default:
		fmt.Print(e)
		ctx.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
	}
}
