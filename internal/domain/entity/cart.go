package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	CartItemID   uuid.UUID `json:"cart_item_id" gorm:"type:varchar(36);primaryKey"`
	UserID       uuid.UUID `json:"user_id" gorm:"type:varchar(36);foreignKey"`
	ProductID    uuid.UUID `json:"product_id" gorm:"type:varchar(36);foreignKey"`
	ProductName  string    `json:"product_name" gorm:"type:text"`
	SellerID     uuid.UUID `json:"seller_id" gorm:"type:varchar(36);foreignKey"`
	Count        uint8     `json:"count" gorm:"type:smallint unsigned"`
	Price        uint64    `json:"price" gorm:"type:bigint unsigned"`
	RentDuration uint8     `json:"rent_duration" gorm:"type:smallint unsigned"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

func (c Cart) ParseToDTOCreateCart() dto.CreateCart {
	return dto.CreateCart{
		CartItemID:   c.CartItemID,
		UserID:       c.UserID,
		ProductID:    c.ProductID,
		ProductName:  c.ProductName,
		SellerID:     c.SellerID,
		Count:        c.Count,
		Price:        c.Price,
		RentDuration: c.RentDuration,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

func (c Cart) ParseToDTOUpdateCart() dto.UpdateCart {
	return dto.UpdateCart{
		CartItemID:   c.CartItemID,
		UserID:       c.UserID,
		ProductID:    c.ProductID,
		ProductName:  c.ProductName,
		SellerID:     c.SellerID,
		Count:        c.Count,
		Price:        c.Price,
		RentDuration: c.RentDuration,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

func (c Cart) ParseToDTOGetCartByCartID() dto.GetCartByCartID {
	return dto.GetCartByCartID{
		CartItemID:   c.CartItemID,
		UserID:       c.UserID,
		ProductID:    c.ProductID,
		ProductName:  c.ProductName,
		SellerID:     c.SellerID,
		Count:        c.Count,
		Price:        c.Price,
		RentDuration: c.RentDuration,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

func (c Cart) ParseToDTOGetCartsByUserID() dto.GetCartsByUserID {
	return dto.GetCartsByUserID{
		CartItemID:   c.CartItemID,
		UserID:       c.UserID,
		ProductID:    c.ProductID,
		ProductName:  c.ProductName,
		SellerID:     c.SellerID,
		Count:        c.Count,
		Price:        c.Price,
		RentDuration: c.RentDuration,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

func (c Cart) ParseToDTOGetCartsByUserIDAndSellerID() dto.GetCartsByUserIDAndSellerID {
	return dto.GetCartsByUserIDAndSellerID{
		CartItemID:   c.CartItemID,
		UserID:       c.UserID,
		ProductID:    c.ProductID,
		ProductName:  c.ProductName,
		SellerID:     c.SellerID,
		Count:        c.Count,
		Price:        c.Price,
		RentDuration: c.RentDuration,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}

func (c Cart) ParseToDTOGetCartSummary() dto.GetCartSummary {
	return dto.GetCartSummary{
		ProductName: c.ProductName,
		Count:       c.Count,
		Price:       c.Price,
	}
}

// func (c Cart) ParseToDTOGetCartSummary() dto.GetCartSummary {
// 	return dto.GetCartSummary{
// 		UserID:             c.UserID,
// 		ProductCount:       c.ProductCount,
// 		DeliveryCost:       c.DeliveryCost,
// 		ServiceCost:        c.SeriveCost,
// 		DepositeAmout:      c.DepositeAmount,
// 		DepositePercentage: c.DepositePercentage,
// 		Voucher:            c.Voucher,
// 		TotalPrice:         c.TotalPrice,
// 	}
// }

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
