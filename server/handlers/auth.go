package handlers

import (
	"github.com/EdlanioJ/nlwpacetime/server/services"
	"github.com/gofiber/fiber/v2"
)

type AuthDto struct {
	Code string `json:"code"`
}

func Auth(c *fiber.Ctx) error {
	dto := new(AuthDto)

	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := services.Auth(dto.Code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
