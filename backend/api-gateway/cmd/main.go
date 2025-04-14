package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/reallvaibhav/E-com-Go/backend/api-gateway/internal/middleware"
)

const (
	apiGatewayPort   = ":8000"
	inventoryService = "http://localhost:8081"
	orderService     = "http://localhost:8080"
)

func main() {
	r := gin.Default()

	// Middleware
	r.Use(middleware.Logger())
	r.Use(middleware.AuthMiddleware())

	// Proxy: Inventory Service
	inventoryProxy := createReverseProxy(inventoryService)
	r.Any("/inventory/*proxyPath", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("proxyPath")
		inventoryProxy.ServeHTTP(c.Writer, c.Request)
	})

	// Proxy: Order Service
	orderProxy := createReverseProxy(orderService)
	r.Any("/orders/*proxyPath", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("proxyPath")
		orderProxy.ServeHTTP(c.Writer, c.Request)
	})

	// Health Check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "✅ API Gateway Running"})
	})

	log.Printf("🚀 API Gateway started on %s", apiGatewayPort)
	if err := r.Run(apiGatewayPort); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

// createReverseProxy creates a reverse proxy to the given target URL
func createReverseProxy(target string) *httputil.ReverseProxy {
	parsedURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid service URL %s: %v", target, err)
	}
	return httputil.NewSingleHostReverseProxy(parsedURL)
}
