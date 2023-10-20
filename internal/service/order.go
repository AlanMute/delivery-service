package service

import (
	"fmt"

	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/KrizzMU/delivery-service/internal/repository"
	"github.com/KrizzMU/delivery-service/pkg/cache"
)

type OrderService struct {
	repo repository.Order
	c    *cache.Cache
}

func NewOrderService(r repository.Order, c *cache.Cache) *OrderService {
	return &OrderService{
		repo: r,
		c:    c,
	}
}

func (s *OrderService) RecoveryCache(ords []core.Order) {
	for _, ord := range ords {
		s.c.Add(ord.OrderUID, ord)
		fmt.Println("Востановлено")
		fmt.Println(s.c.Get("b563feb7b2b84b6test"))
		fmt.Println()
	}
}

func (s *OrderService) Create(ord core.Order) error {

	err := s.repo.Add(ord)

	if err != nil {
		return err
	}

	s.c.Add(ord.OrderUID, ord)
	fmt.Println(s.c.Get("b563feb7b2b84b6test"))
	return nil
}
