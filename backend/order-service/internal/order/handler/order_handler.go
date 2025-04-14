package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reallvaibhav/E-com-Go/backend/order-service/internal/order/usecase"
)

type OrderHandler struct {
	uc usecase.OrderUsecase
}

func NewOrderHandler(r *gin.Engine, uc usecase.OrderUsecase) {
	h := &OrderHandler{uc}

	r.GET("/orders", h.GetOrders)
	r.GET("/orders/:id", h.GetOrder)
	r.POST("/orders", h.CreateOrder)
	r.DELETE("/orders/:id", h.DeleteOrder)
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, h.uc.GetAll())
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, ok := h.uc.GetByID(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req usecase.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.uc.Create(req)
	c.JSON(http.StatusCreated, req)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if ok := h.uc.Delete(id); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
