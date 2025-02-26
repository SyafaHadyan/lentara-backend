package rest

import (
	"lentara-backend/internal/app/product/usecase"
	"lentara-backend/internal/domain/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
}

func (h ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	res := h.ProductUseCase.Intermediary()
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
