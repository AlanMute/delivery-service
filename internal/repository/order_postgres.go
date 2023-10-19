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
	delivery := ord.Delivery
	payment := ord.Payment
	items := ord.Items

	tx := r.db.Begin()

	if err := tx.Create(&delivery).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range items {
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
