package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:char(36):primaryKey"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text;not null"`
	Category    string    `gorm:"type:text;not null"`
	Price       int64     `gorm:"type:bigint;not null"`
	Stock       int32     `gorm:"type:int;not null"`
	RentCount   int32     `gorm:"type:int"`
	Rating      float32   `gorm:"type:float"`
	PhotoUrl    string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamp;autoUpdateTime"`
}

func (p Product) ParseToDTOGetAllProducts() dto.GetAllProducts {
	return dto.GetAllProducts{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOGetSpecificProduct() dto.GetSpecificProduct {
	return dto.GetSpecificProduct{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOGetProductCategory() dto.GetProductCategory {
	return dto.GetProductCategory{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOSearchProduct() dto.SearchProduct {
	return dto.SearchProduct{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOResponseCreateProduct() dto.ResponseCreateProduct {
	return dto.ResponseCreateProduct{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTODeleteProduct() dto.DeleteProduct {
	return dto.DeleteProduct{
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOResponseDeleteProduct() dto.ResponseDeleteProduct {
	return dto.ResponseDeleteProduct{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
