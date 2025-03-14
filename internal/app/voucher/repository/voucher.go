package repository

import (
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VoucherMySQLItf interface {
	CreateVoucher(voucher *entity.Voucher) error
	UpdateVoucher(voucher *entity.Voucher) error
	GetVoucherByID(voucher *entity.Voucher, voucherID uuid.UUID) error
	GetAllVouchers(voucher *[]entity.Voucher) error
}

type VoucherMySQL struct {
	db *gorm.DB
}

func NewVoucherMySQL(db *gorm.DB) VoucherMySQLItf {
	return &VoucherMySQL{db}
}

func (r *VoucherMySQL) CreateVoucher(voucher *entity.Voucher) error {
	return r.db.Debug().Create(voucher).Error
}

func (r *VoucherMySQL) UpdateVoucher(voucher *entity.Voucher) error {
	return r.db.Debug().Updates(voucher).Error
}

func (r *VoucherMySQL) GetVoucherByID(voucher *entity.Voucher, voucherID uuid.UUID) error {
	return r.db.Debug().First(voucher, voucherID).Error
}

func (r *VoucherMySQL) GetAllVouchers(voucher *[]entity.Voucher) error {
	return r.db.Debug().Find(voucher).Error
}
