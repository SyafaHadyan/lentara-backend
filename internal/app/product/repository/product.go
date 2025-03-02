package repository

import (
	"lentara-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type ProductMySQLItf interface {
	GetAllProducts(products *[]entity.Product) error
	GetSpecificProduct(products *entity.Product) error
	GetProductCategory(products *[]entity.Product, category string) error
	SearchProduct(products *[]entity.Product, query string) error
	Create(product *entity.Product) error
	UpdateProduct(products *entity.Product) error
	DeleteProduct(product *entity.Product) error
}

type ProductMySQL struct {
	db *gorm.DB
}

func NewProductMySQL(db *gorm.DB) ProductMySQLItf {
	return &ProductMySQL{db}
}

func (r ProductMySQL) GetAllProducts(products *[]entity.Product) error {
	return r.db.Debug().Find(products).Error
}

func (r ProductMySQL) GetSpecificProduct(products *entity.Product) error {
	return r.db.Debug().First(products).Error
}

func (r ProductMySQL) GetProductCategory(products *[]entity.Product, category string) error {
	return r.db.Debug().Where("category = ?", category).Find(products).Error
}

func (r ProductMySQL) SearchProduct(products *[]entity.Product, query string) error {
	return r.db.Debug().Where("title LIKE ?", "%"+query+"%").Find(products).Error
}

func (r ProductMySQL) Create(product *entity.Product) error {
	return r.db.Debug().Create(product).Error
}

func (r ProductMySQL) UpdateProduct(product *entity.Product) error {
	return r.db.Debug().Updates(product).Error
}

func (r ProductMySQL) DeleteProduct(product *entity.Product) error {
	return r.db.Debug().Delete(product).Error
}
