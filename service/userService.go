package service

import (
	"fatura-yonetim-sistemi/entity/dtos/auth"
	"fatura-yonetim-sistemi/entity/errors"
	"fatura-yonetim-sistemi/infrastructure/helpers"
	"fatura-yonetim-sistemi/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (service *UserService) Login(dto *auth.LoginDto) (string, error) {
	user, err := service.Repo.FindUserForLogin(dto.IdentityNumber)
	if err != nil {
		return "", err
	}
	validate := helpers.ValidatePassword(dto.Password, user.Password)
	if !validate {
		return "", errors.WrongPasswordError{Message: "Wrong password", Code: 400}
	}
	err = service.Repo.SetSessionActive(user.ID)
	if err != nil {
		return "", err
	}
	token, err := helpers.GenerateJwt(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
