package entity

import "time"

type Topup struct {
	ID            int       `json:"id"`
	CustomerID    int       `json:"customer_id"`
	MerchantID    int       `json:"merchant_id"`
	ProductID     int       `json:"product_id"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type ConfirmTopup struct {
	TopupID       int     `json:"topup_id"`
	Amount        float64 `json:"amount"`
	Price         float64 `json:"price"`
	PaymentMethod string  `json:"payment_method"`
	PaymentStatus string  `json:"payment_status"`
}
