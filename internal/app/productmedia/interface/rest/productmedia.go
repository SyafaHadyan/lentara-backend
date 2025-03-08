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

	routerGroup = routerGroup.Group("/productmedia")

	routerGroup.Get("/:id", handler.GetProductMedia)
	routerGroup.Patch("/:id", handler.UpdateProductMedia)
	routerGroup.Post("/:id", handler.CreateProductMedia)
	routerGroup.Delete("/:id", handler.DeleteProductMedia)
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

	err := ctx.BodyParser(&request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

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

func (h ProductMediaHandler) DeleteProductMedia(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid product id")
	}

	res, err := h.ProductMediaUseCase.DeleteProductMedia(productID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to delete product media")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully deleted product media",
		"payload": res,
	})
}
