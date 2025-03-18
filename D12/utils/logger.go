package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

func LogEntry(c echo.Context) *logrus.Entry {
	log := logrus.NewEntry(logrus.StandardLogger())
	log.Logger.SetLevel(logrus.DebugLevel)
	if c == nil {
		return log
	}

	// hal penting pada logrus
	return log.WithFields(logrus.Fields{
		"request_id": c.Request().Header.Get("X-Request-ID"),
		"ip":         c.RealIP(),
		"user_agent": c.Request().UserAgent(),
		"path":       c.Path(),
		// ...  populate informasi apapun yang kita inginkan
	})
}

func LogEntry2(c echo.Context) *logrus.Entry {
	log := logrus.New() // untuk logrus
	log.SetFormatter(&ecslogrus.Formatter{})
	if c == nil {
		return log.WithFields(logrus.Fields{})
	}

	return log.WithFields(logrus.Fields{
		"request_id": c.Request().Header.Get("X-Request-ID"),
		"ip":         c.RealIP(),
		"user_agent": c.Request().UserAgent(),
		"path":       c.Path(),
		// ...  populate informasi apapun yang kita inginkan
	})
}
