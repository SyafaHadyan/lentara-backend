package entity

import (
	"lentara-backend/internal/domain/dto"

	"github.com/google/uuid"
)

type Voucher struct {
	ID    uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	Title string    `json:"title" gorm:"type:text"`
	Count uint32    `json:"count" gorm:"type:int unsigned"`
	Value float64   `json:"value" gorm:"type:float"`
	Type  string    `json:"type" gorm:"type:text"`
}

func (v Voucher) ParseToDTOCreateVoucher() dto.CreateVoucher {
	return dto.CreateVoucher{
		ID:    v.ID,
		Title: v.Title,
		Count: v.Count,
		Value: v.Value,
		Type:  v.Type,
	}
}

func (v Voucher) ParseToDTOResponseCreateVoucher() dto.ResponseCreateVoucher {
	return dto.ResponseCreateVoucher{
		ID:    v.ID,
		Title: v.Title,
		Count: v.Count,
		Value: v.Value,
		Type:  v.Type,
	}
}

func (v Voucher) ParseToDTOUpdateVoucher() dto.UpdateVoucher {
	return dto.UpdateVoucher{
		ID:    v.ID,
		Title: v.Title,
		Count: v.Count,
		Value: v.Value,
		Type:  v.Type,
	}
}

func (v Voucher) ParseToDTOGetVoucherByID() dto.GetVoucherByID {
	return dto.GetVoucherByID{
		ID:    v.ID,
		Title: v.Title,
		Count: v.Count,
		Value: v.Value,
		Type:  v.Type,
	}
}

func (v Voucher) ParseToDTOGetAllVouchers() dto.GetAllVouchers {
	return dto.GetAllVouchers{
		ID:    v.ID,
		Title: v.Title,
		Count: v.Count,
		Value: v.Value,
		Type:  v.Type,
	}
}
