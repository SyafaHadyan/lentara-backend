package fiber

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

const IdleTimeout = 5 * time.Second

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout:  IdleTimeout,
		ErrorHandler: HandleError,
	})

	return app
}

// func HandleError(ctx *fiber.Ctx, err error) error {
// 	return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
// }

func HandleError(ctx *fiber.Ctx, err error) error {
	// code := fiber.StatusInternalServerError
	// code := fiber.StatusTeapot
	code := fiber.StatusBadRequest

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return ctx.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
}
