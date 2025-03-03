package repository

import (
	"lentara-backend/internal/domain/entity"

	"gorm.io/gorm"
)

type UserMySQLItf interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	Login(user *entity.User) error
}

type UserMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserMySQLItf {
	return &UserMySQL{db}
}

func (r *UserMySQL) Create(user *entity.User) error {
	return r.db.Debug().Create(user).Error
}

func (r *UserMySQL) Update(user *entity.User) error {
	return r.db.Debug().Updates(user).Error
}

func (r *UserMySQL) Login(user *entity.User) error {
	return r.db.Debug().First(user).Error
}
