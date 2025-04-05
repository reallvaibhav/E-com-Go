package handler

import (
	"net/http"
	"order-service/internal/order/usecase"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	uc *usecase.OrderUsecase
}

func NewOrderHandler(r *gin.Engine, uc *usecase.OrderUsecase) {
	h := &OrderHandler{uc: uc}

	r.POST("/orders", h.CreateOrder)
	r.GET("/orders", h.ListOrders)
	r.GET("/orders/:id", h.GetOrder)
	r.PATCH("/orders/:id/status", h.UpdateOrderStatus)
}

type CreateOrderRequest struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o, _ := h.uc.Create(req.ProductID, req.Quantity, req.Price)
	c.JSON(http.StatusCreated, o)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	c.JSON(http.StatusOK, h.uc.List())
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	o, found := h.uc.GetByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, o)
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, found := h.uc.UpdateStatus(id, req.Status)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}
