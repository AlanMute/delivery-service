package repository

import (
	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/jinzhu/gorm"
)

type OrderPostgres struct {
	db *gorm.DB
}

func NewOrderPostgres(pd *gorm.DB) *OrderPostgres {
	return &OrderPostgres{db: pd}
}

func (r *OrderPostgres) Add(ord core.Order) error {
	tx := r.db.Begin()

	if err := tx.Create(&ord.Delivery).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&ord.Payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range ord.Items {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Create(&ord).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *OrderPostgres) GetAll() []core.Order {
	var orders []core.Order

	r.db.Preload("Delivery").Preload("Payment").Find(&orders)

	for i := range orders {
		r.db.Where("track_number = ?", orders[i].TrackNumber).Find(&orders[i].Items)
	}

	return orders
}
