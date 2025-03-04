package middleware

import (
	"lentara-backend/internal/infra/jwt"

	"github.com/gofiber/fiber/v2"
)

type MiddlewareItf interface {
	Authentication(ctx *fiber.Ctx) error
	Authorization(ctx *fiber.Ctx) error
}

type Middleware struct {
	jwt jwt.JWT
}

func NewMiddleWare(jwt jwt.JWT) MiddlewareItf {
	return &Middleware{
		jwt: jwt,
	}
}
