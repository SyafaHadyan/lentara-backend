package rest

import (
	"fmt"
	usecase "lentara-backend/internal/app/cart/usecase"
	productusecase "lentara-backend/internal/app/product/usecase"
	userusecase "lentara-backend/internal/app/user/usecase"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
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
	ProductUseCase productusecase.ProductUseCaseItf
	UserUseCase    userusecase.UserUseCaseItf
}

func NewCartHandler(routerGroup fiber.Router, validator *validator.Validate, middleware middleware.MiddlewareItf, cartUseCase usecase.CartUseCaseItf, userUseCase userusecase.UserUseCaseItf, productUseCase productusecase.ProductUseCaseItf) {
	cartHandler := CartHandler{
		Validator:      validator,
		Middleware:     middleware,
		CartUseCase:    cartUseCase,
		ProductUseCase: productUseCase,
		UserUseCase:    userUseCase,
	}

	routerGroup = routerGroup.Group("/cart")

	routerGroup.Post("/", middleware.Authentication, cartHandler.CreateCart)
	routerGroup.Patch("/:cartid", middleware.Authentication, cartHandler.UpdateCart)
	routerGroup.Get("/cartid/:id", middleware.Authentication, cartHandler.GetCartByID)
	routerGroup.Get("/cartuser/", middleware.Authentication, cartHandler.GetCartsByUserID)
	routerGroup.Get("/cartseller/", middleware.Authentication, cartHandler.GetCartsByUserIDAndSellerID)
	routerGroup.Delete("/cartid/:id", middleware.Authentication, cartHandler.DeleteCartByCartID)
	routerGroup.Delete("/cartuser/", middleware.Authentication, cartHandler.DeleteCartByUserID)
	routerGroup.Get("/ordersummary", middleware.Authentication, cartHandler.GetOrderSummary)
	routerGroup.Get("/summary", middleware.Authentication, cartHandler.GetCartSummary)
}

func (h CartHandler) CreateCart(ctx *fiber.Ctx) error {
	var create dto.CreateCart
	err := ctx.BodyParser(&create)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
	}

	err = h.Validator.Struct(create)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	validateRentDuration := create.RentDuration

	if validateRentDuration != 1 && validateRentDuration != 3 && validateRentDuration != 5 && validateRentDuration != 7 {
		return fiber.NewError(http.StatusBadRequest, "invalid rent duration")
	}

	productInfo, err := h.ProductUseCase.GetProductByID(create.ProductID)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "can't find product with current id")
	}

	if productInfo.Stock < uint32(create.Count) {
		return fiber.NewError(http.StatusBadRequest, "product count can't exceed product stock")
	}

	var price uint64

	switch create.RentDuration {
	case 1:
		price = uint64(productInfo.Price_1)
	case 3:
		price = uint64(productInfo.Price_3)
	case 5:
		price = uint64(productInfo.Price_5)
	case 7:
		price = uint64(productInfo.Price_7)
	}

	price *= uint64(create.Count)

	res, err := h.CartUseCase.CreateCart(create, productInfo.Title, userID, productInfo.SellerID, price)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create cart")
	}

	updateProductStock := entity.Product{
		Stock:     productInfo.Stock - uint32(create.Count),
		RentCount: productInfo.RentCount + int32(create.Count),
	}

	err = h.ProductUseCase.UpdateProduct(create.ProductID, updateProductStock.ParseToDTOUpdateProduct())
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update product stock")
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

	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
	}

	cartInfo, err := h.CartUseCase.GetCartByID(update.CartItemID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get cart info")
	}

	if userID != cartInfo.UserID {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
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

	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
	}

	res, err := h.CartUseCase.GetCartByID(cartID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get cart by id")
	}

	if userID != res.UserID {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully get cart by id",
		"payload": res,
	})
}

func (h CartHandler) GetCartsByUserID(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
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
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
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

	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
	}

	userInfo, err := h.CartUseCase.GetCartByID(cartID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get cart info")
	}

	if userInfo.UserID != userID {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
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
	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
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

func (h CartHandler) GetOrderSummary(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
	}

	res, err := h.CartUseCase.GetOrderSummary(userID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get cart summary from user id")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully get order summary",
		"payload": res,
	})
}

func (h CartHandler) GetCartSummary(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
	}

	res, totalPrice, err := h.CartUseCase.GetCartSumamry(userID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get cart summary")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":     "successfully get cart summary",
		"payload":     res,
		"total_price": totalPrice,
	})
}
