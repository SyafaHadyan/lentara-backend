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
	CreateProductSpecification(productSpecification dto.CreateProductSpecification) (dto.ResponseCreateProductSpecification, error)
	UpdateProductSpecification(ProductID uuid.UUID, productSpecification dto.UpdateProductSpecification) (dto.ResponseUpdateProductSpecification, error)
	GetProductSpecification(productID uuid.UUID) (*[]dto.GetProductSpecification, error)
	DeleteProductSpecification(productID uuid.UUID, request dto.DeleteProductSpecification) (dto.DeleteProductSpecification, error)
}

type ProductSpecificationUsecase struct {
	ProductSpecificationRepository repository.ProductSpecificationMySQLItf
}

func NewProductSpecificationUsecase(productSpecificationRepository repository.ProductSpecificationMySQLItf) ProductSpecificationUsecaseItf {
	return &ProductSpecificationUsecase{
		ProductSpecificationRepository: productSpecificationRepository,
	}
}

func (u ProductSpecificationUsecase) CreateProductSpecification(productSpecification dto.CreateProductSpecification) (dto.ResponseCreateProductSpecification, error) {
	product := entity.ProductSpecification{
		ID:              productSpecification.ID,
		Specification_1: productSpecification.Specification1,
		Specification_2: productSpecification.Specification2,
		Specification_3: productSpecification.Specification3,
		Specification_4: productSpecification.Specification4,
		Specification_5: productSpecification.Specification5,
	}

	err := u.ProductSpecificationRepository.CreateProductSpecification(product)
	if err != nil {
		return dto.ResponseCreateProductSpecification{}, fiber.NewError(http.StatusInternalServerError, "failed to create product specifications")
	}
	return product.ParseToDTOResponseCreateProductSpecification(), nil
}

func (u ProductSpecificationUsecase) UpdateProductSpecification(ProductID uuid.UUID, productSpecification dto.UpdateProductSpecification) (dto.ResponseUpdateProductSpecification, error) {
	product := &entity.ProductSpecification{
		ID:              ProductID,
		Specification_1: productSpecification.Specification1,
		Specification_2: productSpecification.Specification2,
		Specification_3: productSpecification.Specification3,
		Specification_4: productSpecification.Specification4,
		Specification_5: productSpecification.Specification5,
	}

	err := u.ProductSpecificationRepository.UpdateProductSpecification(product, ProductID)
	if err != nil {
		return dto.ResponseUpdateProductSpecification{}, fiber.NewError(http.StatusInternalServerError, "failed to update product specifications")
	}

	return product.ParseToDTOResponseUpdateProductSpecification(), nil
}

func (u ProductSpecificationUsecase) GetProductSpecification(productID uuid.UUID) (*[]dto.GetProductSpecification, error) {
	productSpecification := new([]entity.ProductSpecification)

	err := u.ProductSpecificationRepository.GetProductSpecification(productSpecification, productID)
	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, "failed to get product specifications")
	}

	res := make([]dto.GetProductSpecification, len(*productSpecification))
	for i, product := range *productSpecification {
		res[i] = product.ParseToDTOGetProductSpecification()
	}

	return &res, nil
}

func (u ProductSpecificationUsecase) DeleteProductSpecification(productID uuid.UUID, request dto.DeleteProductSpecification) (dto.DeleteProductSpecification, error) {
	productSpecification := &entity.ProductSpecification{
		ID: productID,
	}

	err := u.ProductSpecificationRepository.DeleteProductSpecification(productSpecification)
	if err != nil {
		return dto.DeleteProductSpecification{}, fiber.NewError(http.StatusBadRequest, "failed to delete product specification")
	}

	return productSpecification.ParseToDTODeleteProductSpecification(), nil
}
