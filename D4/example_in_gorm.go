package main

import (
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// hanya di exekusi 1x
func main2() {
	db, err := gorm.Open(postgres.Open(ComposeConnStr()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// config conneciton pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	e := echo.New()

	// built-in middleware
	// e.Use(middleware.WithLogger)
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Gzip())

	// e.HTTPErrorHandler = utils.HTTPErrorHandler

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func ComposeConnStr() string {
	return fmt.Sprintf(`host=%s user=%s  password=%s  dbname=%s  port=%s  sslmode=%s  TimeZone=%s`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)
}
