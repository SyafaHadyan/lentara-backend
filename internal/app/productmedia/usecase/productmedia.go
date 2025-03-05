package usecase

import (
	repository "lentara-backend/internal/app/productmedia/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductMediaUsecaseItf interface {
	CreateProductMedia(productID uuid.UUID, productMedia dto.CreateProductMedia) (dto.CreateProductMedia, error)
	UpdateProductMedia(productID uuid.UUID, productMedia dto.UpdateProductMedia) (dto.ResponseUpdateProductMedia, error)
	GetProductMedia(productID uuid.UUID) (*[]dto.GetProductMedia, error)
}

type ProductMediaUsecase struct {
	ProductMediaRepository repository.ProductMediaMySQLItf
}

func NewProductMediaUsecase(productMediaRepository repository.ProductMediaMySQLItf) ProductMediaUsecaseItf {
	return &ProductMediaUsecase{
		ProductMediaRepository: productMediaRepository,
	}
}

func (u ProductMediaUsecase) CreateProductMedia(productID uuid.UUID, productMedia dto.CreateProductMedia) (dto.CreateProductMedia, error) {
	product := &entity.ProductMedia{
		ID:       productID,
		Media_1:  productMedia.Media1,
		Media_2:  productMedia.Media2,
		Media_3:  productMedia.Media3,
		Media_4:  productMedia.Media4,
		Media_5:  productMedia.Media5,
		Media_6:  productMedia.Media6,
		Media_7:  productMedia.Media7,
		Media_8:  productMedia.Media8,
		Media_9:  productMedia.Media9,
		Media_10: productMedia.Media10,
	}

	err := u.ProductMediaRepository.CreateProductMedia(product)
	if err != nil {
		return dto.CreateProductMedia{}, fiber.NewError(http.StatusInternalServerError, "failed to create product media")
	}
	return product.ParseToDTOCreateProductMedia(), nil
}

func (u ProductMediaUsecase) UpdateProductMedia(productID uuid.UUID, productMedia dto.UpdateProductMedia) (dto.ResponseUpdateProductMedia, error) {
	product := entity.ProductMedia{
		ID:       productID,
		Media_1:  productMedia.Media1,
		Media_2:  productMedia.Media2,
		Media_3:  productMedia.Media3,
		Media_4:  productMedia.Media4,
		Media_5:  productMedia.Media5,
		Media_6:  productMedia.Media6,
		Media_7:  productMedia.Media7,
		Media_8:  productMedia.Media8,
		Media_9:  productMedia.Media9,
		Media_10: productMedia.Media10,
	}

	err := u.ProductMediaRepository.UpdateProductMedia(&product, productID)
	if err != nil {
		return dto.ResponseUpdateProductMedia{}, fiber.NewError(http.StatusInternalServerError, "failed to update product media")
	}
	return product.ParseToDTOResponseUpdateProductMedia(), nil
}

func (u ProductMediaUsecase) GetProductMedia(productID uuid.UUID) (*[]dto.GetProductMedia, error) {
	productMedia := new([]entity.ProductMedia)

	err := u.ProductMediaRepository.GetProductMedia(productMedia, productID)
	if err != nil {
		return nil, fiber.NewError(http.StatusBadRequest, "product id not found")
	}

	res := make([]dto.GetProductMedia, len(*productMedia))
	for i, product := range *productMedia {
		res[i] = product.ParseToDTOGetProductMedia()
	}

	return &res, err
}
