package service

import (
	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/KrizzMU/delivery-service/internal/repository"
	"github.com/KrizzMU/delivery-service/pkg/cache"
)

type OrderService struct {
	repo repository.Order
	c    cache.Cache
}

func NewOrderService(r repository.Order) *OrderService {
	return &OrderService{
		repo: r,
		c:    *cache.NewCache(),
	}
}

func (s *OrderService) Create(ord core.Order) error {

	err := s.repo.Add(ord)

	if err != nil {
		return err
	}

	s.c.Add(ord.OrderUID, ord)

	return nil
}
