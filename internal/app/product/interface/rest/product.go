package rest

import (
	"lentara-backend/internal/app/product/usecase"
	"lentara-backend/internal/domain/dto"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	ProductUseCase usecase.ProductUsecaseItf
}

func NewProductHandler(routerGroup fiber.Router, productUseCase usecase.ProductUsecaseItf) {
	handler := ProductHandler{
		ProductUseCase: productUseCase,
	}

	routerGroup = routerGroup.Group("/products")

	routerGroup.Get("/", handler.GetAllProducts)

	routerGroup.Post("/", handler.CreateProduct)
}

func (h ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	res := h.ProductUseCase.Intermediary()

	return ctx.JSON(fiber.Map{
		"message": res,
	})
	// return ctx.SendString("succesfully get all products")
}

func (h ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	// var request dto.RequestCreateProduct
	request := new(dto.RequestCreateProduct)

	err := ctx.BodyParser(request)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "failed to parse request body",
		})
	}

	return ctx.JSON(request)

	// return ctx.JSON(fiber.Map{
	// 	"message": "post",
	// })
}
