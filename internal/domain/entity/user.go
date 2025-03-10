package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	Name           string    `json:"name" gorm:"type:text;not null"`
	Email          string    `json:"email" gorm:"type:text;not null;unique"`
	Username       string    `json:"username" gorm:"type:text;not null;unique"`
	Password       string    `json:"password" gorm:"type:text;not null"`
	IsAdmin        bool      `json:"is_admin" gorm:"type:boolean;default:0"`
	RentProposal   int32     `json:"rent_proposal" gorm:"type:int"`
	RentComplete   int32     `json:"rent_complete" gorm:"type:int"`
	CancelledRent  int32     `json:"cancelled_rent" gorm:"type:int"`
	RentCount      int32     `json:"rent_count" gorm:"type:int"`
	ProfilePicture string    `json:"profile_picture" gorm:"type:text"`
	CreatedAt      time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

func (u User) ParseToDTOResponseRegister() dto.ResponseRegister {
	return dto.ResponseRegister{
		ID:             u.ID,
		Name:           u.Name,
		Email:          u.Email,
		Username:       u.Username,
		Password:       u.Password,
		IsAdmin:        u.IsAdmin,
		ProfilePicture: u.ProfilePicture,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}

func (u User) ParseToDTOResponseLogin() dto.ResponseLogin {
	return dto.ResponseLogin{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		IsAdmin:   u.IsAdmin,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
