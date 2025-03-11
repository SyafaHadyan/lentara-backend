package repository

import (
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartMySQLItf interface {
	CreateCart(cart *entity.Cart) error
	UpdateCart(cart *entity.Cart) error
	GetCartByID(cart *entity.Cart) error
	GetCartsByUserID(cart *[]entity.Cart, userID uuid.UUID) error
	DeleteCartByCartID(cart *entity.Cart) error
	DeleteCartByUserID(cart *entity.Cart, userID uuid.UUID) error
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

func (r *CartMySQL) GetCartByID(cart *entity.Cart) error {
	return r.db.Debug().First(cart).Error
}

func (r *CartMySQL) GetCartsByUserID(cart *[]entity.Cart, userID uuid.UUID) error {
	return r.db.Debug().Where("user_id = ?", userID).Find(cart).Error
}

func (r *CartMySQL) DeleteCartByCartID(cart *entity.Cart) error {
	return r.db.Debug().Delete(cart).Error
}

func (r *CartMySQL) DeleteCartByUserID(cart *entity.Cart, userID uuid.UUID) error {
	return r.db.Debug().Where("user_id = ?", userID).Delete(cart).Error
}
