package repository

import "github.com/jinzhu/gorm"

type OrderPostgres struct {
	p *gorm.DB
}

func NewOrderPostgres(pd *gorm.DB) *OrderPostgres {
	return &OrderPostgres{p: pd}
}
