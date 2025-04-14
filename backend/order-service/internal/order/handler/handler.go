package handler

import (
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	uc *usecase.orderUsecase
}

func NewOrderHandler(r *gin.Engine, uc *usecase.orderUsecase) {
	h := &OrderHandler{uc: uc}
	r.GET("/orders", h.GetOrders)
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	orders := h.uc.repo.FetchAll()
	c.JSON(200, orders)
}
