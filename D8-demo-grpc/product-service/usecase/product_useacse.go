package usecase

import (
	"context"
	"fmt"
	"hacktiv/model"

	"hacktiv/repository"

	userPB "hacktiv/pb/user"
)

type productUsecase struct {
	productRepo repository.IProductRepository
	userClient  userPB.UserServiceClient
}

type IProductUsecase interface {
	CreateProduct(ctx context.Context, product model.Product) error
	GetAllProducts(ctx context.Context) ([]model.Product, error)
}

func NewProductUsecase(productRepo repository.IProductRepository, userClient userPB.UserServiceClient) IProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
		userClient:  userClient,
	}
}

func (u *productUsecase) CreateProduct(ctx context.Context, product model.Product) error {
	// berisi logic business (validation, etc)
	// ....
	return u.productRepo.CreateProduct(ctx, product)
}

func (u *productUsecase) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	// berisi logic business (validation, etc)
	// ....

	// MISALNYA USER CLIENT BUAT VALIDATE TOKEN
	users, err := u.userClient.ListUsers(ctx, &userPB.ListUsersRequest{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Users dari grpc: ", users)

	var products []model.Product
	products, err = u.productRepo.GetAllProducts(ctx)
	if err != nil {
		fmt.Println(err)
		return products, err
	}

	return products, nil
}
