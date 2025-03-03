package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Authentication(ctx *fiber.Ctx) error {
	authToken := ctx.GetReqHeaders()["Authorization"]

	if len(authToken) < 1 {
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	bearerToken := authToken[0]
	token := strings.Split(bearerToken, " ")

	userID, isAdmin, err := m.jwt.ValidateToken((token[1]))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "token invalid")
	}

	ctx.Locals("userID", userID.String())
	ctx.Locals("isAmdin", isAdmin)

	return ctx.Next()
}
