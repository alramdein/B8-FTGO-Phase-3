package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type User struct {
	ID      int
	Name    string
	Age     int
	Address string
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	user := User{
		ID:      8888,
		Name:    "Benny",
		Age:     20,
		Address: "Jakarta",
	}

	DefaultExp := 300 * time.Second

	userByte, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	cacheKey := GenerateUserCacheKey(user.ID)

	err = rdb.Set(ctx, cacheKey, userByte, DefaultExp).Err()
	if err != nil {
		panic(err)
	}

	res := rdb.Get(ctx, cacheKey)
	if res.Err() != nil {
		panic(res.Err())
	}

	fmt.Println(res.Val())

}

func GenerateUserCacheKey(id int) string {
	return fmt.Sprintf("user:%v", id)
}
