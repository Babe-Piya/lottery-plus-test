package beef

import (
	"lottery-plus/beef/handler"
	"lottery-plus/beef/service"

	"github.com/gofiber/fiber/v2"
)

func AddRoute(api fiber.Router) {
	beefService := service.New()
	beefHandler := handler.New(beefService)

	v1 := api.Group("/beef")

	v1.Get("/summary", beefHandler.BeefSummary)
}
