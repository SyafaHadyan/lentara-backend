package rest

import (
	"lentara-backend/internal/app/product/usecase"

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
}

func (h ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	res := h.ProductUseCase.Intermediary()

	return ctx.JSON(fiber.Map{
		"message": res,
	})
	// return ctx.SendString("succesfully get all products")
}
