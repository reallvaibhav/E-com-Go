package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []Product{
	{ID: "1", Name: "Laptop", Price: 1200},
	{ID: "2", Name: "Phone", Price: 800},
}

func RegisterRoutes(r *gin.Engine) {
	productRoutes := r.Group("/products")
	{
		productRoutes.GET("/", GetAllProducts)
		productRoutes.GET("/:id", GetProductByID)
	}
}

func GetAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}
