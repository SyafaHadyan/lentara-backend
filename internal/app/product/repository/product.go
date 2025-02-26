package repository

import (
	"lentara-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type ProductMySQLItf interface {
	GetAllProducts(products *[]entity.Product) error
	Create(product *entity.Product) error
	GetSpecificProduct(products *entity.Product) error
}

type ProductMySQL struct {
	db *gorm.DB
}

func NewProductMySQL(db *gorm.DB) ProductMySQLItf {
	return &ProductMySQL{db}
}

func (r ProductMySQL) GetAllProducts(products *[]entity.Product) error {
	return r.db.Debug().Find(products).Error

	// return "I use Arch btw"
}

func (r ProductMySQL) GetSpecificProduct(products *entity.Product) error {
	return r.db.Debug().First(products).Error
}

func (r ProductMySQL) Create(product *entity.Product) error {
	return r.db.Debug().Create(product).Error
}
