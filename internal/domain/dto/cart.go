package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateCart struct {
	CartItemID   uuid.UUID `json:"cart_item_id"`
	UserID       uuid.UUID `json:"user_id"`
	ProductID    uuid.UUID `json:"product_id" validate:"required"`
	SellerID     uuid.UUID `json:"seller_id"`
	Count        uint8     `json:"count" validate:"required"`
	RentDuration uint8     `json:"rent_duration" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateCart struct {
	CartItemID   uuid.UUID `json:"cart_item_id"`
	UserID       uuid.UUID `json:"user_id"`
	ProductID    uuid.UUID `json:"product_id"`
	SellerID     uuid.UUID `json:"seller_id"`
	Count        uint8     `json:"count"`
	RentDuration uint8     `json:"rent_duration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetCartByCartID struct {
	CartItemID   uuid.UUID `json:"cart_item_id"`
	UserID       uuid.UUID `json:"user_id"`
	ProductID    uuid.UUID `json:"product_id"`
	SellerID     uuid.UUID `json:"seller_id"`
	Count        uint8     `json:"count"`
	RentDuration uint8     `json:"rent_duration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetCartsByUserID struct {
	CartItemID   uuid.UUID `json:"cart_item_id"`
	UserID       uuid.UUID `json:"user_id"`
	ProductID    uuid.UUID `json:"product_id"`
	SellerID     uuid.UUID `json:"seller_id"`
	Count        uint8     `json:"count"`
	RentDuration uint8     `json:"rent_duration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetCartsByUserIDAndSellerID struct {
	CartItemID   uuid.UUID `json:"cart_item_id"`
	UserID       uuid.UUID `json:"user_id"`
	ProductID    uuid.UUID `json:"product_id"`
	SellerID     uuid.UUID `json:"seller_id"`
	Count        uint8     `json:"count"`
	RentDuration uint8     `json:"rent_duration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DeleteCartByCartID struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
}

type DeleteCartByUserID struct {
	UserID uuid.UUID `json:"user_id"`
}
