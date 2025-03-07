package main

import (
	"context"
	"hacktiv/config"
	"hacktiv/handler"
	"hacktiv/middleware"
	"hacktiv/repository"
	"hacktiv/usecase"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "hacktiv/docs"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()
	client, db := config.InitMongoDB(ctx)
	defer func() { _ = client.Disconnect(ctx) }()

	e := echo.New()

	// built-in middleware
	e.Use(middleware.WithLogger)
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Gzip())

	// e.HTTPErrorHandler = utils.HTTPErrorHandler

	e.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusPreconditionFailed, handler.ErrUserIsNotSuperadmin)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	userRepo := repository.NewUserRepository(db)
	userUseacase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUseacase)

	userHandler.RegisterUserRoutes(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
