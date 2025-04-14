package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reallvaibhav/E-com-Go/backend/order-service/internal/order/handler"
	"github.com/reallvaibhav/E-com-Go/backend/order-service/internal/order/repository"
	"github.com/reallvaibhav/E-com-Go/backend/order-service/internal/order/usecase"
)

func main() {
	r := gin.Default()

	repo := repository.NewInMemoryOrderRepo()
	uc := usecase.NewOrderUsecase(repo)
	handler.NewOrderHandler(r, uc)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "🛒 Order Service Running"})
	})

	r.Run(":8080")
}
