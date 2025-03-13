package middleware

import (
	"github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/utils"

	"github.com/labstack/echo/v4"
)

func WithLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		utils.LogEntry(c).Info("request started")
		return next(c)
	}
}
