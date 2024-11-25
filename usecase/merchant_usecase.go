package usecase

import (
	"test-mnc/entity"
	"test-mnc/repository"
	"time"
)

type merchantUsecase struct {
	repo repository.MerchantRepository
}

// CreateMerchant implements MerchantUsecase.
func (m *merchantUsecase) CreateMerchant(payload entity.Merchant) (entity.Merchant, error) {
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	return m.repo.CreateMerchant(payload)
}

type MerchantUsecase interface {
	CreateMerchant(payload entity.Merchant) (entity.Merchant, error)
}

func NewMerchantUsecase(repo repository.MerchantRepository) MerchantUsecase {
	return &merchantUsecase{repo: repo}
}
