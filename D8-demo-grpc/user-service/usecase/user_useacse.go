package usecase

import (
	"context"
	"fmt"

	"github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/model"

	"github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/repository"
)

type userUsecase struct {
	userRepo repository.IUserRepository
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, user model.User) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
}

func NewUserUsecase(userRepo repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, user model.User) error {
	// berisi logic business (validation, etc)
	// ....
	return u.userRepo.CreateUser(ctx, user)
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]model.User, error) {
	// berisi logic business (validation, etc)
	// ....
	var users []model.User
	users, err := u.userRepo.GetAllUsers(ctx)
	if err != nil {
		fmt.Println(err)
		return users, err
	}

	return users, nil
}
