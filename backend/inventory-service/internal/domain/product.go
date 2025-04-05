package domain

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

// Repository interface (abstraction)
type ProductRepository interface {
	Create(product *Product) error
	FindByID(id string) (*Product, error)
	Update(product *Product) error
	Delete(id string) error
	FindAll() ([]Product, error)
}

// Usecase interface (abstraction)
type ProductUsecase interface {
	CreateProduct(product *Product) error
	GetProduct(id string) (*Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id string) error
	ListProducts() ([]Product, error)
}
