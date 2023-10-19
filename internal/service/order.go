package service

import (
	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/KrizzMU/delivery-service/internal/repository"
)

type OrderService struct {
	r *repository.Order
}

func NewOrderService(repo *repository.Order) *OrderService {
	return &OrderService{r: repo}
}

func (*OrderService) Create(ord core.Order) error {
	return nil
}
