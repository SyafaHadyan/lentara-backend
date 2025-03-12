package rest

import (
	"fmt"
	usecase "lentara-backend/internal/app/cart/usecase"
	productusecase "lentara-backend/internal/app/product/usecase"
	userusecase "lentara-backend/internal/app/user/usecase"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/middleware"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CartHandler struct {
	Validator      *validator.Validate
	Middleware     middleware.MiddlewareItf
	CartUseCase    usecase.CartUseCaseItf
	UserUseCase    userusecase.UserUseCaseItf
	ProductUseCase productusecase.ProductUseCaseItf
}

func NewCartHandler(routerGroup fiber.Router, validator *validator.Validate, middleware middleware.MiddlewareItf, cartUseCase usecase.CartUseCaseItf, userUseCase userusecase.UserUseCaseItf, productUseCase productusecase.ProductUseCaseItf) {
	cartHandler := CartHandler{
		Validator:      validator,
		Middleware:     middleware,
		CartUseCase:    cartUseCase,
		UserUseCase:    userUseCase,
		ProductUseCase: productUseCase,
	}

	routerGroup = routerGroup.Group("/cart")

	routerGroup.Post("/", middleware.Authentication, cartHandler.CreateCart)
	routerGroup.Patch("/:cartid", cartHandler.UpdateCart)
	routerGroup.Get("/cartid/:id", cartHandler.GetCartByID)
	routerGroup.Get("/cartuser/", middleware.Authentication, cartHandler.GetCartsByUserID)
	routerGroup.Get("/cartseller/", middleware.Authentication, cartHandler.GetCartsByUserIDAndSellerID)
	routerGroup.Delete("/cartid/:id", cartHandler.DeleteCartByCartID)
	routerGroup.Delete("/cartuser/:id", cartHandler.DeleteCartByUserID)
}

func (h CartHandler) CreateCart(ctx *fiber.Ctx) error {
	var create dto.CreateCart
	err := ctx.BodyParser(&create)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid product id")
	}

	err = h.Validator.Struct(create)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	validateRentDuration := create.RentDuration

	if validateRentDuration != 1 && validateRentDuration != 3 && validateRentDuration != 5 && validateRentDuration != 7 {
		return fiber.NewError(http.StatusBadRequest, "invalid rent duration")
	}

	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	if userID == uuid.Nil {
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	productInfo, err := h.ProductUseCase.GetProductByID(create.ProductID)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "can't find product with current id")
	}

	res, err := h.CartUseCase.CreateCart(create, userID, productInfo.SellerID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create cart")
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "successfully created cart item",
		"payload": res,
	})
}

func (h CartHandler) UpdateCart(ctx *fiber.Ctx) error {
	var update dto.UpdateCart
	err := ctx.BodyParser(&update)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

	err = h.Validator.Struct(update)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	validateRentDuration := update.RentDuration

	if validateRentDuration != 0 && validateRentDuration != 1 && validateRentDuration != 3 && validateRentDuration != 5 && validateRentDuration != 7 {
		return fiber.NewError(http.StatusBadRequest, "invalid rent duration")
	}

	cartID, err := uuid.Parse(ctx.Params("cartid"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid cart id")
	}

	_, err = h.CartUseCase.UpdateCart(update, cartID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update cart")
	}

	resUpdate, err := h.CartUseCase.GetCartByID(cartID)
	if err != nil {
		/* Proceed even if failed to get cart from database */
		log.Println("failed to get update form database")
		log.Println(err)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"messag":  "successfully udpated cart",
		"payload": resUpdate,
	})
}

func (h CartHandler) GetCartByID(ctx *fiber.Ctx) error {
	cartID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid cart id")
	}

	res, err := h.CartUseCase.GetCartByID(cartID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get cart by id")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully get cart by id",
		"payload": res,
	})
}

func (h CartHandler) GetCartsByUserID(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	res, err := h.CartUseCase.GetCartsByUserID(userID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get carts by user id")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully get carts by user id",
		"payload": res,
	})
}

func (h CartHandler) GetCartsByUserIDAndSellerID(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		fmt.Println(userID)
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	seller, err := h.CartUseCase.GetSellerListFromUserCart(userID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get seller list from user id")
	}

	for i := 0; i < len(seller); i++ {
		fmt.Println(i)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully get carts from current user grouped by seller id",
		"payload": seller,
	})
}

func (h CartHandler) DeleteCartByCartID(ctx *fiber.Ctx) error {
	cartID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid cart id")
	}

	res, err := h.CartUseCase.DeleteCartByCartID(cartID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to delete cart by id")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully deleted cart",
		"payload": res,
	})
}

func (h CartHandler) DeleteCartByUserID(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid user id")
	}

	res, err := h.CartUseCase.DeleteCartByUserID(userID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to delete cart by user id")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully deleted all carts form user id",
		"payload": res,
	})
}
