package mysql

import (
	"lentara-backend/internal/domain/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	errProduct := db.AutoMigrate(entity.Product{})
	if errProduct != nil {
		return errProduct
	}

	errProductSpecification := db.AutoMigrate(entity.ProductSpecification{})
	if errProductSpecification != nil {
		return errProductSpecification
	}

	return nil
}
