package rest

import (
	"lentara-backend/internal/app/product/usecase"
	"lentara-backend/internal/domain/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	Validator      *validator.Validate
	ProductUseCase usecase.ProductUsecaseItf
}

func NewProductHandler(routerGroup fiber.Router, validator *validator.Validate, productUseCase usecase.ProductUsecaseItf) {
	handler := ProductHandler{
		Validator:      validator,
		ProductUseCase: productUseCase,
	}

	routerGroup = routerGroup.Group("/products")

	routerGroup.Get("/", handler.GetAllProducts)
	routerGroup.Post("/", handler.CreateProduct)
	routerGroup.Get("/:id", handler.GetSpecificProduct)
	routerGroup.Patch("/:id", handler.UpdateProduct)
}

func (h ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	res, err := h.ProductUseCase.GetAllProducts()
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": res,
	})
	// return ctx.SendString("succesfully get all products")
}

func (h ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var request dto.RequestCreateProduct
	// request := new(dto.RequestCreateProduct)

	err := ctx.BodyParser(&request)
	if err != nil {
		// return err
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to parse request body",
		})
	}

	// TODO parse error
	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	res, err := h.ProductUseCase.CreateProduct(request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "succesfully created product",
		"payload": res,
	})

	// return ctx.JSON(res)

	// return ctx.JSON(request)

	// return ctx.JSON(fiber.Map{
	// 	"message": "post",
	// })
}

func (h ProductHandler) GetSpecificProduct(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "not a valid uuid")
	}

	res, err := h.ProductUseCase.GetSpecificProduct(productID)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "can't find uuid")
	}

	return ctx.Status(http.StatusOK).JSON(res)
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

	res, err := h.ProductUseCase.GetSpecificProduct(productID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get product info")
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
