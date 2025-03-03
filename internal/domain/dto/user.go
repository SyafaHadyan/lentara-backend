package dto

import (
	"time"

	"github.com/google/uuid"
)

type Register struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" validate:"required,min=3"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8"`
	IsAdmin   string    `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseRegister struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsAdmin   string    `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
