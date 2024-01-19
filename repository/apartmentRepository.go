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

func (r ApartmentRepository) GetApartments() ([]models.Apartment, error) {
	var apartments []models.Apartment
	err := r.DB.Preload("users").Preload("Dues").Preload("Invoice").Find(&apartments).Error
	if err != nil {
		return nil, err
	}
	return apartments, nil
}
