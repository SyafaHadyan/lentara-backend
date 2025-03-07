package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateCart struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
	UserID     uuid.UUID `json:"user_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Count      int32     `json:"count"`
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

type DeleteCart struct {
	CartItemID uuid.UUID `json:"cart_item_id"`
	UserID     uuid.UUID `json:"user_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Count      int32     `json:"count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
