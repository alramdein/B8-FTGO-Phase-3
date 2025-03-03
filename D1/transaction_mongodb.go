package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	ctx := context.Background()
	clientOpts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Disconnect(ctx) }()

	collBatch8 := client.Database(os.Getenv("MONGODB_DATABASE")).Collection("batch8")
	collTx := client.Database(os.Getenv("MONGODB_DATABASE")).Collection("transactions")

	callback := func(sesctx context.Context) (interface{}, error) {
		if _, err := collBatch8.InsertOne(sesctx, bson.D{{"hoho", 1}}); err != nil {
			return nil, err
		}
		if _, err := collTx.InsertOne(sesctx, bson.D{{"hehe", 999}}); err != nil {
			return nil, err
		}

		return nil, errors.New("ini error")
	}

	session, err := client.StartSession()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("result: %v\n", result)
}
