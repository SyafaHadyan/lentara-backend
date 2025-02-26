package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:char(36):primaryKey"`
	Title       string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Category    string    `gorm:"type:text;not null"`
	Price       int64     `gorm:"type:bigint;not null"`
	Stock       int32     `gorm:"type:smallint;not null"`
	PhotoUrl    string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamp;autoUpdateTime"`
}

func (p Product) ParseToDTO() dto.ResponseCreateProduct {
	return dto.ResponseCreateProduct{
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		PhotoUrl:    p.PhotoUrl,
	}
}

func (p Product) ParseToDTOGetAllProducts() dto.GetAllProducts {
	return dto.GetAllProducts{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		PhotoUrl:    p.PhotoUrl,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
