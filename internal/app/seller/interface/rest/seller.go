package rest

import (
	"lentara-backend/internal/app/seller/usecase"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/middleware"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SellerHandler struct {
	Validator     *validator.Validate
	Middleware    middleware.MiddlewareItf
	sellerUseCase usecase.SellerUsecaseItf
}

func NewSellerHandler(routerGroup fiber.Router, validator *validator.Validate, middleware middleware.MiddlewareItf, sellerUseCase usecase.SellerUsecaseItf) {
	SellerHandler := SellerHandler{
		Validator:     validator,
		Middleware:    middleware,
		sellerUseCase: sellerUseCase,
	}

	routerGroup = routerGroup.Group("/seller")

	routerGroup.Post("/register", SellerHandler.SellerRegister)
	routerGroup.Post("/login", SellerHandler.SellerLogin)
	routerGroup.Patch("/update", middleware.Authentication, SellerHandler.UpdateSellerInfo)
	routerGroup.Get("/info", middleware.Authentication, SellerHandler.GetSellerInfo)
}

func (h *SellerHandler) SellerRegister(ctx *fiber.Ctx) error {
	var register dto.SellerRegister
	err := ctx.BodyParser(&register)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

	err = h.Validator.Struct(register)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	res, err := h.sellerUseCase.SellerRegister(register)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create seller user")
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "successfully created seller user",
		"payload": res,
	})
}

func (h *SellerHandler) SellerLogin(ctx *fiber.Ctx) error {
	var login dto.SellerLogin
	err := ctx.BodyParser(&login)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

	err = h.Validator.Struct(login)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	token, err := h.sellerUseCase.SellerLogin(login)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "username or password is invalid")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "logged in",
		"token":   token,
	})
}

func (h *SellerHandler) UpdateSellerInfo(ctx *fiber.Ctx) error {
	var update dto.UpdateSellerInfo

	err := ctx.BodyParser(&update)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

	sellerID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid seller id")
	}

	err = h.Validator.Struct(&update)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	log.Println(update)

	_, err = h.sellerUseCase.UpdateSellerInfo(update, sellerID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to uppdate seller info")
	}

	var seller dto.GetSellerInfo

	res, err := h.sellerUseCase.GetSellerInfo(seller, sellerID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get updated seller data")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully updated seller info",
		"payload": res,
	})
}

func (h *SellerHandler) GetSellerInfo(ctx *fiber.Ctx) error {
	sellerID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
	}

	var sellerInfo dto.GetSellerInfo

	res, err := h.sellerUseCase.GetSellerInfo(sellerInfo, sellerID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get seller info")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully get seller info",
		"payload": res,
	})
}
