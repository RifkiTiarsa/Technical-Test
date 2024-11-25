package entity

import "time"

type Product struct {
	ID         int       `json:"id"`
	MerchantID int       `json:"merchant_id"`
	Name       string    `json:"name"`
	Nominal    float64   `json:"nominal"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
