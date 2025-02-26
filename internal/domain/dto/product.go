package dto

import (
	"time"

	"github.com/google/uuid"
)

type RequestCreateProduct struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=5"`
	Category    string `json:"category" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Stock       int32  `json:"stock" validate:"required"`
	PhotoUrl    string `json:"photo_url"`
}

type ResponseCreateProduct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int64  `json:"price"`
	Stock       int32  `json:"stock"`
	PhotoUrl    string `json:"photo_url"`
}

type GetAllProducts struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Stock       int32     `json:"stock"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
