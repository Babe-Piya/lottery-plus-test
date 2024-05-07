package handler

import "github.com/gofiber/fiber/v2"

func (h handler) BeefSummary(c *fiber.Ctx) error {
	resp, err := h.Service.BeefSummary()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(resp)
}
