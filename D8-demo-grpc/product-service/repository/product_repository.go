package repository

import (
	"context"
	"fmt"
	"hacktiv/model"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(ctx context.Context, product model.Product) error
	GetAllProducts(ctx context.Context) ([]model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{
		db: db,
	}
}

func (u *productRepository) CreateProduct(ctx context.Context, product model.Product) error {
	// productRole := model.ProductRole{
	// 	Name:  product.Name,
	// 	Email: product.Email,
	// }
	res := u.db.Create(&product)
	fmt.Println("USER ID: ", product.ID)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}

	// if product.Age.Valid {
	// 	//
	// }

	return nil
}

func (u *productRepository) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	res := u.db.Find(&products)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}

	return products, nil
}
