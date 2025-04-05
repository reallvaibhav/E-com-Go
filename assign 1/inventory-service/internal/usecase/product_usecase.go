package usecase

import (
	"inventory-service/internal/domain"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{repo: repo}
}

func (uc *productUsecase) CreateProduct(product *domain.Product) error {
	return uc.repo.Create(product)
}

func (uc *productUsecase) GetProduct(id string) (*domain.Product, error) {
	return uc.repo.FindByID(id)
}

func (uc *productUsecase) UpdateProduct(product *domain.Product) error {
	return uc.repo.Update(product)
}

func (uc *productUsecase) DeleteProduct(id string) error {
	return uc.repo.Delete(id)
}

func (uc *productUsecase) ListProducts() ([]domain.Product, error) {
	return uc.repo.FindAll()
}
