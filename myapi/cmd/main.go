package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"myapi/db"
	"myapi/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	// Connect to database
	db.Connect()

	// Initialize Gin router
	r := gin.Default()

	// Register all routes
	routes.SetupRouter(r)

	// Start the server
	log.Println("🚀 Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
