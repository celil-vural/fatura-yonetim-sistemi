package middleware

import (
	"fatura-yonetim-sistemi/entity/errors"
	"fatura-yonetim-sistemi/infrastructure/config"
	"fatura-yonetim-sistemi/infrastructure/database"
	"fatura-yonetim-sistemi/infrastructure/helpers"
	"fatura-yonetim-sistemi/repository"
	"fatura-yonetim-sistemi/service"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// AddJwtMiddleware is a function to add jwt middleware for handler
func AddJwtMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.ConfigEnv("JWT_SECRET"))},
	})
}
func TokenSessionControl() fiber.Handler {
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
		userDto, err := userService.FindUserForSessionControl(id)
		if err != nil {
			return errors.UserNotFoundError{ErrorType: errors.ErrorType{Message: "User not found", Code: 400}}
		}
		if !userDto.SessionActive {
			return errors.UserIsNotActiveError{
				ErrorType: errors.ErrorType{
					Message: "User is not active, please login again", Code: 401,
				},
			}
		}
		return c.Next()
	}
}
