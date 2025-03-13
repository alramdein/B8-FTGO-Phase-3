package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	httpHandler "github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/handler/http"
	"github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/middleware"
	"github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/model"
	"github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/repository"
	"github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/usecase"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	_ "github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/docs"
	grpcHandler "github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/handler/grpc"
	userPB "github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/pb/user"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(composeConnStr()), &gorm.Config{})
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

	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	quitCh := make(chan bool, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		for {
			select {
			case <-sigCh:
				quitCh <- true
			case err := <-errCh:
				log.Fatal(err)
				quitCh <- true
			}
		}
	}()

	go func() {
		InitHTTPServer(db, errCh)
	}()

	go func() {
		InitGrpcServer(db, errCh)
	}()

	<-quitCh
	fmt.Println("exiting program...")
}

func InitHTTPServer(db *gorm.DB, errCh chan error) {
	e := echo.New()

	// built-in middleware
	e.Use(middleware.WithLogger)
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Gzip())

	// e.HTTPErrorHandler = utils.HTTPErrorHandler

	e.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusPreconditionFailed, httpHandler.ErrUserIsNotSuperadmin)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	userRepo := repository.NewUserRepository(db)
	userUseacase := usecase.NewUserUsecase(userRepo)
	userHandler := httpHandler.NewUserHandler(userUseacase)

	userHandler.RegisterUserRoutes(e)

	// e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

	errCh <- e.Start(":" + os.Getenv("PORT"))
}

func InitGrpcServer(db *gorm.DB, errChan chan error) {
	port := ":1234" // pake env
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(withAuth),
	}

	userRepo := repository.NewUserRepository(db)
	userUseacase := usecase.NewUserUsecase(userRepo)

	userSvc := grpcHandler.NewUserGrpcServer(userUseacase)
	gprcServer := grpc.NewServer(opts...)

	userPB.RegisterUserServiceServer(gprcServer, userSvc)

	fmt.Println("Grpc server started at port ", port)
	if err := gprcServer.Serve(lis); err != nil {
		errChan <- err
	}
}

func withAuth(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	if len(md["authorization"]) == 0 {
		return nil, fmt.Errorf("missing token")
	}

	// validate token blabla
	token, ok := md["authorization"]
	if !ok {
		return nil, fmt.Errorf("missing token")
	}

	if token[0] != "kajnfjanfandfkljadnALIF" {
		return nil, fmt.Errorf("invalid token")
	}

	return handler(ctx, req)
}

func composeConnStr() string {
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
