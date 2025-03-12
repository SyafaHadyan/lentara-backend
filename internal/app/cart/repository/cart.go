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
	GetCartsByUserIDAndSellerID(cart *[]entity.Cart, userID uuid.UUID, sellerID uuid.UUID) error
	DeleteCartByCartID(cart *entity.Cart) error
	DeleteCartByUserID(cart *entity.Cart, userID uuid.UUID) error
	GetSellerListFromUserCart(cart *[]string, userID uuid.UUID) error
	GetUserCartsBySellerID(cart *[]entity.Cart, userID uuid.UUID, sellerID uuid.UUID) error
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

func (r *CartMySQL) GetCartsByUserIDAndSellerID(cart *[]entity.Cart, userID uuid.UUID, sellerID uuid.UUID) error {
	return r.db.Debug().Where("user_id = ?", userID).Where("seller_id = ?", sellerID).Find(cart).Error
}

func (r *CartMySQL) DeleteCartByCartID(cart *entity.Cart) error {
	return r.db.Debug().Delete(cart).Error
}

func (r *CartMySQL) DeleteCartByUserID(cart *entity.Cart, userID uuid.UUID) error {
	return r.db.Debug().Where("user_id = ?", userID).Delete(cart).Error
}

func (r *CartMySQL) GetSellerListFromUserCart(cart *[]string, userID uuid.UUID) error {
	return r.db.Debug().Raw("SELECT UNIQUE seller_id FROM carts WHERE user_id = ?", userID).Find(cart).Error
}

func (r *CartMySQL) GetUserCartsBySellerID(cart *[]entity.Cart, userID uuid.UUID, sellerID uuid.UUID) error {
	return r.db.Debug().Raw("SELECT * FROM 'carts' WHERE 'user_id' = ? AND 'seller_id' = ?", userID, sellerID).Find(cart).Error
}
