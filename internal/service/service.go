package service

import (
	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/KrizzMU/delivery-service/internal/repository"
)

type Order interface {
	Create(ord core.Order) error
}

type Service struct {
	Order
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(&r.Order),
	}
}
