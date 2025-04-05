package main

import (
	"api-gateway/internal/middleware"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.AuthMiddleware())

	// Proxy to Inventory Service
	inventoryURL, _ := url.Parse("http://localhost:8081")
	inventoryProxy := httputil.NewSingleHostReverseProxy(inventoryURL)
	r.Any("/inventory/*proxyPath", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("proxyPath")
		inventoryProxy.ServeHTTP(c.Writer, c.Request)
	})

	// Proxy to Order Service
	orderURL, _ := url.Parse("http://localhost:8080")
	orderProxy := httputil.NewSingleHostReverseProxy(orderURL)
	r.Any("/orders/*proxyPath", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("proxyPath")
		orderProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API Gateway Running"})
	})

	r.Run(":8000") // Gateway runs on port 8000
}
