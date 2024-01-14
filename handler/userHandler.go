package handler

import (
	"fatura-yonetim-sistemi/entity"
	"fatura-yonetim-sistemi/entity/dtos/userDtos"
	"fatura-yonetim-sistemi/infrastructure/database"
	"fatura-yonetim-sistemi/repository"
	"fatura-yonetim-sistemi/service"
	"github.com/gofiber/fiber/v2"
)

// Login userDtos
func Login(c *fiber.Ctx) error {
	var userService = service.UserService{Repo: repository.UserRepository{DB: database.Database.Db}}
	identityNumber := c.FormValue("identity_number")
	password := c.FormValue("password")
	var dto userDtos.LoginDto
	dto.IdentityNumber = identityNumber
	dto.Password = password
	token, err := userService.Login(&dto)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(
		entity.GlobalResponse{
			Data: map[string]string{
				"token": token,
			},
			Success:     true,
			ErrorDetail: entity.ErrorResponse{},
		},
	)
}

// Register userDtos
func Register(c *fiber.Ctx) error {
	var userService = service.UserService{Repo: repository.UserRepository{DB: database.Database.Db}}
	var dto userDtos.RegisterDto
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	err := userService.Register(&dto)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(
		entity.GlobalResponse{
			Data: map[string]string{
				"message": "User created successfully",
			},
			Success:     true,
			ErrorDetail: entity.ErrorResponse{},
		},
	)
}
