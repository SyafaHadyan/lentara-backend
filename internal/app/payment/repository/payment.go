package repository

import (
	"lentara-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type PaymentMySQLItf interface {
	StorePayment(payment *entity.Payment) error
	UpdatePayment(payment *entity.Payment) error
}

type PaymentMySQL struct {
	db *gorm.DB
}

func NewPaymentMySQL(db *gorm.DB) PaymentMySQLItf {
	return &PaymentMySQL{db}
}

func (r *PaymentMySQL) StorePayment(payment *entity.Payment) error {
	return r.db.Debug().Create(payment).Error
}

func (r *PaymentMySQL) UpdatePayment(payment *entity.Payment) error {
	return r.db.Debug().Updates(payment).Error
}
