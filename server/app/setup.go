package app

import (
	"fmt"

	"github.com/EdlanioJ/nlwpacetime/server/configs"
	"github.com/EdlanioJ/nlwpacetime/server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupAndRun() error {
	err := configs.Load()
	if err != nil {
		return err
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New())
	router.SetupRoutes(app)
	cfg := configs.GetApi()

	err = app.Listen(fmt.Sprintf(":%s", cfg.Port))
	return err
}
