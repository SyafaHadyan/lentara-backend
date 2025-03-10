package dto

import (
	"time"

	"github.com/google/uuid"
)

type SellerRegister struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email" validate:"required,email"`
	Username       string    `json:"username" validate:"required,min=3,max=128"`
	Password       string    `json:"password" validate:"required,min=8,max=128"`
	StoreName      string    `json:"store_name" validate:"required,min=3,max=128"`
	StoreLocation  string    `json:"store_location" validate:"required,min=3,max=128"`
	PhoneNumber    string    `json:"phone_number" validate:"required"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ResponseSellerRegister struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	StoreName      string    `json:"store_name"`
	StoreLocation  string    `json:"store_location"`
	PhoneNumber    string    `json:"phone_number"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SellerLogin struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}

type SellerInfo struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type ResponseSellerLogin struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	StoreLocation string    `json:"store_location"`
	PhoneNumber   string    `json:"phone_number"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UpdateSellerInfo struct {
	Name           string    `json:"name" validate:"omitempty,required,min=3"`
	Email          string    `json:"email" validate:"omitempty,required,email"`
	Username       string    `json:"username" validate:"omitempty,required,min=3,max=128"`
	Password       string    `json:"password" validate:"omitempty,required,min=8,max=128"`
	StoreLocation  string    `json:"store_location" validate:"omitempty,required,min=3,max=256"`
	PhoneNumber    string    `json:"phone_number" validate:"omitempty,required"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GetSellerInfo struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	StoreLocation  string    `json:"store_location"`
	PhoneNumber    string    `json:"phone_number"`
	ProfilePicture string    `json:"profile_picture"`
	Password       string    `json:"password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
