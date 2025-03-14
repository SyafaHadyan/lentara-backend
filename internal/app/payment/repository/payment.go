package repository

import (
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMySQLItf interface {
	StorePayment(payment *entity.Payment) error
	UpdatePayment(payment *entity.Payment) error
	GetPaymentInfo(payment *entity.Payment, userID uuid.UUID, orderID uuid.UUID) error
	GetPaymentUserInfo(payment *entity.Payment, orderID uuid.UUID) error
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

func (r *PaymentMySQL) GetPaymentInfo(payment *entity.Payment, userID uuid.UUID, orderID uuid.UUID) error {
	return r.db.Debug().Find(payment, userID, orderID).Error
}

func (r *PaymentMySQL) GetPaymentUserInfo(payment *entity.Payment, orderID uuid.UUID) error {
	return r.db.Debug().First(payment, orderID).Error
}
