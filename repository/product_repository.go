package repository

import (
	"database/sql"
	"test-mnc/config"
	"test-mnc/entity"
)

type productRepository struct {
	db *sql.DB
}

// CreateProduct implements ProductRepository.
func (p *productRepository) CreateProduct(payload entity.Product) (entity.Product, error) {
	if err := p.db.QueryRow(config.InsertProduct, payload.MerchantID, payload.Name, payload.Nominal, payload.Price, payload.CreatedAt, payload.UpdatedAt).Scan(&payload.ID); err != nil {
		return entity.Product{}, err
	}

	return payload, nil
}

// DeleteProduct implements ProductRepository.
func (p *productRepository) DeleteProduct(id int) error {
	_, err := p.db.Exec(config.DeleteProduct, id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllProducts implements ProductRepository.
func (p *productRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product

	rows, err := p.db.Query(config.GetAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.MerchantID, &product.Name, &product.Nominal, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetProductById implements ProductRepository.
func (p *productRepository) GetProductById(id int) (entity.Product, error) {
	var product entity.Product

	if err := p.db.QueryRow(config.GetProductById, id).Scan(&product.ID, &product.MerchantID, &product.Name, &product.Nominal, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

// UpdateProduct implements ProductRepository.
func (p *productRepository) UpdateProduct(payload entity.Product) (entity.Product, error) {
	row, err := p.db.Exec(config.UpdateProduct, payload.ID, payload.MerchantID, payload.Name, payload.Nominal, payload.Price, payload.CreatedAt, payload.UpdatedAt)
	if err != nil {
		return entity.Product{}, err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return entity.Product{}, err
	}

	return payload, nil
}

type ProductRepository interface {
	CreateProduct(payload entity.Product) (entity.Product, error)
	GetProductById(id int) (entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	UpdateProduct(payload entity.Product) (entity.Product, error)
	DeleteProduct(id int) error
}

// NewProductRepository returns a new ProductRepository instance.
func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
