package rest

import (
	"lentara-backend/internal/app/user/usecase"
	"lentara-backend/internal/domain/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Validator   *validator.Validate
	userUsecase usecase.UserUsecaseItf
}

func NewUserHandler(routerGroup fiber.Router, validator *validator.Validate, userUsecase usecase.UserUsecaseItf) {
	userHandler := UserHandler{
		Validator:   validator,
		userUsecase: userUsecase,
	}

	routerGroup = routerGroup.Group("/users")

	routerGroup.Post("/register", userHandler.RegisterUser)
	routerGroup.Post("/login", userHandler.LoginUser)
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
