package rest

import (
	"lentara-backend/internal/app/productspecification/usecase"
	"lentara-backend/internal/domain/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductSpecificationHandler struct {
	Validator                   *validator.Validate
	ProductSpecificationUseCase usecase.ProductSpecificationUsecaseItf
}

func NewProductSpecificationHandler(routerGroup fiber.Router, validator *validator.Validate, productSpecificationUseCase usecase.ProductSpecificationUsecaseItf) {
	productSpecificationHandler := ProductSpecificationHandler{
		Validator:                   validator,
		ProductSpecificationUseCase: productSpecificationUseCase,
	}

	routerGroup = routerGroup.Group("/productspec")

	routerGroup.Post("/:id", productSpecificationHandler.CreateProductSpecification)
	routerGroup.Patch("/:id", productSpecificationHandler.UpdateProductSpecification)
	routerGroup.Get("/:id", productSpecificationHandler.GetProductSpecification)
	routerGroup.Delete(":id", productSpecificationHandler.DeleteProductSpecification)
}

func (h ProductSpecificationHandler) CreateProductSpecification(ctx *fiber.Ctx) error {
	var request dto.CreateProductSpecification

	err := ctx.BodyParser(&request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create product specification")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid product id")
	}

	res, err := h.ProductSpecificationUseCase.CreateProductSpecification(request, productID)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "product specification already exist, use PATCH instead")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductSpecificationHandler) UpdateProductSpecification(ctx *fiber.Ctx) error {
	var request dto.UpdateProductSpecification

	err := ctx.BodyParser(&request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to parse request body")
	}

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to get product id")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to validate payload")
	}

	res, err := h.ProductSpecificationUseCase.UpdateProductSpecification(request, productID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update product specifications")
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductSpecificationHandler) GetProductSpecification(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	// res, err := h.ProductSpecificationUseCase.GetProductSpecification(uuid.Parse(ctx.Params("id")))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "not a valid uuid")
	}

	res, err := h.ProductSpecificationUseCase.GetProductSpecification(productID)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "product specifications doesn't exist")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductSpecificationHandler) DeleteProductSpecification(ctx *fiber.Ctx) error {
	var request dto.DeleteProductSpecification

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid product id")
	}

	res, err := h.ProductSpecificationUseCase.DeleteProductSpecification(productID, request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to delete product specification")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}
