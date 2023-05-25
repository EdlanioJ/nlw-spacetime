package handlers

import (
	"fmt"
	"strings"

	"github.com/EdlanioJ/nlwpacetime/server/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ListProfileMemories(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)
	memories, err := services.ListMemoriesByUser(sub)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(memories)
}

func GetMemory(c *fiber.Ctx) error {
	id := c.Params("id")
	memory, err := services.GetMemory(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(memory)
}

func CreateMemory(c *fiber.Ctx) error {
	dto := new(services.CreateMemoryDto)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)

	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("ValidationError: %s", err.Error()),
		})
	}

	dto.UserID = sub
	memory, err := services.CreateMemory(*dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(memory)
}

func UploadMemory(c *fiber.Ctx) error {
	dto := new(services.UpdateMemoryDto)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)
	id := c.Params("id")

	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	dto.UserID = sub
	dto.ID = id

	memory, err := services.UpdateMemory(*dto)
	if err != nil {
		if strings.Contains(err.Error(), "InvalidUserError") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(memory)
}

func DeleteMemory(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	sub := claims["sub"].(string)
	id := c.Params("id")

	dto := services.DeleteMemoryDto{
		ID:     id,
		UserID: sub,
	}

	err := services.DeleteMemory(dto)
	if err != nil {
		if strings.Contains(err.Error(), "InvalidUserError") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
