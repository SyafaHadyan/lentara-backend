package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) AdminUser(ctx *fiber.Ctx) error {
	isAdmin := ctx.Locals("isAdmin")
	if isAdmin == false {
		return fiber.NewError(http.StatusUnauthorized, "user unauthorized")
	}

	return ctx.Next()
}
