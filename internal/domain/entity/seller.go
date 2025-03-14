package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type Seller struct {
	ID             uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	Name           string    `json:"name" gorm:"type:text;not null"`
	Email          string    `json:"email" gorm:"type:text;not null;unique"`
	Username       string    `json:"username" gorm:"type:text;not null;unique"`
	Password       string    `json:"password" gorm:"type:text;not null"`
	StoreName      string    `json:"store_name" gorm:"type:text;not null"`
	StoreLocation  string    `json:"store_location" gorm:"type:text;not null"`
	PhoneNumber    string    `json:"phone_number" gorm:"type:text;not null"`
	ProfilePicture string    `json:"profile_picture" gorm:"type:text"`
	CreatedAt      time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

func (s Seller) ParseToDTOResponseSellerRegister() dto.ResponseSellerRegister {
	return dto.ResponseSellerRegister{
		ID:             s.ID,
		Name:           s.Name,
		Email:          s.Email,
		Username:       s.Username,
		Password:       s.Password,
		StoreName:      s.StoreName,
		StoreLocation:  s.StoreLocation,
		PhoneNumber:    s.PhoneNumber,
		ProfilePicture: s.ProfilePicture,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}

func (s Seller) ParseToDTOResponseSellerLogin() dto.ResponseSellerLogin {
	return dto.ResponseSellerLogin{
		ID:             s.ID,
		Name:           s.Name,
		Email:          s.Email,
		Username:       s.Username,
		Password:       s.Password,
		StoreName:      s.StoreName,
		StoreLocation:  s.StoreLocation,
		PhoneNumber:    s.PhoneNumber,
		ProfilePicture: s.ProfilePicture,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}

func (s Seller) ParseToDTOResponseUpdateSellerInfo() dto.ResponseUpdateSellerInfo {
	return dto.ResponseUpdateSellerInfo{
		ID:             s.ID,
		Name:           s.Name,
		Email:          s.Email,
		Username:       s.Username,
		Password:       s.Password,
		StoreName:      s.StoreName,
		StoreLocation:  s.StoreLocation,
		PhoneNumber:    s.PhoneNumber,
		ProfilePicture: s.ProfilePicture,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}

func (s Seller) ParseToDTOGetSellerInfo() dto.GetSellerInfo {
	return dto.GetSellerInfo{
		ID:             s.ID,
		Name:           s.Name,
		Email:          s.Email,
		Username:       s.Username,
		Password:       s.Password,
		StoreName:      s.StoreName,
		StoreLocation:  s.StoreLocation,
		PhoneNumber:    s.PhoneNumber,
		ProfilePicture: s.ProfilePicture,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}

func (s Seller) ParseToDTOGetPublicSellerInfo() dto.GetPublicSellerInfo {
	return dto.GetPublicSellerInfo{
		Name:           s.Name,
		StoreName:      s.StoreName,
		StoreLocation:  s.StoreLocation,
		ProfilePicture: s.ProfilePicture,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}
