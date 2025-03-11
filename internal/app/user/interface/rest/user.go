package rest

import (
	"lentara-backend/internal/app/user/usecase"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	Validator   *validator.Validate
	Middleware  middleware.MiddlewareItf
	userUsecase usecase.UserUseCaseItf
}

func NewUserHandler(routerGroup fiber.Router, validator *validator.Validate, middleware middleware.MiddlewareItf, userUsecase usecase.UserUseCaseItf) {
	userHandler := UserHandler{
		Validator:   validator,
		Middleware:  middleware,
		userUsecase: userUsecase,
	}

	routerGroup = routerGroup.Group("/users")

	routerGroup.Post("/register", userHandler.RegisterUser)
	routerGroup.Post("/login", userHandler.LoginUser)
	routerGroup.Get("/info", middleware.Authentication, userHandler.GetUserInfoByUserID)
}

func (u *UserHandler) RegisterUser(ctx *fiber.Ctx) error {
	var register dto.Register
	err := ctx.BodyParser(&register)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

	err = u.Validator.Struct(register)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	res, err := u.userUsecase.Register(register)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create user")
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "successfully created user",
		"payload": res,
	})
}

func (h *UserHandler) LoginUser(ctx *fiber.Ctx) error {
	var login dto.Login

	err := ctx.BodyParser(&login)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

	err = h.Validator.Struct(login)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request body")
	}

	token, err := h.userUsecase.Login(login)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "username or password is invalid")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "logged in",
		"token":   token,
	})
}

func (h UserHandler) GetUserInfoByUserID(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	res, err := h.userUsecase.GetUserInfoByUserID(userID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get user info by user id")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully get user info by user id",
		"payload": res,
	})
}
