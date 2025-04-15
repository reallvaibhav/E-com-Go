package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/products", getAllProducts)
	r.GET("/products/:id", getProductByID)
	r.POST("/products", createProduct)
	r.PATCH("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)
	return r
}

func TestGetAllProducts(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateProduct(t *testing.T) {
	router := setupRouter()
	product := Product{ID: "5", Name: "TV", Price: 999}
	body, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
