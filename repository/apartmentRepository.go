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
