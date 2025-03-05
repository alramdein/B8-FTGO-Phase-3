package usecase

import (
	"context"
	"fmt"
	"hacktiv/model"

	"hacktiv/repository"
)

type userUsecase struct {
	userRepo        repository.IUserRepository
	dbTransactioner repository.DBTransactioner
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, user model.User) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
	DeleteUser(ctx context.Context, id int64) error
	TransactionExample(ctx context.Context, user model.User) error
}

func NewUserUsecase(userRepo repository.IUserRepository, dbTransactioner repository.DBTransactioner) IUserUsecase {
	return &userUsecase{
		userRepo:        userRepo,
		dbTransactioner: dbTransactioner,
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

func (u *userUsecase) DeleteUser(ctx context.Context, id int64) error {
	// berisi logic business (validation, etc)
	// ....
	return u.userRepo.DeleteUser(ctx, id)
}

func (u *userUsecase) TransactionExample(ctx context.Context, user model.User) error {

	tx := u.dbTransactioner.BeginTransaction()

	err := u.userRepo.CreateUserWithTransaction(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = u.userRepo.CreateAddressWithTransaction(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = u.userRepo.UpdateSomethingWithTransaction(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
