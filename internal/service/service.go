package service

import (
	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/KrizzMU/delivery-service/internal/repository"
	"github.com/KrizzMU/delivery-service/pkg/cache"
)

type Order interface {
	Create(ord core.Order) error
	RecoveryCache(ords []core.Order)
}

type Service struct {
	Order
}

func NewService(r *repository.Repository, c *cache.Cache) *Service {
	return &Service{
		Order: NewOrderService(r.Order, c),
	}
}
