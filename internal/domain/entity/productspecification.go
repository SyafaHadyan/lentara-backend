package entity

import (
	"lentara-backend/internal/domain/dto"

	"github.com/google/uuid"
)

type ProductSpecification struct {
	ID              uuid.UUID `gorm:"type:char(36):foreignKey"`
	Specification_1 string    `gorm:"type:text"`
	Specification_2 string    `gorm:"type:text"`
	Specification_3 string    `gorm:"type:text"`
	Specification_4 string    `gorm:"type:text"`
	Specification_5 string    `gorm:"type:text"`
}

func (p ProductSpecification) ParseToDTOCreateProductSpecification() dto.CreateProductSpecification {
	return dto.CreateProductSpecification{
		ID:             p.ID,
		Specification1: p.Specification_1,
		Specification2: p.Specification_2,
		Specification3: p.Specification_3,
		Specification4: p.Specification_4,
		Specification5: p.Specification_5,
	}
}

func (p ProductSpecification) ParseToDTOResponseCreateProductSpecification() dto.ResponseCreateProductSpecification {
	return dto.ResponseCreateProductSpecification{
		ID:             p.ID,
		Specification1: p.Specification_1,
		Specification2: p.Specification_2,
		Specification3: p.Specification_3,
		Specification4: p.Specification_4,
		Specification5: p.Specification_5,
	}
}

func (p ProductSpecification) ParseToDTOUpdateProductSpecification() dto.UpdateProductSpecification {
	return dto.UpdateProductSpecification{
		Specification1: p.Specification_1,
		Specification2: p.Specification_2,
		Specification3: p.Specification_3,
		Specification4: p.Specification_4,
		Specification5: p.Specification_5,
	}
}

func (p ProductSpecification) ParseToDTOResponseUpdateProductSpecification() dto.ResponseUpdateProductSpecification {
	return dto.ResponseUpdateProductSpecification{
		ID:             p.ID,
		Specification1: p.Specification_1,
		Specification2: p.Specification_2,
		Specification3: p.Specification_3,
		Specification4: p.Specification_4,
		Specification5: p.Specification_5,
	}
}

func (p ProductSpecification) ParseToDTOGetProductSpecification() dto.GetProductSpecification {
	return dto.GetProductSpecification{
		ID:             p.ID,
		Specification1: p.Specification_1,
		Specification2: p.Specification_2,
		Specification3: p.Specification_3,
		Specification4: p.Specification_4,
		Specification5: p.Specification_5,
	}
}

func (p ProductSpecification) ParseToDTODeleteProductSpecification() dto.DeleteProductSpecification {
	return dto.DeleteProductSpecification{
		ID: p.ID,
	}
}
