package repository

import (
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductSpecificationMySQLItf interface {
	CreateProductSpecification(productSpecification entity.ProductSpecification) error
	UpdateProductSpecification(productSpecification *entity.ProductSpecification, id uuid.UUID) error
	GetProductSpecification(productSpecification *[]entity.ProductSpecification, id uuid.UUID) error
	DeleteProductSpecification(productSpecification *entity.ProductSpecification) error
}

type ProductSpecificationMySQL struct {
	db *gorm.DB
}

func NewProductSpecificationMySQL(db *gorm.DB) ProductSpecificationMySQLItf {
	return &ProductSpecificationMySQL{db}
}

func (r ProductSpecificationMySQL) CreateProductSpecification(productSpecification entity.ProductSpecification) error {
	return r.db.Debug().Create(productSpecification).Error
}

func (r ProductSpecificationMySQL) UpdateProductSpecification(productSpecification *entity.ProductSpecification, id uuid.UUID) error {
	return r.db.Debug().Save(productSpecification).Error
}

func (r ProductSpecificationMySQL) GetProductSpecification(productSpecification *[]entity.ProductSpecification, id uuid.UUID) error {
	return r.db.Debug().Where("id = ?", id).First(productSpecification).Error
}

func (r ProductSpecificationMySQL) DeleteProductSpecification(productSpecification *entity.ProductSpecification) error {
	return r.db.Debug().Delete(productSpecification).Error
}
