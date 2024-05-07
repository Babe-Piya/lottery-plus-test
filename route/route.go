package route

import (
	"lottery-plus/beef"

	"github.com/gofiber/fiber/v2"
)

func Route() *fiber.App {
	app := fiber.New()
	api := app.Group("/api/v1")

	beef.AddRoute(api)

	return app
}
