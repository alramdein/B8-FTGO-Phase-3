package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitMongoDB(ctx context.Context) (*mongo.Client, *mongo.Database) {
	clientOpts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		panic(err)
	}

	// ping server
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	// ping collection
	db := client.Database(os.Getenv("MONGODB_DATABASE"))
	_, err = db.Collection("ping").Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	return client, db
}
