package entity

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	ID        uuid.UUID `gorm:"type:varchar(36)"`
	Link      string    `gorm:"type:text"`
	Type      string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamp;autoUpdateTime"`
}
