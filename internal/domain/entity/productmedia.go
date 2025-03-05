package entity

import (
	"lentara-backend/internal/domain/dto"

	"github.com/google/uuid"
)

type ProductMedia struct {
	ID       uuid.UUID `gorm:"type:char(36):foreignKey"`
	Media_1  string    `gorm:"type:string"`
	Media_2  string    `gorm:"type:string"`
	Media_3  string    `gorm:"type:string"`
	Media_4  string    `gorm:"type:string"`
	Media_5  string    `gorm:"type:string"`
	Media_6  string    `gorm:"type:string"`
	Media_7  string    `gorm:"type:string"`
	Media_8  string    `gorm:"type:string"`
	Media_9  string    `gorm:"type:string"`
	Media_10 string    `gorm:"type:string"`
}

func (p ProductMedia) ParseToDTOCreateProductMedia() dto.CreateProductMedia {
	return dto.CreateProductMedia{
		ID:      p.ID,
		Media1:  p.Media_1,
		Media2:  p.Media_2,
		Media3:  p.Media_3,
		Media4:  p.Media_4,
		Media5:  p.Media_5,
		Media6:  p.Media_6,
		Media7:  p.Media_7,
		Media8:  p.Media_8,
		Media9:  p.Media_9,
		Media10: p.Media_10,
	}
}

func (p ProductMedia) ParseToDTOUpdateProductMedia() dto.UpdateProductMedia {
	return dto.UpdateProductMedia{
		ID:      p.ID,
		Media1:  p.Media_1,
		Media2:  p.Media_2,
		Media3:  p.Media_3,
		Media4:  p.Media_4,
		Media5:  p.Media_5,
		Media6:  p.Media_6,
		Media7:  p.Media_7,
		Media8:  p.Media_8,
		Media9:  p.Media_9,
		Media10: p.Media_10,
	}
}

func (p ProductMedia) ParseToDTOGetProductMedia() dto.GetProductMedia {
	return dto.GetProductMedia{
		ID:      p.ID,
		Media1:  p.Media_1,
		Media2:  p.Media_2,
		Media3:  p.Media_3,
		Media4:  p.Media_4,
		Media5:  p.Media_5,
		Media6:  p.Media_6,
		Media7:  p.Media_7,
		Media8:  p.Media_8,
		Media9:  p.Media_9,
		Media10: p.Media_10,
	}
}
