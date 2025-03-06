package repository

import (
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductMediaMySQLItf interface {
	CreateProductMedia(productMedia *entity.ProductMedia) error
	UpdateProductMedia(productMedia *entity.ProductMedia, id uuid.UUID) error
	GetProductMedia(productMedia *[]entity.ProductMedia, id uuid.UUID) error
}

type ProductMediaMySQL struct {
	db *gorm.DB
}

func NewProductMediaMySQL(db *gorm.DB) ProductMediaMySQLItf {
	return &ProductMediaMySQL{db}
}

func (r ProductMediaMySQL) CreateProductMedia(productMedia *entity.ProductMedia) error {
	return r.db.Debug().Create(productMedia).Error
}

func (r ProductMediaMySQL) UpdateProductMedia(productMedia *entity.ProductMedia, id uuid.UUID) error {
	return r.db.Debug().Save(productMedia).Error
}

func (r ProductMediaMySQL) GetProductMedia(productMedia *[]entity.ProductMedia, id uuid.UUID) error {
	return r.db.Debug().Where("id = ?", id).First(productMedia).Error
}
