package main

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
	{ID: "1", Name: "Phone", Price: 500},
	{ID: "2", Name: "Laptop", Price: 1200},
}

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, p := range products {
			if p.ID == id {
				c.JSON(http.StatusOK, p)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	r.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.BindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products = append(products, newProduct)
		c.JSON(http.StatusCreated, newProduct)
	})

	r.PATCH("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updated Product
		if err := c.BindJSON(&updated); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, p := range products {
			if p.ID == id {
				products[i].Name = updated.Name
				products[i].Price = updated.Price
				c.JSON(http.StatusOK, products[i])
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, p := range products {
			if p.ID == id {
				products = append(products[:i], products[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	r.Run(":8081")
}
