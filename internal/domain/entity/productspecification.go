package entity

import (
	"lentara-backend/internal/domain/dto"

	"github.com/google/uuid"
)

type ProductSpecification struct {
	ID             uuid.UUID `gorm:"type:char(36):foreignKey"`
	Specification1 string    `gorm:"type:text"`
	Specification2 string    `gorm:"type:text"`
	Specification3 string    `gorm:"type:text"`
	Specification4 string    `gorm:"type:text"`
	Specification5 string    `gorm:"type:text"`
}

func (p ProductSpecification) ParseToDTOCreateProductSpecification() dto.CreateProductSpecification {
	return dto.CreateProductSpecification{
		ID:             p.ID,
		Specification1: p.Specification1,
		Specification2: p.Specification2,
		Specification3: p.Specification3,
		Specification4: p.Specification4,
		Specification5: p.Specification5,
	}
}

func (p ProductSpecification) ParseToDTOResponseCreateProductSpecification() dto.ResponseCreateProductSpecification {
	return dto.ResponseCreateProductSpecification{
		ID:             p.ID,
		Specification1: p.Specification1,
		Specification2: p.Specification2,
		Specification3: p.Specification3,
		Specification4: p.Specification4,
		Specification5: p.Specification5,
	}
}

func (p ProductSpecification) ParseToDTOUpdateProductSpecification() dto.UpdateProductSpecification {
	return dto.UpdateProductSpecification{
		ID:             p.ID,
		Specification1: p.Specification1,
		Specification2: p.Specification2,
		Specification3: p.Specification3,
		Specification4: p.Specification4,
		Specification5: p.Specification5,
	}
}

func (p ProductSpecification) ParseToDTOGetProductSpecification() dto.GetProductSpecification {
	return dto.GetProductSpecification{
		ID:             p.ID,
		Specification1: p.Specification1,
		Specification2: p.Specification2,
		Specification3: p.Specification3,
		Specification4: p.Specification4,
		Specification5: p.Specification5,
	}
}
