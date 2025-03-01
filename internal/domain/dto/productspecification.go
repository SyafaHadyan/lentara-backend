package dto

import "github.com/google/uuid"

type CreateProductSpecification struct {
	ID             uuid.UUID `json:"product_id" validate:"required"`
	Specification1 string    `json:"specification_1" validate:"min=3"`
	Specification2 string    `json:"specification_2" validate:"min=3"`
	Specification3 string    `json:"specification_3" validate:"min=3"`
	Specification4 string    `json:"specification_4" validate:"min=3"`
	Specification5 string    `json:"specification_5" validate:"min=3"`
}

type UpdateProductSpecification struct {
	ID             uuid.UUID `json:"product_id"`
	Specification1 string    `json:"specification_1" validate:"omitempty,min=3"`
	Specification2 string    `json:"specification_2" validate:"omitempty,min=3"`
	Specification3 string    `json:"specification_3" validate:"omitempty,min=3"`
	Specification4 string    `json:"specification_4" validate:"omitempty,min=3"`
	Specification5 string    `json:"specification_5" validate:"omitempty,min=3"`
}

type GetProductSpecification struct {
	ID             uuid.UUID `json:"product_id"`
	Specification1 string    `json:"specification_1"`
	Specification2 string    `json:"specification_2"`
	Specification3 string    `json:"specification_3"`
	Specification4 string    `json:"specification_4"`
	Specification5 string    `json:"specification_5"`
}
