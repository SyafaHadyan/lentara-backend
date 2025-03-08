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
	routerGroup.Patch("/:id", cartHandler.UpdateCart)
	routerGroup.Get("/cartid/:id", cartHandler.GetCartByID)
}

func (c *CartHandler) CreateCart(ctx *fiber.Ctx) error {
	var create dto.CreateCart
	err := ctx.BodyParser(&create)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid product id")
	}

	userID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid user id")
	}

	res, err := c.cartUsecase.CreateCart(create, userID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create cart")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully created cart item",
		"payload": res,
	})
}

func (c *CartHandler) UpdateCart(ctx *fiber.Ctx) error {
	var update dto.UpdateCart
	err := ctx.BodyParser(&update)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

	_, err = c.cartUsecase.UpdateCart(update)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update cart")
	}

	cartID, err := uuid.Parse(update.CartItemID.String())
	if err != nil {
		/* Proceed even if cart id is invalid */
	}

	resUpdate, err := c.cartUsecase.GetCartByID(cartID)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"messag":  "successfully udpated cart",
		"payload": resUpdate,
	})
}

func (c *CartHandler) GetCartByID(ctx *fiber.Ctx) error {
	cartID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid cart id")
	}

	res, err := c.cartUsecase.GetCartByID(cartID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get cart by id")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully get cart by id",
		"payload": res,
	})
}
