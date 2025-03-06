package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig(os.Getenv("DB_CONN_STRING"))
	if err != nil {
		panic(err)
	}

	config.MaxConnLifetime = time.Minute * 30
	config.MaxConnIdleTime = time.Minute * 10
	config.HealthCheckPeriod = time.Minute * 2
	config.MaxConnLifetimeJitter = time.Minute * 5

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Release()

	users, err := conn.Query(ctx, "SELECT * FROM users;")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", users)
}
