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
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetProductsBySellerID struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetProductByID struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
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
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SortProductByPrice struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FilterProductByRating struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
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
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SearchAndCategoryProduct struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestCreateProduct struct {
	Title       string    `json:"title" validate:"required,min=3"`
	Description string    `json:"description" validate:"required,min=5"`
	Category    string    `json:"category" validate:"required,min=1"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1" validate:"required"`
	Price_3     int64     `json:"price_3" validate:"required"`
	Price_5     int64     `json:"price_5" validate:"required"`
	Price_7     int64     `json:"price_7" validate:"required"`
	Stock       int32     `json:"stock" validate:"required"`
	PhotoUrl    string    `json:"photo_url" validate:"required"`
}

type ResponseCreateProduct struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateProduct struct {
	Title       string    `json:"title" validate:"omitempty,min=3"`
	Description string    `json:"description" validate:"omitempty,min=5"`
	Category    string    `json:"category" validate:"omitempty,min=1"`
	Origin      string    `json:"origin" validate:"omitempty,min=3"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
}

type DeleteProduct struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Origin      string    `json:"origin"`
	SellerID    uuid.UUID `json:"seller_id"`
	Price_1     int64     `json:"price_1"`
	Price_3     int64     `json:"price_3"`
	Price_5     int64     `json:"price_5"`
	Price_7     int64     `json:"price_7"`
	Stock       int32     `json:"stock"`
	RentCount   int32     `json:"rent_count"`
	Rating      float32   `json:"rating"`
	PhotoUrl    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResponseDeleteProduct struct {
	ID uuid.UUID `json:"id"`
}
