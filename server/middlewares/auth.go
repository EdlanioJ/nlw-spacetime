package middlewares

import (
	"github.com/EdlanioJ/nlwpacetime/server/configs"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware() fiber.Handler {
	conf := configs.GetApi()

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(conf.JwtSecret)},
	})
}
