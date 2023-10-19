package repository

import "github.com/jinzhu/gorm"

type Order interface {
}

type Repository struct {
	Order
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
	}
}
