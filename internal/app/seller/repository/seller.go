package repository

import (
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SellerMySQLItf interface {
	SellerRegister(seller *entity.Seller) error
	SellerLogin(seller *entity.Seller) error
	GetSellerLoginInfo(seller *entity.Seller, sellerLogin dto.SellerInfo) error
	UpdateSellerInfo(seller *entity.Seller) error
	GetSellerInfo(seller *entity.Seller, sellerID uuid.UUID) error
}

type SellerMySQL struct {
	db *gorm.DB
}

func NewSellerMySQL(db *gorm.DB) SellerMySQLItf {
	return &SellerMySQL{db}
}

func (r *SellerMySQL) GetSellerLoginInfo(seller *entity.Seller, sellerLogin dto.SellerInfo) error {
	return r.db.Debug().First(seller, sellerLogin).Error
}

func (r *SellerMySQL) SellerRegister(seller *entity.Seller) error {
	return r.db.Debug().Create(seller).Error
}

func (r *SellerMySQL) SellerLogin(seller *entity.Seller) error {
	return r.db.Debug().First(seller).Error
}

func (r *SellerMySQL) UpdateSellerInfo(seller *entity.Seller) error {
	return r.db.Debug().Updates(seller).Error
}

func (r *SellerMySQL) GetSellerInfo(seller *entity.Seller, sellerID uuid.UUID) error {
	// return r.db.Debug().Raw("SELECT * FROM sellers WHERE id = ?", sellerID).Find(seller).Error
	return r.db.Debug().First(seller, sellerID).Error
}
