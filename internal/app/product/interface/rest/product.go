package rest

import "github.com/gofiber/fiber/v2"

type ProductHandler struct{}

func NewProductHandler(routerGroup fiber.Router) {
	handler := ProductHandler{}

	routerGroup = routerGroup.Group("/products")

	routerGroup.Get("/", handler.GetAllProducts)
}

func (h *ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	return ctx.SendString("succesfully get all products")
}
