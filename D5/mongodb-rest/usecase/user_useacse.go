package usecase

import (
	"context"
	"fmt"
	"hacktiv/model"

	"hacktiv/repository"
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
	// harusnya ada logic validaito disini
	// ...
	err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
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
