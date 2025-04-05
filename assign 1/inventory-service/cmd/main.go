package main

import (
	"database/sql"
	"inventory-service/internal/handler"
	"inventory-service/internal/repository/postgres"
	"inventory-service/internal/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize DB
	db := initDB()
	defer db.Close()

	// Clean Architecture wiring+
	repo := postgres.NewProductRepository(db)
	usecase := usecase.NewProductUsecase(repo)
	handler := handler.NewProductHandler(usecase)

	// Setup router
	r := gin.Default()

	r.POST("/products", handler.CreateProduct)
	r.GET("/products/:id", handler.GetProduct)
	r.PATCH("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)
	r.GET("/products", handler.ListProducts)

	r.Run(":8080")
}

func initDB() *sql.DB {
	connStr := "user=postgres password=1234 dbname=ecommerce sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (...);`)
	if err != nil {
		panic(err)
	}

	return db
}
