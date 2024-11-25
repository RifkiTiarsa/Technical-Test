package repository

import (
	"database/sql"
	"test-mnc/config"
	"test-mnc/entity"
)

type merchantRepository struct {
	db *sql.DB
}

// CreateMerchant implements MerchantRepository.
func (m *merchantRepository) CreateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	if err := m.db.QueryRow(config.InsertMerchant, payload.Name, payload.Balance, payload.CreatedAt, payload.UpdatedAt).Scan(&payload.ID); err != nil {
		return entity.Merchant{}, err
	}

	return payload, nil
}

type MerchantRepository interface {
	CreateMerchant(payload entity.Merchant) (entity.Merchant, error)
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db: db}
}
