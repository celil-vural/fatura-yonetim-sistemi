package middleware

import (
	"fatura-yonetim-sistemi/infrastructure/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// AddJwtMiddleware is a function to add jwt middleware for handler
func AddJwtMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.ConfigEnv("JWT_SECRET"))},
	})
}
