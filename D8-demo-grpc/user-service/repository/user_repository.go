package repository

import (
	"context"
	"errors"
	"fmt"
	"hacktiv/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
	DeleteUser(ctx context.Context, id int64) error

	CreateUserWithTransaction(ctx context.Context, tx *gorm.DB, user model.User) error
	CreateAddressWithTransaction(ctx context.Context, tx *gorm.DB, user model.User) error
	UpdateSomethingWithTransaction(ctx context.Context, tx *gorm.DB, user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

// MICROSERVICES => GraphQL & gRPC
// A - B - C
// -- API GATEWAY --

func (u *userRepository) CreateUser(ctx context.Context, user model.User) error {
	// userRole := model.UserRole{
	// 	Name:  user.Name,
	// 	Email: user.Email,
	// }
	res := u.db.Create(&user)
	fmt.Println("USER ID: ", user.ID)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}

	// if user.Age.Valid {
	// 	//
	// }

	return nil
}

func (u *userRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	res := u.db.Find(&users)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}

	return users, nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id int64) error {
	res := u.db.Delete(&model.User{}, id)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}

	return nil
}

func (u *userRepository) CreateUserWithTransaction(ctx context.Context, tx *gorm.DB, user model.User) error {
	res := tx.Create(&user)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}
	return nil
}

func (u *userRepository) CreateAddressWithTransaction(ctx context.Context, tx *gorm.DB, user model.User) error {
	user.Name = "CREATE ADRESS 2"
	res := tx.Create(&user)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}
	return nil
}

func (u *userRepository) UpdateSomethingWithTransaction(ctx context.Context, tx *gorm.DB, user model.User) error {
	user.Name = "UPDATE SOMETHING 2"
	res := tx.Save(&user)

	// error palsu
	if res.Error != nil {
		res.Error = errors.New("ADUH")
		fmt.Println(res.Error)
		return res.Error
	}
	return nil
}
