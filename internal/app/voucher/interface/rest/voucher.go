package rest

import (
	"errors"
	"lentara-backend/internal/app/voucher/usecase"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type VoucherHandler struct {
	Validator      *validator.Validate
	Middleware     middleware.MiddlewareItf
	VoucherUseCase usecase.VoucherUseCaseItf
}

func NewVoucherHandler(routerGroup fiber.Router, validator *validator.Validate, middleware middleware.MiddlewareItf, voucherUseCase usecase.VoucherUseCaseItf) {
	voucherHandler := VoucherHandler{
		Validator:      validator,
		Middleware:     middleware,
		VoucherUseCase: voucherUseCase,
	}

	routerGroup = routerGroup.Group("/voucher")

	routerGroup.Post("/new", middleware.Authentication, middleware.AdminUser, voucherHandler.CreateVoucher)
	routerGroup.Post("/", middleware.Authentication, middleware.AdminUser, voucherHandler.CreateVoucher)
	routerGroup.Get("/:id", voucherHandler.GeVoucherByID)
	routerGroup.Get("/", voucherHandler.GetAllVouchers)
}

func (h VoucherHandler) CreateVoucher(ctx *fiber.Ctx) error {
	var create dto.CreateVoucher

	err := ctx.BodyParser(&create)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	err = h.Validator.Struct(create)
	if create.Type != "percentage" && create.Type != "price" {
		err = errors.New("invalid type")
	}
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	res, err := h.VoucherUseCase.CreateVoucher(create)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create new voucher")
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "voucher creation successful",
		"payload": res,
	})
}

func (h VoucherHandler) UpdateVouhcer(ctx *fiber.Ctx) error {
	var update dto.UpdateVoucher

	voucherID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid voucher id")
	}

	err = ctx.BodyParser(&update)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	err = h.Validator.Struct(update)
	if update.Type != "percentage" && update.Type != "price" {
		err = errors.New("invalid type")
	}
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	res, err := h.VoucherUseCase.UpdateVoucher(update, voucherID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update voucher info")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "update voucher successful",
		"payload": res,
	})
}

func (h VoucherHandler) GeVoucherByID(ctx *fiber.Ctx) error {
	voucherID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid voucher id")
	}

	res, err := h.VoucherUseCase.GetVoucherByID(voucherID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get voucher by id")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfuly get voucher by id",
		"payload": res,
	})
}

func (h VoucherHandler) GetAllVouchers(ctx *fiber.Ctx) error {
	res, err := h.VoucherUseCase.GetAllVouchers()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get all products")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfuly get all vouchers",
		"payload": res,
	})
}
