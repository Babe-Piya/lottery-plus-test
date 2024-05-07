package handler

import (
	"lottery-plus/beef/service"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	Service service.Service
}

type Handler interface {
	BeefSummary(c *fiber.Ctx) error
}

func New(service service.Service) Handler {
	return handler{
		Service: service,
	}
}
