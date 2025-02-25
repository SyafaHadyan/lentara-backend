package usecase

import "lentara-backend/internal/app/product/repository"

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

func (u ProductUsecase) Intermediary() string {
	return u.ProductRepository.GetProducts()
}
