package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:char(36):primaryKey"`
	Title       string    `gorm:"type:varchar(100); not null"`
	Description string    `gorm:"type:text"`
	Price       int64     `gorm:"type:bigint;not null"`
	Stock       int32     `gorm:"type:smallint; not null"`
	PhotoUrl    string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamp;autoUpdateTime"`
}

func (p Product) ParseToDTO() dto.ResponseCreateProduct {
	return dto.ResponseCreateProduct{
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		PhotoUrl:    p.PhotoUrl,
	}
}

