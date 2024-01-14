package middleware

import (
	"fatura-yonetim-sistemi/infrastructure/database"
	"fatura-yonetim-sistemi/infrastructure/helpers"
	"fatura-yonetim-sistemi/repository"
	"fatura-yonetim-sistemi/service"
	"github.com/gofiber/fiber/v2"
)

func IsManager() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userService := service.UserService{Repo: repository.UserRepository{DB: database.Database.Db}}
		token := c.Get("Authorization")[7:]
		claims, err := helpers.ExtractClaims(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		id := claims["id"].(string)
		control, err := userService.UserIsManager(id)
		if err != nil || !control {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()
	}
}
