package beef

import (
	"net/http"

	"lottery-plus/beef/handler"
	"lottery-plus/beef/service"

	repoAPI "lottery-plus/beef/repository/api"

	"github.com/gofiber/fiber/v2"
)

func AddRoute(api fiber.Router) {
	client := repoAPI.New(http.Client{})
	beefService := service.New(client)
	beefHandler := handler.New(beefService)

	v1 := api.Group("/beef")

	v1.Get("/summary", beefHandler.BeefSummary)
}
