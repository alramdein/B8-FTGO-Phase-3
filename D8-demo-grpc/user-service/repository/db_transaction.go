package repository

import "gorm.io/gorm"

type DBTransactioner interface {
	BeginTransaction() *gorm.DB
}

type dbTransactioner struct {
	db *gorm.DB
}

func NewDBTransactioner(db *gorm.DB) DBTransactioner {
	return &dbTransactioner{
		db: db,
	}
}

func (d dbTransactioner) BeginTransaction() *gorm.DB {
	return d.db.Begin()
}
