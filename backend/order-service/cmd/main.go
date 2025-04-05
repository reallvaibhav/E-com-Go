package main

import (
	"order-service/internal/order/handler"
	"order-service/internal/order/repository"
	"order-service/internal/order/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	repo := repository.NewInMemoryOrderRepo()
	uc := usecase.NewOrderUsecase(repo)
	handler.NewOrderHandler(r, uc)

	r.Run(":8080") // Run on port 8080
}
