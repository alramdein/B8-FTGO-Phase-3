package usecase

import (
	"context"
	"fmt"
	"hacktiv/model"

	"hacktiv/repository"
)

type productUsecase struct {
	productRepo repository.IProductRepository
	userClient  userPB.UserClient
}

type IProductUsecase interface {
	CreateProduct(ctx context.Context, product model.Product) error
	GetAllProducts(ctx context.Context) ([]model.Product, error)
}

func NewProductUsecase(productRepo repository.IProductRepository, dbTransactioner repository.DBTransactioner) IProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
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
	var products []model.Product
	products, err := u.productRepo.GetAllProducts(ctx)
	if err != nil {
		fmt.Println(err)
		return products, err
	}

	return products, nil
}
