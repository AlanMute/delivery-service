package handlers

import (
	"github.com/KrizzMU/delivery-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	r.Handle("GET", "/order/:id", h.GetById)

	r.Static("/info", "./static")

	return r
}
