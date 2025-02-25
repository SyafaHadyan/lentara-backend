package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `gorm:"type:char(36):primaryKey"`
	Title       string    `gorm:"type:varchar(100); not null"`
	Description string    `gorm:"type:text"`
	Price       int64     `gorm:"type:bigint;not null"`
	Stock       int32     `gorm:"type:smallint; not null"`
	PhotoURL    string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:timestamp;autoUpdateTime"`
}

