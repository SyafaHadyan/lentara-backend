package dto

import (
	"time"

	"github.com/google/uuid"
)

type MidtransResponseNewPayment struct {
	Token    string `json:"token"`
	Redirect string `json:"redirect"`
}

type MidtransUpdatePaymentStatus struct {
	TransasctionTime  string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	StatusMessage     string `json:"status_message"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	GrossAmount       string `json:"gross_amount"`
	FraudStatus       string `json:"fraud_status"`
	Currenncy         string `json:"currency"`
}

type MidtransResponseUpdatePaymentStatus struct {
	TransasctionTime  string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	GrossAmount       string `json:"gross_amount"`
	Currenncy         string `json:"currency"`
}

type StorePayment struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	TotalPrice uint64    `json:"total_price"`
	Status     string    `json:"transaction_status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UpdatePaymentStatus struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	TotalPrice uint64    `json:"total_price"`
	Status     string    `json:"transaction_status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetPaymentStatus struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	TotalPrice uint64    `json:"total_price"`
	Status     string    `json:"transaction_status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
