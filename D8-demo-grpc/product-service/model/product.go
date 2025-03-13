package model

import (
	"time"
)

type Product struct {
	ID        int64  `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
}

type Role struct {
	ID   int64
	Name string
}

type ProductRole struct {
	ProductID int64 `gorm:"foreign_key"`
	RoleID    int64 `gorm:"foreign_key"`
}

type ProductResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
