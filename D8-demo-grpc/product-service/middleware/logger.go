package middleware

import (
	"hacktiv/utils"

	"github.com/labstack/echo/v4"
)

func WithLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		utils.LogEntry(c).Info("request started")
		return next(c)
	}
}
