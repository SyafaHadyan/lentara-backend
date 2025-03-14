package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateMedia struct {
	ID        uuid.UUID `json:"id"`
	Link      string    `json:"link"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
