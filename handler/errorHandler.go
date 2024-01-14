package handler

import (
	"fatura-yonetim-sistemi/entity"
	"fatura-yonetim-sistemi/entity/errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errors.WrongPasswordError:
		return c.Status(e.Code).JSON(
			entity.GlobalResponse{
				Data:    nil,
				Success: false,
				ErrorDetail: entity.ErrorResponse{
					Message: e.ErrorType.Message,
					Error:   e,
				},
			},
		)
	case errors.UserNotFoundError:
		return c.Status(e.Code).JSON(
			entity.GlobalResponse{
				Data:    nil,
				Success: false,
				ErrorDetail: entity.ErrorResponse{
					Message: e.ErrorType.Message,
					Error:   e,
				},
			},
		)
	case errors.UserIsNotActiveError:
		return c.Status(e.Code).JSON(
			entity.GlobalResponse{
				Data:    nil,
				Success: false,
				ErrorDetail: entity.ErrorResponse{
					Message: e.ErrorType.Message,
					Error:   e,
				},
			},
		)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(
			entity.GlobalResponse{
				Data:    nil,
				Success: false,
				ErrorDetail: entity.ErrorResponse{
					Message: "Something went wrong",
					Error:   e,
				},
			},
		)
	}
}
