package dto

import (
	"time"

	"github.com/google/uuid"
)

type Register struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"omitempty,min=3"`
	Email     string    `json:"email" validate:"required,email"`
	Username  string    `json:"username" validate:"required,min=3,max=128"`
	Password  string    `json:"password" validate:"required,min=8,max=128"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseUserRegister struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Login struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserParam struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type ResponseRegister struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseLogin struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserInfo struct {
	Name         string    `json:"name" validate:"omitempty,required,min=3"`
	Email        string    `json:"email" validate:"omitempty,required,email"`
	Username     string    `json:"username" validate:"omitempty,required,min=3,max=128"`
	Password     string    `json:"password" validate:"omitempty,required,min=8,max=128"`
	UserLocation string    `json:"user_location" validate:"omitempty,required,min=3,max=256"`
	PhoneNumber  string    `json:"phone_number" validate:"omitempy,required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
