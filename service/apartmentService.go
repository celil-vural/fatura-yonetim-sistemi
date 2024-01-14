package service

import (
	"fatura-yonetim-sistemi/entity/dtos/apartmentDtos"
	"fatura-yonetim-sistemi/repository"
)

type ApartmentService struct {
	Repo repository.ApartmentRepository
}

func (service ApartmentService) CreateApartment(dto *apartmentDtos.ApartmentDtoForCreate) error {
	return service.Repo.CreateApartment(dto.ToApartment())
}
