package rest

import (
	"lentara-backend/internal/app/product/usecase"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	Validator      *validator.Validate
	ProductUseCase usecase.ProductUsecaseItf
	Middleware     middleware.MiddlewareItf
}

func NewProductHandler(routerGroup fiber.Router, validator *validator.Validate, productUseCase usecase.ProductUsecaseItf, middleware middleware.MiddlewareItf) {
	handler := ProductHandler{
		Validator:      validator,
		ProductUseCase: productUseCase,
		Middleware:     middleware,
	}

	routerGroup = routerGroup.Group("/")

	routerGroup.Get("/products", handler.GetAllProducts)
	routerGroup.Get("/products/:id", handler.GetSpecificProduct)
	routerGroup.Get("/products/category/:category", handler.GetProductCategory)
	routerGroup.Get("/search/:title", handler.SearchProduct)
	routerGroup.Post("/products", middleware.Authentication, handler.CreateProduct)
	routerGroup.Patch("/products/:id", middleware.Authentication, handler.UpdateProduct)
	routerGroup.Delete("/products/:id", middleware.Authentication, handler.DeleteProduct)
}

func (h ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	res, err := h.ProductUseCase.GetAllProducts()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get products")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductHandler) GetSpecificProduct(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "not a valid uuid")
	}

	product, err := h.ProductUseCase.GetSpecificProduct(productID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get product")
	}

	return ctx.Status(http.StatusOK).JSON(product)
}

func (h ProductHandler) GetProductCategory(ctx *fiber.Ctx) error {
	res, err := h.ProductUseCase.GetProductCategory(ctx.Params("category"))
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get product category")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductHandler) SearchProduct(ctx *fiber.Ctx) error {
	res, err := h.ProductUseCase.SearchProduct(ctx.Params("title"))
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to search product")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var request dto.RequestCreateProduct
	// request := new(dto.RequestCreateProduct)

	err := ctx.BodyParser(&request)
	if err != nil {
		fiber.NewError(http.StatusInternalServerError, "failed to parse request body")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		fiber.NewError(http.StatusBadRequest, "failed to validate request")
	}

	res, err := h.ProductUseCase.CreateProduct(request)
	if err != nil {
		fiber.NewError(http.StatusInternalServerError, "failed to create product")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully created prodcut",
		"payload": res,
	})
}

func (h ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {
	var request dto.UpdateProduct

	err := ctx.BodyParser(&request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update product")
	}

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "can't find product with current id")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "payload invalid")
	}

	err = h.ProductUseCase.UpdateProduct(productID, request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update product")
	}

	product, err := h.ProductUseCase.GetSpecificProduct(productID)
	if err != nil {
		fiber.NewError(http.StatusInternalServerError, "failed to get product info")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "product udpated",
		"payload": product,
	})

	// product, productSpecification, err := h.ProductUseCase.GetSpecificProduct(productID)
	// if err != nil {
	// 	return fiber.NewError(http.StatusInternalServerError, "failed to get product info")
	// }

	// return ctx.Status(http.StatusOK).JSON(fiber.Map{
	// "product":               product,
}

func (h ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {
	var request dto.DeleteProduct

	err := ctx.BodyParser(request)
	if err != nil {
		fiber.NewError(http.StatusBadRequest, "failed to parse request with current id")
	}

	ProductID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		fiber.NewError(http.StatusInternalServerError, "failed to get product id")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		fiber.NewError(http.StatusBadRequest, "invalid product id")
	}

	res, err := h.ProductUseCase.DeleteProduct(ProductID, request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete product")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}
