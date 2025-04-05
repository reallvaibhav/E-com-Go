package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"inventory-service/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}

	_, err := r.db.Exec(
		`INSERT INTO products (id, name, category, price, stock) 
		 VALUES ($1, $2, $3, $4, $5)`,
		product.ID, product.Name, product.Category, product.Price, product.Stock,
	)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

func (r *ProductRepository) GetByID(id string) (*domain.Product, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	row := r.db.QueryRow(
		`SELECT id, name, category, price, stock FROM products WHERE id = $1`,
		id,
	)

	var product domain.Product
	err := row.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Stock)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrProductNotFound // Custom error (define in domain)
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	return &product, nil
}

func (r *ProductRepository) Update(product *domain.Product) error {
	if product == nil {
		return errors.New("product cannot be nil")
	}

	result, err := r.db.Exec(
		`UPDATE products 
		 SET name = $1, category = $2, price = $3, stock = $4 
		 WHERE id = $5`,
		product.Name, product.Category, product.Price, product.Stock, product.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrProductNotFound
	}
	return nil
}

func (r *ProductRepository) Delete(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}

	result, err := r.db.Exec(
		`DELETE FROM products WHERE id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.ErrProductNotFound
	}
	return nil
}

func (r *ProductRepository) ListAll() ([]domain.Product, error) {
	rows, err := r.db.Query(
		`SELECT id, name, category, price, stock FROM products`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Category,
			&product.Price,
			&product.Stock,
		); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return products, nil
}
