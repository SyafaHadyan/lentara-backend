package entity

import (
	"lentara-backend/internal/domain/dto"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	Name      string    `json:"name" gorm:"type:text;not null"`
	Email     string    `json:"email" gorm:"type:text;not null"`
	Password  string    `json:"password" gorm:"type:text;not null"`
	IsAdmin   string    `json:"is_admin" gorm:"type:boolean;default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}

func (u User) ParseToDTOResponseRegister() dto.ResponseRegister {
	return dto.ResponseRegister{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		IsAdmin:   u.IsAdmin,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
