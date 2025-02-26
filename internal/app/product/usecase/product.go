package usecase

import (
	"lentara-backend/internal/app/product/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"

	"github.com/google/uuid"
)

type ProductUsecaseItf interface {
	GetAllProducts() (*[]dto.GetAllProducts, error)
	CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error)
	GetSpecificProduct(productID uuid.UUID) (dto.GetSpecificProduct, error)
}

type ProductUsecase struct {
	ProductRepository repository.ProductMySQLItf
}

func NewProductUsecase(productRepository repository.ProductMySQLItf) ProductUsecaseItf {
	return &ProductUsecase{
		ProductRepository: productRepository,
	}
}

func (u ProductUsecase) GetAllProducts() (*[]dto.GetAllProducts, error) {
	products := new([]entity.Product)

	err := u.ProductRepository.GetAllProducts(products)
	if err != nil {
		return nil, err
	}

	res := make([]dto.GetAllProducts, len(*products))
	for i, product := range *products {
		res[i] = product.ParseToDTOGetAllProducts()
	}

	return &res, nil

	// return u.db.Find().Error()
}

func (u ProductUsecase) CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error) {
	product := entity.Product{
		ID:            uuid.New(),
		Title:         request.Title,
		Description:   request.Description,
		Specification: request.Specification,
		Category:      request.Category,
		Price:         request.Price,
		Stock:         request.Stock,
		PhotoUrl:      request.PhotoUrl,
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

func (u ProductUsecase) GetSpecificProduct(productID uuid.UUID) (dto.GetSpecificProduct, error) {
	product := &entity.Product{
		ID: productID,
	}

	err := u.ProductRepository.GetSpecificProduct(product)
	if err != nil {
		return dto.GetSpecificProduct{}, err
	}

	return product.ParseToDTOGetSpecificProduct(), err
}
