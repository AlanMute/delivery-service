package repository

import (
	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/jinzhu/gorm"
)

type Order interface {
	Add(ord core.Order) error
}

type Repository struct {
	Order
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
	}
}
