package dto

import "github.com/google/uuid"

type CreateProductSpecification struct {
	ID             uuid.UUID `json:"product_id" validate:"required"`
	Specification1 string    `json:"specification_1"`
	Specification2 string    `json:"specification_2"`
	Specification3 string    `json:"specification_3"`
	Specification4 string    `json:"specification_4"`
	Specification5 string    `json:"specification_5"`
}

type ResponseCreateProductSpecification struct {
	ID             uuid.UUID `json:"product_id"`
	Specification1 string    `json:"specification_1"`
	Specification2 string    `json:"specification_2"`
	Specification3 string    `json:"specification_3"`
	Specification4 string    `json:"specification_4"`
	Specification5 string    `json:"specification_5"`
}

type UpdateProductSpecification struct {
	ID             uuid.UUID `json:"product_id"`
	Specification1 string    `json:"specification_1"`
	Specification2 string    `json:"specification_2"`
	Specification3 string    `json:"specification_3"`
	Specification4 string    `json:"specification_4"`
	Specification5 string    `json:"specification_5"`
}

type ResponseUpdateProductSpecification struct {
	ID             uuid.UUID `json:"product_id"`
	Specification1 string    `json:"specification_1"`
	Specification2 string    `json:"specification_2"`
	Specification3 string    `json:"specification_3"`
	Specification4 string    `json:"specification_4"`
	Specification5 string    `json:"specification_5"`
}

type GetProductSpecification struct {
	ID             uuid.UUID `json:"product_id"`
	Specification1 string    `json:"specification_1"`
	Specification2 string    `json:"specification_2"`
	Specification3 string    `json:"specification_3"`
	Specification4 string    `json:"specification_4"`
	Specification5 string    `json:"specification_5"`
}

type DeleteProductSpecification struct {
	ID uuid.UUID `json:"product_id"`
}
