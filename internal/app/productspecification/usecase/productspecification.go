package usecase

import (
	"lentara-backend/internal/app/productspecification/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductSpecificationUsecaseItf interface {
	GetProductSpecification(productID uuid.UUID) (*[]dto.GetProductSpecification, error)
}

type ProductSpecificationUsecase struct {
	ProductSpecificationRepository repository.ProductSpecificationMySQLItf
}

func NewProductSpecificationUsecase(productSpecificationRepository repository.ProductSpecificationMySQLItf) ProductSpecificationUsecaseItf {
	return &ProductSpecificationUsecase{
		ProductSpecificationRepository: productSpecificationRepository,
	}
}

func (u ProductSpecificationUsecase) GetProductSpecification(productID uuid.UUID) (*[]dto.GetProductSpecification, error) {
	productSpecification := new([]entity.ProductSpecification)

	err := u.ProductSpecificationRepository.GetProductSpecification(productSpecification, productID)
	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, "failed to get product specification")
	}

	res := make([]dto.GetProductSpecification, len(*productSpecification))
	for i, product := range *productSpecification {
		res[i] = product.ParseToDTOProductSpecification()
	}

	return &res, nil
}
