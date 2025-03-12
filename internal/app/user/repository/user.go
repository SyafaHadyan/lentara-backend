package repository

import (
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserMySQLItf interface {
	RegisterUser(user *entity.User) error
	LoginUser(user *entity.User) error
	UpdateUserInfo(user *entity.User) error
	GetUserUsername(user *entity.User, userParam dto.UserParam) error
	GetUserInfoByUserID(user *entity.User, userID uuid.UUID) error
}

type UserMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserMySQLItf {
	return &UserMySQL{db}
}

func (r *UserMySQL) RegisterUser(user *entity.User) error {
	return r.db.Debug().Create(user).Error
}

func (r *UserMySQL) LoginUser(user *entity.User) error {
	return r.db.Debug().First(user).Error
}

func (r *UserMySQL) UpdateUserInfo(user *entity.User) error {
	return r.db.Debug().Updates(user).Error
}

func (r *UserMySQL) GetUserUsername(user *entity.User, userParam dto.UserParam) error {
	return r.db.Debug().First(&user, userParam).Error
}

func (r *UserMySQL) GetUserInfoByUserID(user *entity.User, userID uuid.UUID) error {
	return r.db.Debug().First(user, userID).Error
}
