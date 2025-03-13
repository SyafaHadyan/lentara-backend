package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID         uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:varchar(36);foreignKey"`
	TotalPrice uint64    `json:"total_price" gorm:"type:bigint unsigned"`
	Status     string    `json:"status" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

func (p Payment) ParseToDTOStorePayment() dto.StorePayment {
	return dto.StorePayment{
		ID:         p.ID,
		UserID:     p.UserID,
		TotalPrice: p.TotalPrice,
		Status:     p.Status,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func (p Payment) ParseToDTOUpdatePayment() dto.UpdatePaymentStatus {
	return dto.UpdatePaymentStatus{
		ID:         p.ID,
		UserID:     p.UserID,
		TotalPrice: p.TotalPrice,
		Status:     p.Status,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}
