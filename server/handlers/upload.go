package handlers

import (
	"github.com/EdlanioJ/nlwpacetime/server/services"
	"github.com/gofiber/fiber/v2"
)

const (
	MB = 1 << 20
)

func Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if file.Size > 5*MB {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "uploaded file must be less than 5MB",
		})
	}

	response, err := services.Upload(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response)
}
