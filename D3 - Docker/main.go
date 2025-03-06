package main

import (
	"fmt"
	"hacktiv/handler"
	"hacktiv/middleware"
	"hacktiv/model"
	"hacktiv/repository"
	"hacktiv/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "hacktiv/docs"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// hanya di exekusi 1x
func main() {
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

	db.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRole{})

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

	dbTransactioner := repository.NewDBTransactioner(db)
	userRepo := repository.NewUserRepository(db)
	userUseacase := usecase.NewUserUsecase(userRepo, dbTransactioner)
	userHandler := handler.NewUserHandler(userUseacase)

	userHandler.RegisterUserRoutes(e)

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
