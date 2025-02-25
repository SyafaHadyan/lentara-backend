package repository

import (
	"lentara-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type ProductMySQLItf interface {
	GetProducts() string
	Create(product *entity.Product) error
}

type ProductMySQL struct {
	db *gorm.DB
}

func NewProductMySQL(db *gorm.DB) ProductMySQLItf {
	return &ProductMySQL{db}
}

func (r ProductMySQL) GetProducts() string {
	return "I use Arch btw"
}

func (r ProductMySQL) Create(product *entity.Product) error {
	return r.db.Create(product).Error
}
