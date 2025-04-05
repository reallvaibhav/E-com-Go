package postgres

import (
	"database/sql"
	"inventory-service/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	_, err := r.db.Exec(
		"INSERT INTO products (id, name, category, price, stock) VALUES ($1, $2, $3, $4, $5)",
		product.ID, product.Name, product.Category, product.Price, product.Stock,
	)
	return err
}

// Implement other CRUD operations similarly...
