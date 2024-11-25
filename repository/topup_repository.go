package repository

import (
	"database/sql"
	"fmt"
	"test-mnc/config"
	"test-mnc/entity"
)

type topupRepository struct {
	db *sql.DB
}

// GetTopupById implements TopupRepository.
func (t *topupRepository) GetTopupById(id int) (entity.Topup, error) {
	var topup entity.Topup

	if err := t.db.QueryRow(config.GetTopupById, id).Scan(&topup.ID, &topup.CustomerID, &topup.MerchantID, &topup.ProductID, &topup.PaymentMethod, &topup.Status, &topup.CreatedAt, &topup.UpdatedAt); err != nil {
		return entity.Topup{}, err
	}

	return topup, nil
}

// UpdateBalanceCustomer implements TopupRepository.
func (t *topupRepository) UpdateBalanceCustomer(tx *sql.Tx, balance float64, idCustomer int) error {
	if _, err := tx.Exec(config.UpdateBalanceCustomer, balance, idCustomer); err != nil {
		return err
	}

	return nil
}

// UpdateBalanceMerchant implements TopupRepository.
func (t *topupRepository) UpdateBalanceMerchant(tx *sql.Tx, balance float64, idMerchant int) error {
	if _, err := tx.Exec(config.UpdateBalanceMerchant, balance, idMerchant); err != nil {
		return err
	}

	return nil
}

// UpdateStatus implements TopupRepository.
func (t *topupRepository) UpdateStatus(tx *sql.Tx, status string, idTopup int) error {
	if _, err := tx.Exec(config.UpdateStatusTopup, status, idTopup); err != nil {
		return err
	}

	return nil
}

// CreateTopup implements TopupRepository.
func (t *topupRepository) CreateTopup(topup entity.Topup) (entity.Topup, error) {
	if err := t.db.QueryRow(config.InsertTopup, topup.CustomerID, topup.MerchantID, topup.ProductID, topup.PaymentMethod, topup.Status, topup.CreatedAt, topup.UpdatedAt).Scan(&topup.ID); err != nil {
		return entity.Topup{}, err
	}

	return topup, nil
}

func (t *topupRepository) TxTopupUpdateAfterPayment(topup entity.ConfirmTopup) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dataTopup, err := t.GetTopupById(topup.TopupID)
	if err != nil {
		return err
	}

	status := "pending"

	if topup.PaymentStatus == "Done" {
		status = "completed"
	}

	if err := t.UpdateStatus(tx, status, dataTopup.ID); err != nil {
		return err
	}

	if err := t.UpdateBalanceCustomer(tx, topup.Amount, dataTopup.CustomerID); err != nil {
		return err
	}

	if err := t.UpdateBalanceMerchant(tx, topup.Amount, dataTopup.MerchantID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	fmt.Println("Topup transaction commited")

	return nil
}

type TopupRepository interface {
	CreateTopup(topup entity.Topup) (entity.Topup, error)
	GetTopupById(id int) (entity.Topup, error)
	UpdateStatus(tx *sql.Tx, status string, idTopup int) error
	UpdateBalanceCustomer(tx *sql.Tx, balance float64, idCustomer int) error
	UpdateBalanceMerchant(tx *sql.Tx, balance float64, idMerchant int) error
	TxTopupUpdateAfterPayment(topup entity.ConfirmTopup) error
}

func NewTopupRepository(db *sql.DB) TopupRepository {
	return &topupRepository{db: db}
}
