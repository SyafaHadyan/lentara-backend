package dto

import "github.com/google/uuid"

type CreateProductMedia struct {
	ID      uuid.UUID `json:"id"`
	Media1  string    `json:"media_1"`
	Media2  string    `json:"media_2"`
	Media3  string    `json:"media_3"`
	Media4  string    `json:"media_4"`
	Media5  string    `json:"media_5"`
	Media6  string    `json:"media_6"`
	Media7  string    `json:"media_7"`
	Media8  string    `json:"media_8"`
	Media9  string    `json:"media_9"`
	Media10 string    `json:"media_10"`
}

type UpdateProductMedia struct {
	ID      uuid.UUID `json:"id"`
	Media1  string    `json:"media_1"`
	Media2  string    `json:"media_2"`
	Media3  string    `json:"media_3"`
	Media4  string    `json:"media_4"`
	Media5  string    `json:"media_5"`
	Media6  string    `json:"media_6"`
	Media7  string    `json:"media_7"`
	Media8  string    `json:"media_8"`
	Media9  string    `json:"media_9"`
	Media10 string    `json:"media_10"`
}

type ResponseUpdateProductMedia struct {
	ID      uuid.UUID `json:"id"`
	Media1  string    `json:"media_1"`
	Media2  string    `json:"media_2"`
	Media3  string    `json:"media_3"`
	Media4  string    `json:"media_4"`
	Media5  string    `json:"media_5"`
	Media6  string    `json:"media_6"`
	Media7  string    `json:"media_7"`
	Media8  string    `json:"media_8"`
	Media9  string    `json:"media_9"`
	Media10 string    `json:"media_10"`
}

type GetProductMedia struct {
	ID      uuid.UUID `json:"id"`
	Media1  string    `json:"media_1"`
	Media2  string    `json:"media_2"`
	Media3  string    `json:"media_3"`
	Media4  string    `json:"media_4"`
	Media5  string    `json:"media_5"`
	Media6  string    `json:"media_6"`
	Media7  string    `json:"media_7"`
	Media8  string    `json:"media_8"`
	Media9  string    `json:"media_9"`
	Media10 string    `json:"media_10"`
}

type DeleteProductMedia struct {
	ID uuid.UUID `json:"id"`
}
