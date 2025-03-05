package rest

import (
	"lentara-backend/internal/app/productmedia/usecase"
	"lentara-backend/internal/domain/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductMediaHandler struct {
	ProductMediaUseCase usecase.ProductMediaUsecaseItf
}

func NewProductMediahandler(routerGroup fiber.Router, productMediaUseCase usecase.ProductMediaUsecaseItf) {
	handler := ProductMediaHandler{
		ProductMediaUseCase: productMediaUseCase,
	}

	routerGroup = routerGroup.Group("/")

	routerGroup.Get("/productmedia/:id", handler.GetProductMedia)
	routerGroup.Patch("/productmedia/:id", handler.UpdateProductMedia)
	routerGroup.Post("/productmedia/:id", handler.CreateProductMedia)
}

func (h ProductMediaHandler) CreateProductMedia(ctx *fiber.Ctx) error {
	var request dto.CreateProductMedia

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid product id")
	}

	err = ctx.BodyParser(&request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

	res, err := h.ProductMediaUseCase.CreateProductMedia(productID, request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create product media")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully created product media",
		"payload": res,
	})
}

func (h ProductMediaHandler) UpdateProductMedia(ctx *fiber.Ctx) error {
	var request dto.UpdateProductMedia

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "inavlid product id")
	}

	res, err := h.ProductMediaUseCase.UpdateProductMedia(productID, request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update product media")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully updated product media",
		"payload": res,
	})
}

func (h ProductMediaHandler) GetProductMedia(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid product id")
	}

	productMedia, err := h.ProductMediaUseCase.GetProductMedia(productID)
	if err != nil {
		return fiber.NewError(http.StatusNoContent, "no record found")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": productMedia,
	})
}
