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
	Price        uint64    `json:"price"`
	RentDuration uint8     `json:"rent_duration" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateCart struct {
	CartItemID   uuid.UUID `json:"cart_item_id"`
	UserID       uuid.UUID `json:"user_id"`
	ProductID    uuid.UUID `json:"product_id" validate:"omitempty,required"`
	SellerID     uuid.UUID `json:"seller_id"`
	Count        uint8     `json:"count" validate:"omitempty,required"`
	Price        uint64    `json:"price"`
	RentDuration uint8     `json:"rent_duration" validate:"omitempty,required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetCartByCartID struct {
	CartItemID   uuid.UUID `json:"cart_item_id"`
	UserID       uuid.UUID `json:"user_id"`
	ProductID    uuid.UUID `json:"product_id"`
	SellerID     uuid.UUID `json:"seller_id"`
	Count        uint8     `json:"count"`
	Price        uint64    `json:"price"`
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
	Price        uint64    `json:"price"`
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
	Price        uint64    `json:"price"`
	RentDuration uint8     `json:"rent_duration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetCartSummary struct {
	UserID             uuid.UUID `json:"user_id"`
	ProductCount       uint8     `json:"product_count"`
	DeliveryCost       uint64    `json:"delivery_cost"`
	ServiceCost        uint64    `json:"service_cost"`
	DepositeAmount     uint64    `json:"deposite_amount"`
	DepositePercentage uint64    `json:"deposite_percentage"`
	Voucher            uint64    `json:"voucher"`
	TotalPrice         uint64    `json:"total_price"`
}

type DeleteCartByCartID struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
}

type DeleteCartByUserID struct {
	UserID uuid.UUID `json:"user_id"`
}
