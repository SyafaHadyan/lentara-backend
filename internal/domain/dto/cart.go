package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateCart struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
	UserID     uuid.UUID `json:"user_id"`
	ProductID  uuid.UUID `json:"product_id" validate:"required"`
	Count      int32     `json:"count" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UpdateCart struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
	UserID     uuid.UUID `json:"user_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Count      int32     `json:"count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetCartByID struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
	UserID     uuid.UUID `json:"user_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Count      int32     `json:"count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DeleteCartByCartID struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
}

type DeleteCartByUserID struct {
	UserID uuid.UUID `json:"user_id"`
}
