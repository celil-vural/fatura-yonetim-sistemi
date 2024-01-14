package userDtos

import "fatura-yonetim-sistemi/entity/models"

type LoginDto struct {
	IdentityNumber string `json:"identity_number"`
	Password       string `json:"password"`
}
type RegisterDto struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Phone          string `json:"phone"`
	IdentityNumber string `json:"identity_number"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}
type UserDtoForSessionControl struct {
	SessionActive bool `json:"session_active"`
}

func (d UserDtoForSessionControl) FromUser(user *models.User) {
	d.SessionActive = user.SessionActive
}
func (d LoginDto) ToUser() *models.User {
	var user models.User
	user.IdentityNumber = d.IdentityNumber
	user.Password = d.Password
	return &user
}
func (d RegisterDto) ToUser() *models.User {
	var user models.User
	user.Name = d.Name
	user.Surname = d.Surname
	user.Phone = d.Phone
	user.IdentityNumber = d.IdentityNumber
	user.Email = d.Email
	user.Password = d.Password
	return &user
}
