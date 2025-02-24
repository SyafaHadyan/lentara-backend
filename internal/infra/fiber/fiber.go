package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const IdleTimeout = 5 * time.Second

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout: IdleTimeout,
	})

	return app
}
