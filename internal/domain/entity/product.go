package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:varchar(36):primaryKey"`
	Title       string    `gorm:"type:text;not null"`
	Description string    `gorm:"type:text;not null"`
	Category    string    `gorm:"type:text;not null"`
	Origin      string    `gorm:"type:text;not null"`
	SellerID    uuid.UUID `gorm:"type:varchar(36);not null;foreignKey"`
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
		Origin:      p.Origin,
		SellerID:    p.SellerID,
		Price:       p.Price,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOGetProductByID() dto.GetProductByID {
	return dto.GetProductByID{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Origin:      p.Origin,
		SellerID:    p.SellerID,
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
		Origin:      p.Origin,
		SellerID:    p.SellerID,
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
		Origin:      p.Origin,
		SellerID:    p.SellerID,
		Price:       p.Price,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOSearchAndCategoryProduct() dto.SearchAndCategoryProduct {
	return dto.SearchAndCategoryProduct{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Origin:      p.Origin,
		SellerID:    p.SellerID,
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
		Origin:      p.Origin,
		SellerID:    p.SellerID,
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
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Origin:      p.Origin,
		SellerID:    p.SellerID,
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
		ID: p.ID,
	}
}
