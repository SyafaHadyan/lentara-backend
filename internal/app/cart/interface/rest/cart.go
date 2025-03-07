package rest

import (
	"lentara-backend/internal/app/cart/usecase"
	"lentara-backend/internal/domain/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CartHandler struct {
	Validator   *validator.Validate
	cartUsecase usecase.CartUsecaseItf
}

func NewCartHandler(routerGroup fiber.Router, validator *validator.Validate, cartUsecase usecase.CartUsecaseItf) {
	cartHandler := CartHandler{
		Validator:   validator,
		cartUsecase: cartUsecase,
	}

	routerGroup = routerGroup.Group("/cart")

	routerGroup.Post("/:id", cartHandler.CreateCart)
}

func (c *CartHandler) CreateCart(ctx *fiber.Ctx) error {
	var create dto.CreateCart
	userID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid user id")
	}

	res, err := c.cartUsecase.CreateCart(create, userID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create cart")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully created cart",
		"payload": res,
	})
}
