package router

import (
	"github.com/EdlanioJ/nlwpacetime/server/handlers"
	"github.com/EdlanioJ/nlwpacetime/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/auth", handlers.Auth)
	app.Post("/upload", handlers.Upload)

	jwt := middlewares.NewAuthMiddleware()

	memory := app.Group("/memories")
	memory.Use(jwt)
	memory.Get("/", handlers.ListProfileMemories)
	memory.Get("/:id", handlers.GetMemory)
	memory.Post("/", handlers.CreateMemory)
	memory.Patch("/:id", handlers.UploadMemory)
	memory.Delete("/:id", handlers.DeleteMemory)
}

