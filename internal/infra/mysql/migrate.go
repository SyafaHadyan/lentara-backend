package mysql

import (
	"lentara-backend/internal/domain/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(entity.Product{}, entity.ProductSpecification{}, entity.User{}, entity.ProductMedia{}, entity.Cart{})
	if err != nil {
		return err
	}

	return nil
}
