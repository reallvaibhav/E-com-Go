package handler

import (
	"net/http"

	"inventory-service/internal/domain"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase domain.ProductUsecase
}

func NewProductHandler(usecase domain.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// Implement other handler methods similarly...
