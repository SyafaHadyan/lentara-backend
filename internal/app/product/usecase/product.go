package usecase

import (
	"lentara-backend/internal/app/product/repository"
	"lentara-backend/internal/domain/dto"
)

type ProductUsecaseItf interface {
	Intermediary() string
}

type ProductUsecase struct {
	ProductRepository repository.ProductMySQLItf
}

func NewProductUsecase(productRepository repository.ProductMySQLItf) ProductUsecaseItf {
	return &ProductUsecase{
		ProductRepository: productRepository,
	}
}

func (u ProductUsecase) Intermediary() (dto.ResponseCreateProduct, string) {
	return u.ProductRepository.GetProducts()
}

func (u ProductUseCase) CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error) {
}
