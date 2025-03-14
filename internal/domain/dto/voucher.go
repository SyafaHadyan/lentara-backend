package dto

import "github.com/google/uuid"

type CreateVoucher struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title" validate:"required"`
	Count uint32    `json:"count" validate:"required"`
	Value float64   `json:"value" validate:"required"`
	Type  string    `json:"type" validate:"required"`
}

type ResponseCreateVoucher struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Value float64   `json:"value"`
	Count uint32    `json:"count"`
	Type  string    `json:"type"`
}

type UpdateVoucher struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title" validate:"omitempty,required"`
	Count uint32    `json:"count" validate:"omitempty,required"`
	Value float64   `json:"value" validate:"omitempty,required"`
	Type  string    `json:"type" validate:"omitempty,required"`
}

type GetVoucherByID struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Count uint32    `json:"count"`
	Value float64   `json:"value"`
	Type  string    `json:"type"`
}

type GetAllVouchers struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Count uint32    `json:"count"`
	Value float64   `json:"value"`
	Type  string    `json:"type"`
}
