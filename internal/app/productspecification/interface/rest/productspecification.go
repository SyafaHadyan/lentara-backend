package rest

import (
	"lentara-backend/internal/app/productspecification/usecase"
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

	routerGroup = routerGroup.Group("/productspecification")

	routerGroup.Get("/:id", productSpecificationHandler.GetProductSpecification)
}

func (h ProductSpecificationHandler) GetProductSpecification(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	// res, err := h.ProductSpecificationUseCase.GetProductSpecification(uuid.Parse(ctx.Params("id")))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Not a valid uuid")
	}

	res, err := h.ProductSpecificationUseCase.GetProductSpecification(productID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "faled to get product specification")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}
