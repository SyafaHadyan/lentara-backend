package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	CartItemID uuid.UUID `json:"cart_item_id" gorm:"type:varchar(36);primaryKey"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:varchar(36);foreignKey"`
	ProductID  uuid.UUID `json:"product_id" gorm:"type:varchar(36);foreignKey"`
	Count      int32     `json:"count" gorm:"type:int"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

func (c Cart) ParseToDTOCreateCart() dto.CreateCart {
	return dto.CreateCart{
		CartItemID: c.CartItemID,
		UserID:     c.UserID,
		ProductID:  c.ProductID,
		Count:      c.Count,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

func (c Cart) ParseToDTOUpdateCart() dto.UpdateCart {
	return dto.UpdateCart{
		CartItemID: c.CartItemID,
		UserID:     c.UserID,
		ProductID:  c.ProductID,
		Count:      c.Count,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

func (c Cart) ParseToDTOGetCartByID() dto.GetCartByID {
	return dto.GetCartByID{
		CartItemID: c.CartItemID,
		UserID:     c.UserID,
		ProductID:  c.ProductID,
		Count:      c.Count,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

func (c Cart) ParseToDTODeleteCartByCartID() dto.DeleteCartByCartID {
	return dto.DeleteCartByCartID{
		CartItemID: c.CartItemID,
	}
}

func (c Cart) ParseToDTODeleteCartByUserID() dto.DeleteCartByUserID {
	return dto.DeleteCartByUserID{
		UserID: c.UserID,
	}
}
