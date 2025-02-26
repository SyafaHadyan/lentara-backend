package usecase

import (
	"lentara-backend/internal/app/product/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
)

type ProductUsecaseItf interface {
	Intermediary() string
	CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error)
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

func (u ProductUsecase) CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error) {
	product := entity.Product{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		Category:    request.Category,
		Price:       request.Price,
		Stock:       request.Stock,
		PhotoUrl:    request.PhotoUrl,
	}

	err := u.ProductRepository.Create(&product)
	if err != nil {
		return dto.ResponseCreateProduct{}, err
	}

	return product.ParseToDTO(), nil

	// res := dto.ResponseCreateProduct{
	// 	Title:       product.Title,
	// 	Description: product.Description,
	// 	Price:       product.Price,
	// 	Stock:       product.Stock,
	// 	PhotoUrl:    product.PhotoUrl,
	// }

	// return dto.ResponseCreateProduct{}, nil
}
