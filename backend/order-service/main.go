package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

var orders = make(map[string]Order)
var mu sync.Mutex

func createOrder(c *gin.Context) {
	var newOrder Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mu.Lock()
	orders[newOrder.ID] = newOrder
	mu.Unlock()
	c.JSON(http.StatusCreated, newOrder)
}

func getOrder(c *gin.Context) {
	id := c.Param("id")
	mu.Lock()
	order, exists := orders[id]
	mu.Unlock()
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func updateOrder(c *gin.Context) {
	id := c.Param("id")
	var updatedOrder Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	order, exists := orders[id]
	mu.Unlock()
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.Amount = updatedOrder.Amount
	order.Status = updatedOrder.Status

	mu.Lock()
	orders[id] = order
	mu.Unlock()

	c.JSON(http.StatusOK, order)
}

func deleteOrder(c *gin.Context) {
	id := c.Param("id")
	mu.Lock()
	_, exists := orders[id]
	if exists {
		delete(orders, id)
	}
	mu.Unlock()
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func main() {
	r := gin.Default()

	r.POST("/orders", createOrder)
	r.GET("/orders/:id", getOrder)
	r.PUT("/orders/:id", updateOrder)
	r.DELETE("/orders/:id", deleteOrder)

	r.Run(":8080") // Run on port 8080
}
