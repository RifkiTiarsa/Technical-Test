package usecase

import (
	"test-mnc/entity"
	"test-mnc/repository"
	"time"
)

type topupUsecase struct {
	repo repository.TopupRepository
}

// CreateTopup implements TopupUsecase.
func (t *topupUsecase) CreateTopup(topup entity.Topup) (entity.Topup, error) {
	topup.Status = "pending"

	topup.CreatedAt = time.Now()
	topup.UpdatedAt = time.Now()

	return t.repo.CreateTopup(topup)
}

// GetTopupById implements TopupUsecase.
func (t *topupUsecase) GetTopupById(id int) (entity.Topup, error) {
	return t.repo.GetTopupById(id)
}

func (t *topupUsecase) UpdateBalanceAfterPayment(payload entity.ConfirmTopup) error {
	if err := t.repo.TxTopupUpdateAfterPayment(payload); err != nil {
		return err
	}

	return nil
}

type TopupUsecase interface {
	CreateTopup(topup entity.Topup) (entity.Topup, error)
	GetTopupById(id int) (entity.Topup, error)
	UpdateBalanceAfterPayment(payload entity.ConfirmTopup) error
}

func NewTopupUsecase(repo repository.TopupRepository) TopupUsecase {
	return &topupUsecase{repo: repo}
}
