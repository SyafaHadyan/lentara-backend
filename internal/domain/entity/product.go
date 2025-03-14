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
	Price_1     int64     `gorm:"type:bigint;not null"`
	Price_3     int64     `gorm:"type:bigint;not null"`
	Price_5     int64     `gorm:"type:bigint;not null"`
	Price_7     int64     `gorm:"type:bigint;not null"`
	Stock       uint32    `gorm:"type:int unsigned;not null"`
	RentCount   int32     `gorm:"type:int;not null"`
	Rating      float32   `gorm:"type:float;"`
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
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOGetProductsBySellerID() dto.GetProductsBySellerID {
	return dto.GetProductsBySellerID{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Origin:      p.Origin,
		SellerID:    p.SellerID,
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
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
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
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
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOSortProductByPrice() dto.SortProductByPrice {
	return dto.SortProductByPrice{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Origin:      p.Origin,
		SellerID:    p.SellerID,
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOFilterProductByRating() dto.FilterProductByRating {
	return dto.FilterProductByRating{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Origin:      p.Origin,
		SellerID:    p.SellerID,
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
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
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
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
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
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
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (p Product) ParseToDTOUpdateProduct() dto.UpdateProduct {
	return dto.UpdateProduct{
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Origin:      p.Origin,
		SellerID:    p.SellerID,
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
		Stock:       p.Stock,
		RentCount:   p.RentCount,
		Rating:      p.Rating,
		PhotoUrl:    p.PhotoUrl,
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
		Price_1:     p.Price_1,
		Price_3:     p.Price_3,
		Price_5:     p.Price_5,
		Price_7:     p.Price_7,
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
