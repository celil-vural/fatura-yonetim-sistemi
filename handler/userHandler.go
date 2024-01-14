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
func GetUsers(c *fiber.Ctx) error {
	var userService = service.UserService{Repo: repository.UserRepository{DB: database.Database.Db}}
	users, err := userService.GetUsers()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(
		entity.GlobalResponse{
			Data:        users,
			Success:     true,
			ErrorDetail: entity.ErrorResponse{},
		},
	)
}
func UpdateUser(c *fiber.Ctx) error {
	var userService = service.UserService{Repo: repository.UserRepository{DB: database.Database.Db}}
	var dto userDtos.UserDtoUpdateForManager
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	err := userService.UpdateUser(&dto)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(
		entity.GlobalResponse{
			Data: map[string]string{
				"message": "User updated successfully",
			},
			Success:     true,
			ErrorDetail: entity.ErrorResponse{},
		},
	)
}
func DeleteUser(c *fiber.Ctx) error {
	var userService = service.UserService{Repo: repository.UserRepository{DB: database.Database.Db}}
	id := c.FormValue("id")
	err := userService.DeleteUser(id)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(
		entity.GlobalResponse{
			Data: map[string]string{
				"message": "User deleted successfully",
			},
			Success:     true,
			ErrorDetail: entity.ErrorResponse{},
		},
	)
}
