package model

import (
	"time"
)

type User struct {
	ID    int64  `gorm:"primary_key"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
	// Age        sql.NullInt16 // pointer artinya nullable (data bisa null)
	ActiveUser *bool
	CreatedAt  time.Time
	// UpdatedAt  sql.NullTime
	// DeletedAt gorm.DeletedAt // soft delete
	// RoleID     int64

	// Role Role
}

type Role struct {
	ID   int64
	Name string
}

type UserRole struct {
	UserID int64 `gorm:"foreign_key"`
	RoleID int64 `gorm:"foreign_key"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
