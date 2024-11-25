package repository

import (
	"database/sql"
	"test-mnc/config"
	"test-mnc/entity"
)

type customerRepository struct {
	db *sql.DB
}

// CreateCustomer implements CustomerRepository.
func (c *customerRepository) CreateCustomer(payload entity.Customer) (entity.Customer, error) {
	if err := c.db.QueryRow(config.InsertCustomer, payload.Name, payload.Email, payload.Password, payload.Balance, payload.CreatedAt, payload.UpdatedAt).Scan(&payload.ID); err != nil {
		return entity.Customer{}, err
	}

	return payload, nil
}

// GetCustomerById implements CustomerRepository.
func (c *customerRepository) GetCustomerById(id string) (entity.Customer, error) {
	var customer entity.Customer

	if err := c.db.QueryRow(config.GetCustomerById, id).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Password, &customer.Balance, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
		return entity.Customer{}, err
	}

	return customer, nil
}

// GetCustomerByUsername implements CustomerRepository.
func (c *customerRepository) GetCustomerByEmail(email string) (entity.Customer, error) {
	var customer entity.Customer

	if err := c.db.QueryRow(config.GetCustomerByUsername, email).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Password, &customer.Balance, &customer.CreatedAt, &customer.UpdatedAt); err != nil {
		return entity.Customer{}, err
	}

	return customer, nil
}

type CustomerRepository interface {
	CreateCustomer(payload entity.Customer) (entity.Customer, error)
	GetCustomerById(id string) (entity.Customer, error)
	GetCustomerByEmail(email string) (entity.Customer, error)
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
