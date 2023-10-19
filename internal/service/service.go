package service

import "github.com/KrizzMU/delivery-service/internal/repository"

type Order interface {
}

type Service struct {
	Order
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(&r.Order),
	}
}
