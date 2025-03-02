package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetAllProducts struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetSpecificProduct struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetProductCategory struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SearchProduct struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestCreateProduct struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=5"`
	Category    string `json:"category" validate:"required,min=1"`
	Price       int64  `json:"price" validate:"required"`
	Stock       int32  `json:"stock" validate:"required"`
	PhotoUrl    string `json:"photo_url" validate:"required"`
}

type ResponseCreateProduct struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateProduct struct {
	Title       string  `json:"title" validate:"omitempty,min=3"`
	Description string  `json:"description" validate:"omitempty,min=5"`
	Category    string  `json:"category" validate:"omitempty,min=1"`
	Price       int64   `json:"price"`
	Stock       int32   `json:"stock"`
	RentCount   int32   `json:"rent_count"`
	Rating      float32 `json:"rating"`
	PhotoUrl    string  `json:"photo_url"`
}

type DeleteProduct struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResponseDeleteProduct struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
