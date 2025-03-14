package repository

import (
	"lentara-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type MediaMySQLItf interface {
	UploadMedia(media *entity.Media) error
}

type MediaMySQL struct {
	db *gorm.DB
}

func NewMediaMySQL(db *gorm.DB) MediaMySQLItf {
	return &MediaMySQL{db}
}

func (r MediaMySQL) UploadMedia(media *entity.Media) error {
	return r.db.Debug().Create(media).Error
}
