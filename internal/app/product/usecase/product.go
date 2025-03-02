package usecase

import (
	repository "lentara-backend/internal/app/product/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductUsecaseItf interface {
	GetAllProducts() (*[]dto.GetAllProducts, error)
	GetSpecificProduct(productID uuid.UUID) (dto.GetSpecificProduct, error)
	GetProductCategory(ProductCategory string) (*[]dto.GetProductCategory, error)
	SearchProduct(query string) (*[]dto.SearchProduct, error)
	CreateProduct(request dto.RequestCreateProduct) (dto.ResponseCreateProduct, error)
	UpdateProduct(ProductID uuid.UUID, request dto.UpdateProduct) error
	DeleteProduct(ProductID uuid.UUID, request dto.DeleteProduct) (dto.ResponseDeleteProduct, error)
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

func (u ProductUsecase) GetProductCategory(productCategory string) (*[]dto.GetProductCategory, error) {
	products := new([]entity.Product)

	err := u.ProductRepository.GetProductCategory(products, productCategory)
	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, "failed to get product category")
	}

	res := make([]dto.GetProductCategory, len(*products))
	for i, product := range *products {
		res[i] = product.ParseToDTOGetProductCategory()
	}

	return &res, nil
}

func (u ProductUsecase) SearchProduct(query string) (*[]dto.SearchProduct, error) {
	products := new([]entity.Product)

	err := u.ProductRepository.SearchProduct(products, query)
	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, "failed to search product")
	}

	res := make([]dto.SearchProduct, len(*products))
	for i, product := range *products {
		res[i] = product.ParseToDTOSearchProduct()
	}

	return &res, nil
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
		return dto.ResponseCreateProduct{}, fiber.NewError(http.StatusBadRequest, "failed to create product")
	}

	return product.ParseToDTOResponseCreateProduct(), nil
}

func (u ProductUsecase) UpdateProduct(ProductID uuid.UUID, request dto.UpdateProduct) error {
	product := &entity.Product{
		ID:          ProductID,
		Title:       request.Title,
		Description: request.Description,
		Category:    request.Category,
		Price:       request.Price,
		Stock:       request.Stock,
		RentCount:   request.RentCount,
		Rating:      request.Rating,
		PhotoUrl:    request.PhotoUrl,
	}

	err := u.ProductRepository.UpdateProduct(product)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to update product")
	}

	return nil
}

func (u ProductUsecase) DeleteProduct(ProductID uuid.UUID, request dto.DeleteProduct) (dto.ResponseDeleteProduct, error) {
	productID := &entity.Product{
		ID: ProductID,
	}

	err := u.ProductRepository.DeleteProduct(productID)
	if err != nil {
		return dto.ResponseDeleteProduct{}, fiber.NewError(http.StatusBadRequest, "failed to find product id")
	}

	return productID.ParseToDTOResponseDeleteProduct(), nil
}
