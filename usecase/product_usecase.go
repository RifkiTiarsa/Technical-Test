package usecase

import (
	"test-mnc/entity"
	"test-mnc/repository"
	"time"
)

type productUsecase struct {
	repo repository.ProductRepository
}

// CreateProduct implements ProductUsecase.
func (p *productUsecase) CreateProduct(payload entity.Product) (entity.Product, error) {
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	return p.repo.CreateProduct(payload)
}

// DeleteProduct implements ProductUsecase.
func (p *productUsecase) DeleteProduct(id int) error {
	product, err := p.repo.GetProductById(id)
	if err != nil {
		return err
	}

	return p.repo.DeleteProduct(product.ID)

}

// GetAllProducts implements ProductUsecase.
func (p *productUsecase) GetAllProducts() ([]entity.Product, error) {
	return p.repo.GetAllProducts()
}

// GetProductById implements ProductUsecase.
func (p *productUsecase) GetProductById(id int) (entity.Product, error) {
	return p.repo.GetProductById(id)
}

// UpdateProduct implements ProductUsecase.
func (p *productUsecase) UpdateProduct(payload entity.Product) (entity.Product, error) {
	product, err := p.repo.GetProductById(payload.ID)
	if err != nil {
		return entity.Product{}, err
	}

	product.Name = payload.Name
	product.Nominal = payload.Nominal
	product.Price = payload.Price
	product.UpdatedAt = time.Now()

	return p.repo.UpdateProduct(product)
}

type ProductUsecase interface {
	CreateProduct(payload entity.Product) (entity.Product, error)
	GetProductById(id int) (entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	UpdateProduct(payload entity.Product) (entity.Product, error)
	DeleteProduct(id int) error
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}
