package repository

import (
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductSpecificationMySQLItf interface {
	GetProductSpecification(productSpecification *[]entity.ProductSpecification, id uuid.UUID) error
}

type ProductSpecificationMySQL struct {
	db *gorm.DB
}

func NewProductSpecificationMySQL(db *gorm.DB) ProductSpecificationMySQLItf {
	return &ProductSpecificationMySQL{db}
}

func (r ProductSpecificationMySQL) GetProductSpecification(productSpecification *[]entity.ProductSpecification, id uuid.UUID) error {
	return r.db.Debug().Where("id = ?", id).Find(productSpecification).Error
}
