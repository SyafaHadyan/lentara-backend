package repository

import (
	"lentara-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type CartMySQLItf interface {
	CreateCart(cart *entity.Cart) error
	UpdateCart(cart *entity.Cart) error
	GetCartById(cart *entity.Cart) error
}

type CartMySQL struct {
	db *gorm.DB
}

func NewCartMySQL(db *gorm.DB) CartMySQLItf {
	return &CartMySQL{db}
}

func (r *CartMySQL) CreateCart(cart *entity.Cart) error {
	return r.db.Debug().Create(cart).Error
}

func (r *CartMySQL) UpdateCart(cart *entity.Cart) error {
	return r.db.Debug().Updates(cart).Error
}

func (r *CartMySQL) GetCartById(cart *entity.Cart) error {
	return r.db.Debug().First(cart).Error
}
