package usecase

import (
	"fmt"
	"test-mnc/entity"
	"test-mnc/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type customerUsecase struct {
	repo repository.CustomerRepository
}

// Create implements CustomerUsecase.
func (c *customerUsecase) RegisterCustomer(payload entity.Customer) (entity.Customer, error) {
	customer, _ := c.repo.GetCustomerByEmail(payload.Email)
	if customer.Email == payload.Email {
		return entity.Customer{}, fmt.Errorf("Email already registered")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.Customer{}, fmt.Errorf("Error hashing password")
	}

	payload.Password = string(passwordHash)
	payload.Balance = 0
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	return c.repo.CreateCustomer(payload)
}

// FindCustomerById implements CustomerUsecase.
func (c *customerUsecase) FindCustomerById(id string) (entity.Customer, error) {
	return c.repo.GetCustomerById(id)
}

// FindCustomerByEmailPassword implements CustomerUsecase.
func (c *customerUsecase) FindCustomerByEmailPassword(email string, password string) (entity.Customer, error) {
	customerExist, err := c.repo.GetCustomerByEmail(email)
	if err != nil {
		return entity.Customer{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(customerExist.Password), []byte(password))
	if err != nil {
		return entity.Customer{}, fmt.Errorf("Invalid email or password")
	}

	return customerExist, nil
}

type CustomerUsecase interface {
	RegisterCustomer(payload entity.Customer) (entity.Customer, error)
	FindCustomerById(id string) (entity.Customer, error)
	FindCustomerByEmailPassword(email, password string) (entity.Customer, error)
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo: repo}
}
