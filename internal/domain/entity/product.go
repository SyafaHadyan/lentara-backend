package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID            uuid.UUID `gorm:"type:char(36):primaryKey"`
	Title         string    `gorm:"type:varchar(100);not null"`
	Description   string    `gorm:"type:text;not null"`
	Specification string    `gorm:"type:text;not null"`
	Category      string    `gorm:"type:text;not null"`
	Price         int64     `gorm:"type:bigint;not null"`
	Stock         int32     `gorm:"type:int;not null"`
	RentCount     int32     `gorm:"type:int"`
	Rating        float32   `gorm:"type:float"`
	PhotoUrl      string    `gorm:"type:text;not null"`
	CreatedAt     time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"type:timestamp;autoUpdateTime"`
}

func (p Product) ParseToDTO() dto.ResponseCreateProduct {
	return dto.ResponseCreateProduct{
		ID:            p.ID,
		Title:         p.Title,
		Description:   p.Description,
		Specification: p.Specification,
		Category:      p.Category,
		Price:         p.Price,
		Stock:         p.Stock,
		PhotoUrl:      p.PhotoUrl,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

func (p Product) ParseToDTOGetAllProducts() dto.GetAllProducts {
	return dto.GetAllProducts{
		ID:            p.ID,
		Title:         p.Title,
		Description:   p.Description,
		Specification: p.Specification,
		Category:      p.Category,
		Price:         p.Price,
		Stock:         p.Stock,
		RentCount:     p.RentCount,
		PhotoUrl:      p.PhotoUrl,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}
