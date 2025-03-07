package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	UserID    uuid.UUID `json:"user_id" gorm:"type:varchar(36);foreignKey"`
	ProductID uuid.UUID `json:"product_id" gorm:"type:varchar(36);foreignKey"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

func (c Cart) ParseToDTOCreateCart() dto.CreateCart {
	return dto.CreateCart{
		UserID:    c.UserID,
		ProductID: c.ProductID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
