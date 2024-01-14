package userDtos

import (
	"fatura-yonetim-sistemi/entity/dtos/apartmentDtos"
	"fatura-yonetim-sistemi/entity/models"
	"gorm.io/gorm"
	"time"
)

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
type UserDtoForManager struct {
	ID               string                                    `json:"id"`
	Name             string                                    `json:"name"`
	Surname          string                                    `json:"surname"`
	Phone            string                                    `json:"phone"`
	IdentityNumber   string                                    `json:"identity_number"`
	Email            string                                    `json:"email"`
	Password         string                                    `json:"password"`
	SessionActive    bool                                      `json:"session_active"`
	CreatedAt        time.Time                                 `json:"created_at"`
	UpdatedAt        time.Time                                 `json:"updated_at"`
	DeletedAt        gorm.DeletedAt                            `json:"deleted_at"`
	IsManager        bool                                      `json:"is_manager"`
	OwnedApartments  []apartmentDtos.ApartmentDtoForManagerGet `json:"ownedApartments"`
	RentedApartments []apartmentDtos.ApartmentDtoForManagerGet `json:"rentedApartments"`
}

func (d UserDtoForManager) FromUser(user *models.User) {
	d.ID = user.ID
	d.Name = user.Name
	d.Surname = user.Surname
	d.Phone = user.Phone
	d.IdentityNumber = user.IdentityNumber
	d.Email = user.Email
	d.Password = user.Password
	d.SessionActive = user.SessionActive
	d.CreatedAt = user.CreatedAt
	d.UpdatedAt = user.UpdatedAt
	d.DeletedAt = user.DeletedAt
	d.IsManager = user.IsManager
	for _, apartment := range user.OwnedApartments {
		var apartmentDto apartmentDtos.ApartmentDtoForManagerGet
		apartmentDto.FromApartment(&apartment)
		d.OwnedApartments = append(d.OwnedApartments, apartmentDto)
	}
	for _, apartment := range user.RentedApartments {
		var apartmentDto apartmentDtos.ApartmentDtoForManagerGet
		apartmentDto.FromApartment(&apartment)
		d.RentedApartments = append(d.RentedApartments, apartmentDto)
	}
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
