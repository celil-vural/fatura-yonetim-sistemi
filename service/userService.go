package service

import (
	"fatura-yonetim-sistemi/entity/dtos/userDtos"
	"fatura-yonetim-sistemi/entity/errors"
	"fatura-yonetim-sistemi/infrastructure/helpers"
	"fatura-yonetim-sistemi/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (service *UserService) Login(dto *userDtos.LoginDto) (string, error) {
	user, err := service.Repo.FindUserForLogin(dto.IdentityNumber)
	if err != nil {
		return "", err
	}
	validate := helpers.ValidatePassword(dto.Password, user.Password)
	if !validate {
		return "", errors.WrongPasswordError{ErrorType: errors.ErrorType{Message: "Wrong password", Code: 400}}
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
func (service *UserService) Register(dto *userDtos.RegisterDto) error {
	var user = dto.ToUser()
	err := service.Repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
func (service *UserService) Logout(id string) error {
	err := service.Repo.SetSessionInactive(id)
	if err != nil {
		return err
	}
	return nil
}
func (service *UserService) FindUserForSessionControl(id string) (*userDtos.UserDtoForSessionControl, error) {
	user, err := service.Repo.FindUserForSessionControl(id)
	var dto userDtos.UserDtoForSessionControl
	dto.FromUser(user)
	if err != nil {
		return nil, err
	}
	return &dto, nil
}
func (service *UserService) UserIsManager(id string) (bool, error) {
	user, err := service.Repo.FindUserForIsManager(id)
	if err != nil {
		return false, err
	}
	if user.IsManager {
		return true, nil
	}
	return false, nil
}
func (service *UserService) GetUsers() ([]userDtos.UserDtoForManager, error) {
	users, err := service.Repo.GetUsers()
	var dtos []userDtos.UserDtoForManager
	for _, user := range users {
		var userDto userDtos.UserDtoForManager
		userDto.FromUser(&user)
		dtos = append(dtos, userDto)
	}
	if err != nil {
		return dtos, err
	}
	return dtos, nil
}
