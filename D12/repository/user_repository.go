package repository

import (
	"context"
	"fmt"
	"hacktiv/model"
	"time"

	"go.elastic.co/apm/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
}

type userRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) IUserRepository {
	return &userRepository{
		userCollection: db.Collection("users"),
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user model.User) error {
	fmt.Println(user)
	res, err := u.userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	fmt.Println("inserted id: ", res.InsertedID)
	return nil
}

func (u *userRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	span, _ := apm.StartSpan(ctx, "repository.GetAllUsers", "repo")
	defer span.End()

	time.Sleep(2 * time.Second)

	var users []model.User
	cursor, err := u.userCollection.Find(ctx, bson.D{})
	if err != nil {
		return users, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		return users, err
	}
	return users, nil
}
