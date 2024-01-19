package handler

import (
	"fatura-yonetim-sistemi/entity"
	"fatura-yonetim-sistemi/entity/dtos/apartmentDtos"
	"fatura-yonetim-sistemi/infrastructure/database"
	"fatura-yonetim-sistemi/repository"
	"fatura-yonetim-sistemi/service"
	"github.com/gofiber/fiber/v2"
)

func CreateApartment(c *fiber.Ctx) error {
	var apartmentService = service.ApartmentService{Repo: repository.ApartmentRepository{DB: database.Database.Db}}
	var dto apartmentDtos.ApartmentDtoForCreate
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	err := apartmentService.CreateApartment(&dto)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(
		entity.GlobalResponse{
			Data: map[string]string{
				"message": "Apartment created successfully",
			},
			Success:     true,
			ErrorDetail: entity.ErrorResponse{},
		},
	)
}
func UpdateApartmentHirer(c *fiber.Ctx) error {
	var apartmentService = service.ApartmentService{Repo: repository.ApartmentRepository{DB: database.Database.Db}}
	var dto apartmentDtos.ApartmentDtoForUpdateHirer
	if err := c.BodyParser(&dto); err != nil {
		return err
	}
	err := apartmentService.UpdateApartmentHirer(&dto)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(
		entity.GlobalResponse{
			Data: map[string]string{
				"message": "Apartment hirer updated successfully",
			},
			Success:     true,
			ErrorDetail: entity.ErrorResponse{},
		},
	)

}
func GetApartments(c *fiber.Ctx) error {
	var apartmentService = service.ApartmentService{Repo: repository.ApartmentRepository{DB: database.Database.Db}}
	apartments, err := apartmentService.GetApartments()
	if err != nil {
		return err
	}
	return c.Status(200).JSON(
		entity.GlobalResponse{
			Data:        apartments,
			Success:     true,
			ErrorDetail: entity.ErrorResponse{},
		},
	)
}
