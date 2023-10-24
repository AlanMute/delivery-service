package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetById(c *gin.Context) {
	orderID := c.Param("id")

	order, err := h.services.Order.Get(orderID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	//fmt.Println(order)
	c.JSON(http.StatusOK, order)
}
