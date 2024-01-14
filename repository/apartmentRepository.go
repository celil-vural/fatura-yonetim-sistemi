package repository

import (
	"fatura-yonetim-sistemi/entity/models"
	"gorm.io/gorm"
)

type ApartmentRepository struct {
	DB *gorm.DB
}

func (r ApartmentRepository) CreateApartment(apartment *models.Apartment) error {
	return r.DB.Create(apartment).Error
}

func (r ApartmentRepository) UpdateApartmentHirer(apartmentId uint, hirerId string) error {
	return r.DB.Model(&models.Apartment{}).Where("id = ?", apartmentId).Update("hirer_id", hirerId).Error
}
