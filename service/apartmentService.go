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

func (service ApartmentService) UpdateApartmentHirer(dto *apartmentDtos.ApartmentDtoForUpdateHirer) error {
	return service.Repo.UpdateApartmentHirer(dto.ID, dto.HirerID)
}

func (service ApartmentService) GetApartments() ([]apartmentDtos.ApartmentDtoForManagerGet, error) {
	apartments, err := service.Repo.GetApartments()
	if err != nil {
		return nil, err
	}
	var dtos []apartmentDtos.ApartmentDtoForManagerGet
	for _, apartment := range apartments {
		var dto apartmentDtos.ApartmentDtoForManagerGet
		dto.FromApartment(&apartment)
		dtos = append(dtos, dto)
	}
	return dtos, nil
}
